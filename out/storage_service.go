package main

import (
	"context"
	"fmt"
	"log"

	pb "trpc-go-example/trpc/example/storage"

	"go.mongodb.org/mongo-driver/bson"
	"trpc.group/trpc-go/trpc-database/goredis"
	"trpc.group/trpc-go/trpc-database/kafka"
	"trpc.group/trpc-go/trpc-database/mongodb"
	"trpc.group/trpc-go/trpc-database/mysql"
)

type storageServiceImpl struct {
	pb.UnimplementedStorageService
}

const (
	// Defining database service names as constants so typos won't appear as often
	RedisServiceName   = "trpc.redis.db.kv"
	MysqlServiceName   = "trpc.mysql.db.kv"
	MongoDBServiceName = "trpc.mongodb.db.kv"
	KafkaServiceName   = "trpc.kafka.producer.kv"
)

type KVDocument struct {
	// for mongoDB
	K string `bson:"k"`
	V string `bson:"v"`
}

func sendKafkaLog(ctx context.Context, key, value string) error {
	producer := kafka.NewClientProxy(KafkaServiceName)

	err := producer.Produce(ctx, []byte(key), []byte(value))
	fmt.Println("sent", key, value)
	return err
}

func writeMongo(ctx context.Context, key string, value string) error {
	proxy := mongodb.NewClientProxy(MongoDBServiceName)
	_, err := proxy.InsertOne(ctx, "database", "table", bson.M{"k": key, "v": value})
	if err != nil {
		return err
	}
	return nil
}

func readMongo(ctx context.Context, key string) (string, error) {
	proxy := mongodb.NewClientProxy(MongoDBServiceName)
	var result KVDocument
	err := proxy.FindOne(ctx, "database", "table", bson.M{"k": key}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.V, nil
}

func writeRedis(ctx context.Context, key string, value string) error {
	cli, err := goredis.New(RedisServiceName)
	if err != nil {
		log.Fatal("new fail err=[%v]", err)
	}
	return cli.Set(ctx, key, value, 0).Err()
}

func readRedis(ctx context.Context, key string) (string, error) {
	cli, err := goredis.New(RedisServiceName)
	if err != nil {
		log.Fatal("new fail err=[%v]", err)
	}
	return cli.Get(ctx, key).Result()
}

func writeMysql(ctx context.Context, key string, value string) error {
	db := mysql.NewClientProxy(MysqlServiceName)
	_, err := db.Exec(ctx, "INSERT INTO kv (k, v) VALUES (?, ?) ON DUPLICATE KEY UPDATE v = ?", key, value, value)
	return err
}

func readMysql(ctx context.Context, key string) (string, error) {
	db := mysql.NewClientProxy(MysqlServiceName)
	var value string
	dest := []interface{}{&value}
	err := db.QueryRow(ctx, dest, "SELECT v FROM kv where k = ?", key)
	if err != nil {
		return "", err
	}
	return value, nil
}

// Store Stores a key-value pair into Redis, MongoDB, and MySQL
func (s *storageServiceImpl) Store(
	ctx context.Context,
	req *pb.StoreRequest,
) (*pb.StoreResponse, error) {
	// ---- mysql write ----
	// errWrite := writeMysql(ctx, req.Key, req.Value)
	// if errWrite != nil {
	// 	return nil, errWrite
	// }

	// ---- redis write ----
	// errWrite := writeRedis(ctx, req.Key, req.Value)
	// if errWrite != nil {
	// 	return nil, errWrite
	// }

	// ---- mongodb write ----
	errWrite := writeMongo(ctx, req.Key, req.Value)
	if errWrite != nil {
		return nil, errWrite
	}

	errKafkaProduce := sendKafkaLog(ctx, req.Key, req.Value)
	if errKafkaProduce != nil {
		return nil, errKafkaProduce
	}

	// Reuse code for returning response
	rsp := &pb.StoreResponse{Status: "ok"}
	return rsp, nil
}

// Fetch Fetches a value by key, likely from Redis or fallback store
func (s *storageServiceImpl) Fetch(
	ctx context.Context,
	req *pb.FetchRequest,
) (*pb.FetchResponse, error) {
	// ---- Mysql read ----
	// v, errRead := readMysql(ctx, req.Key)
	// if errRead != nil {
	// 	return nil, errRead
	// }

	// ---- Redis read ----
	// v, errRead := readRedis(ctx, req.Key)
	// if errRead != nil {
	// 	return nil, errRead
	// }

	// ---- MongoDB read ----
	v, errRead := readMongo(ctx, req.Key)
	if errRead != nil {
		return nil, errRead
	}

	// reuse code for response return
	rsp := &pb.FetchResponse{Value: v}
	fmt.Println("value fetched", rsp.Value)
	return rsp, nil
}

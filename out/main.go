package main

import (
	pb "trpc-go-example/trpc/example/storage"

	_ "trpc.group/trpc-go/trpc-database/goredis"
	_ "trpc.group/trpc-go/trpc-database/kafka"
	_ "trpc.group/trpc-go/trpc-database/mongodb"
	_ "trpc.group/trpc-go/trpc-database/mysql"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterStorageServiceService(s.Service("trpc.example.storage.StorageService"), &storageServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}

// Code generated by trpc-go/trpc-cmdline v1.0.9. DO NOT EDIT.
// source: proto/storage.proto

package storage

import (
	"context"
	"errors"
	"fmt"

	_ "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/client"
	"trpc.group/trpc-go/trpc-go/codec"
	_ "trpc.group/trpc-go/trpc-go/http"
	"trpc.group/trpc-go/trpc-go/server"
)

// START ======================================= Server Service Definition ======================================= START

// StorageServiceService defines service.
type StorageServiceService interface {
	// Store Stores a key-value pair into Redis, MongoDB, and MySQL
	Store(ctx context.Context, req *StoreRequest) (*StoreResponse, error)
	// Fetch Fetches a value by key, likely from Redis or fallback store
	Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error)
}

func StorageServiceService_Store_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &StoreRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(StorageServiceService).Store(ctx, reqbody.(*StoreRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func StorageServiceService_Fetch_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {
	req := &FetchRequest{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}) (interface{}, error) {
		return svr.(StorageServiceService).Fetch(ctx, reqbody.(*FetchRequest))
	}

	var rsp interface{}
	rsp, err = filters.Filter(ctx, req, handleFunc)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// StorageServiceServer_ServiceDesc descriptor for server.RegisterService.
var StorageServiceServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.example.storage.StorageService",
	HandlerType: ((*StorageServiceService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.example.storage.StorageService/Store",
			Func: StorageServiceService_Store_Handler,
		},
		{
			Name: "/trpc.example.storage.StorageService/Fetch",
			Func: StorageServiceService_Fetch_Handler,
		},
	},
}

// RegisterStorageServiceService registers service.
func RegisterStorageServiceService(s server.Service, svr StorageServiceService) {
	if err := s.Register(&StorageServiceServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("StorageService register error:%v", err))
	}
}

// START --------------------------------- Default Unimplemented Server Service --------------------------------- START

type UnimplementedStorageService struct{}

// Store Stores a key-value pair into Redis, MongoDB, and MySQL
func (s *UnimplementedStorageService) Store(ctx context.Context, req *StoreRequest) (*StoreResponse, error) {
	return nil, errors.New("rpc Store of service StorageService is not implemented")
}

// Fetch Fetches a value by key, likely from Redis or fallback store
func (s *UnimplementedStorageService) Fetch(ctx context.Context, req *FetchRequest) (*FetchResponse, error) {
	return nil, errors.New("rpc Fetch of service StorageService is not implemented")
}

// END --------------------------------- Default Unimplemented Server Service --------------------------------- END

// END ======================================= Server Service Definition ======================================= END

// START ======================================= Client Service Definition ======================================= START

// StorageServiceClientProxy defines service client proxy
type StorageServiceClientProxy interface {
	// Store Stores a key-value pair into Redis, MongoDB, and MySQL
	Store(ctx context.Context, req *StoreRequest, opts ...client.Option) (rsp *StoreResponse, err error)
	// Fetch Fetches a value by key, likely from Redis or fallback store
	Fetch(ctx context.Context, req *FetchRequest, opts ...client.Option) (rsp *FetchResponse, err error)
}

type StorageServiceClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewStorageServiceClientProxy = func(opts ...client.Option) StorageServiceClientProxy {
	return &StorageServiceClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *StorageServiceClientProxyImpl) Store(ctx context.Context, req *StoreRequest, opts ...client.Option) (*StoreResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.example.storage.StorageService/Store")
	msg.WithCalleeServiceName(StorageServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("example")
	msg.WithCalleeServer("storage")
	msg.WithCalleeService("StorageService")
	msg.WithCalleeMethod("Store")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &StoreResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

func (c *StorageServiceClientProxyImpl) Fetch(ctx context.Context, req *FetchRequest, opts ...client.Option) (*FetchResponse, error) {
	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)
	msg.WithClientRPCName("/trpc.example.storage.StorageService/Fetch")
	msg.WithCalleeServiceName(StorageServiceServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("example")
	msg.WithCalleeServer("storage")
	msg.WithCalleeService("StorageService")
	msg.WithCalleeMethod("Fetch")
	msg.WithSerializationType(codec.SerializationTypePB)
	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)
	rsp := &FetchResponse{}
	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}
	return rsp, nil
}

// END ======================================= Client Service Definition ======================================= END

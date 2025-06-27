package main

import (
	"context"
	"reflect"
	"testing"

	pb "trpc-go-example/trpc/example/storage"

	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/trpc-go-example/trpc/example/storage/storage_mock.go -package=storage -self_package=trpc-go-example/trpc/example/storage --source=stub/trpc-go-example/trpc/example/storage/storage.trpc.go

func Test_storageServiceImpl_Store(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storageServiceService := pb.NewMockStorageServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := storageServiceService.EXPECT().Store(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.StoreRequest) (*pb.StoreResponse, error) {
		s := &storageServiceImpl{}
		return s.Store(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.StoreRequest
		rsp *pb.StoreResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.StoreResponse
			var err error
			if rsp, err = storageServiceService.Store(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("storageServiceImpl.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("storageServiceImpl.Store() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

func Test_storageServiceImpl_Fetch(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	storageServiceService := pb.NewMockStorageServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := storageServiceService.EXPECT().Fetch(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.FetchRequest) (*pb.FetchResponse, error) {
		s := &storageServiceImpl{}
		return s.Fetch(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.FetchRequest
		rsp *pb.FetchResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.FetchResponse
			var err error
			if rsp, err = storageServiceService.Fetch(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("storageServiceImpl.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("storageServiceImpl.Fetch() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}

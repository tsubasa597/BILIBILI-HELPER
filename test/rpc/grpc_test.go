package rpc_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	g "github.com/tsubasa597/BILIBILI-HELPER/rpc"
	"github.com/tsubasa597/BILIBILI-HELPER/rpc/service"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

var (
	_ctx     = context.Background()
	_server  = grpc.NewServer()
	_connent *grpc.ClientConn
)

func TestMain(t *testing.T) {
	service.RegisterCommentServer(_server, &g.Comment{
		Log: zap.NewExample(),
	})
	service.RegisterDynamicServer(_server, &g.Dynamic{
		Log: zap.NewExample(),
	})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := _server.Serve(lis); err != nil {
			t.Error(err)
		}
	}()

	_connent, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err)
	}
}

func TestComment(t *testing.T) {
	commentClient := service.NewCommentClient(_connent)

	// 565304231165094162, 17, 49
	stream, err := commentClient.Get(_ctx, &service.CommentRequest{
		BaseCommentRequest: &service.BaseCommentRequest{
			Type: 17,
			RID:  565304231165094162,
		},
		PageSum: 49,
		PageNum: 1,
		Sort:    int32(comment.Desc),
	}, grpc.EmptyCallOption{})
	if err != nil {
		t.Error(err)
		return
	}

	for {
		_, err := stream.Recv()
		if err != nil {
			t.Error(err)
		}
	}
}

func TestDynamic(t *testing.T) {
	dynamicClient := service.NewDynamicClient(_connent)

	stream, err := dynamicClient.Get(_ctx, &service.DynamicRequest{
		BaseCommentRequest: &service.BaseDynamicRequest{
			UID: 672342685,
		},
	}, grpc.EmptyCallOption{})
	if err != nil {
		t.Error(err)
		return
	}

	for {
		_, err := stream.Recv()
		if err != nil {
			t.Error(err)
		}
	}
}

// 测试超时，取消测试
// func TestAllComment(t *testing.T) {
// 	commentClient := service.NewCommentClient(connent)

// 	// 565304231165094162, 17, 49
// 	stream, err := commentClient.GetAll(ctx, &service.AllCommentRequest{
// 		BaseCommentRequest: &service.BaseCommentRequest{
// 			Type: 17,
// 			RID:  565304231165094162,
// 		},
// 		Time: 0,
// 	}, grpc.EmptyCallOption{})
// 	if err != nil {
// 		t.Log(err)
// 		return
// 	}

// 	var count int
// 	for {
// 		_, err := stream.Recv()
// 		if err != nil {
// 			break
// 		}

// 		count++
// 	}

// 	t.Log(count)
// }

// func TestAllDynamic(t *testing.T) {
// 	dynamicClient := service.NewDynamicClient(connent)

// 	stream, err := dynamicClient.GetAll(ctx, &service.AllDynamicRequest{
// 		BaseCommentRequest: &service.BaseDynamicRequest{
// 			UID: 672342685,
// 		},
// 		Time: 0,
// 	}, grpc.EmptyCallOption{})
// 	if err != nil {
// 		t.Log(err)
// 		return
// 	}

// 	var count int
// 	for {
// 		_, err := stream.Recv()
// 		if err != nil {
// 			break
// 		}

// 		count++
// 	}
// 	t.Log(count)
// }

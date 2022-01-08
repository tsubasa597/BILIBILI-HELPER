package rpc_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/tsubasa597/BILIBILI-HELPER/info"
	g "github.com/tsubasa597/BILIBILI-HELPER/rpc"
	"github.com/tsubasa597/BILIBILI-HELPER/rpc/service"
	"go.uber.org/zap"
	grpc "google.golang.org/grpc"
)

var (
	ctx     context.Context = context.Background()
	server  *grpc.Server    = grpc.NewServer()
	connent *grpc.ClientConn
)

func init() {
	service.RegisterCommentServer(server, &g.Comment{
		Log: zap.NewExample(),
	})
	service.RegisterDynamicServer(server, &g.Dynamic{
		Log: zap.NewExample(),
	})

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	connent, err = grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
}

func TestComment(t *testing.T) {
	commentClient := service.NewCommentClient(connent)

	// 565304231165094162, 17, 49
	stream, err := commentClient.Get(ctx, &service.CommentRequest{
		BaseCommentRequest: &service.BaseCommentRequest{
			Type: 17,
			RID:  565304231165094162,
		},
		PageSum: 49,
		PageNum: 1,
		Sort:    int32(info.SortDesc),
	}, grpc.EmptyCallOption{})
	if err != nil {
		t.Log(err)
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}

		t.Log(res)
	}
}

func TestDynamic(t *testing.T) {
	dynamicClient := service.NewDynamicClient(connent)

	stream, err := dynamicClient.Get(ctx, &service.DynamicRequest{
		BaseCommentRequest: &service.BaseDynamicRequest{
			UID: 672342685,
		},
	}, grpc.EmptyCallOption{})
	if err != nil {
		t.Log(err)
		return
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}

		t.Log(res)
	}
}

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

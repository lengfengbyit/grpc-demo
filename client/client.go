package main

import (
	"context"
	"flag"
	pb "go-tour/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "服务端端口")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayHello(client)

	req := pb.HelloRequest{
		Name: "fym",
	}
	_ = SayList(client, &req)
	_ = SayRecord(client, &req)
	_ = SayRoute(client, &req)
}

func SayHello(client pb.GreeterClient) error {
	resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "fym",
	})
	log.Printf("client.SayHello resp: %s", resp.Message)
	return nil
}

func SayList(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayList(context.Background(), r)
	for {
		// 流式结束服务端返回数据
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("SayList resp: %v", resp)
	}

	return nil
}

func SayRecord(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRecord(context.Background())
	for i := 0; i < 6; i++ {
		// 持续向客户端发送请求
		_ = stream.Send(r)
	}
	// 接收客户端响应
	resp, _ := stream.CloseAndRecv()

	log.Printf("SayRecord resp: %v", resp)
	return nil
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for i := 0; i <= 6; i++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("SayRoute resp err: %v", resp)
	}

	_ = stream.CloseSend()
	return nil
}

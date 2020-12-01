package main

import (
	"context"
	"flag"
	"fmt"
	pb "go-tour/grpc-demo/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8000", "启动端口号")
	flag.Parse()
}

type GreeterServer struct{}

func (s *GreeterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	i := 0
	for {
		// 每接收一次请求，向客户端发送一个响应
		_ = stream.Send(&pb.HelloReply{
			Message: "say.route",
		})

		// 接收客户端请求, 这里会阻塞等待
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		i++
		log.Printf("SayRoute request: %v", request)
	}
}

func (s *GreeterServer) SayRecord(stream pb.Greeter_SayRecordServer) error {
	for {
		// 流式接收客户端的请求
		request, err := stream.Recv()

		// 检查客户端的请求是否结束
		if err == io.EOF {
			message := &pb.HelloReply{
				Message: "say.record",
			}

			// 响应客户端并关闭客户端连接
			return stream.SendAndClose(message)
		}
		if err != nil {
			return err
		}

		log.Printf("SayRecord request: %v", request)
	}
}

func (s *GreeterServer) SayList(request *pb.HelloRequest, stream pb.Greeter_SayListServer) error {
	for i := 0; i <= 6; i++ {
		// 流式向客户端返回数据
		_ = stream.Send(&pb.HelloReply{
			Message: "hello.list",
		})
	}

	log.Printf("SayList request: %v", request)
	return nil
}

func (s *GreeterServer) SayHello(ctx context.Context, r *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("SayHello request: %s", r.String())
	return &pb.HelloReply{Message: "hello.world"}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GreeterServer{})
	lis, _ := net.Listen("tcp", ":"+port)

	fmt.Println("Server listen: http://localhost:" + port)
	err := server.Serve(lis)
	fmt.Println(err)
}

package main

import (
	proto "github.com/FreeFlyXiaoMa/grpc_demo/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

const post  = "127.0.0.1:5004"

type server struct {

}

func (s *server)SayHello(ctx context.Context,in *proto.HelloRequest) (*proto.HelloReplay,error) {
	return &proto.HelloReplay{Message:"hello"+ in.Name},nil
}

func main()  {
	ln,err := net.Listen("tcp",post)

	if err != nil{
		fmt.Println("网络异常",err)
	}

	srv:=grpc.NewServer()
	proto.RegisterGreeterServer(srv,&server{})

	err = srv.Serve(ln)
	if err !=  nil{
		fmt.Println("网络启动异常",err)
	}
}
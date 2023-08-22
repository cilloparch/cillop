package rpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func RunServer(port int, registerServer func(server *grpc.Server)) error {
	addr := fmt.Sprintf(":%v", port)
	return RunServerOnAddr(addr, registerServer)
}

func RunServerOnAddr(addr string, registerServer func(server *grpc.Server)) error {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RPC Listening on %v\n", addr)
	return grpcServer.Serve(listen)
}

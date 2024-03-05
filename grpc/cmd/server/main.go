package main

import (
	"BogProject/grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	server := grpc.NewServer()
	srv := &api.GRPCServer{}

	api.RegisterTelegramBotServer(server, srv)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

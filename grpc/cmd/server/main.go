package main

import (
	"BogProject/grpc/pkg/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Создаем сервер gRPC.
	server := grpc.NewServer()

	// Регистрируем наш сервер в качестве сервера TelegramBot.
	api.RegisterTelegramBotServer(server, &api.GRPCServer{})

	// Слушаем на порту 8080.
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}

	// Запускаем сервер.
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}

package main

import (
	"BogProject/grpc/pkg/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("tcp:8080", grpc.WithInsecure())
	if err != nil {
		// Обработка ошибки
	}
	defer conn.Close()

	// Создание клиента gRPC
	client := api.NewTelegramBotClient(conn)

	// Создание объекта запроса
	request := &api.MessageRequestTelegram{
		User:    "user1",
		Message: "Hello, server!",
	}

	// Вызов метода на сервере
	response, err := client.GetMessages(context.Background(), request)
	if err != nil {
		// Обработка ошибки
	}

	for _, message := range response.Messages {
		fmt.Println(message.User, message.Message)
	}

}

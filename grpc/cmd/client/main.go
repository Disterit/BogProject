package main

import (
	"BogProject/grpc/pkg/api"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	// Устанавливаем соединение с сервером gRPC.
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Не удалось установить соединение с сервером: %v", err)
	}
	defer conn.Close()

	// Создаем клиента gRPC.
	client := api.NewTelegramBotClient(conn)

	// Создаем объект запроса.
	request := &api.MessageRequestTelegram{
		User:    "user1",
		Message: "Hello, server!",
	}

	// Вызываем метод на сервере.
	response, err := client.GetMessages(context.Background(), request)
	if err != nil {
		log.Fatalf("Ошибка при вызове метода GetMessages: %v", err)
	}

	// Выводим полученные сообщения на консоль.
	for _, message := range response.Messages {
		fmt.Println(message.User, message.Message)
	}
}

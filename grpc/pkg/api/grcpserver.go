package api

import (
	"context"
)

type GRPCServer struct{}

func (s *GRPCServer) mustEmbedUnimplementedTelegramBotServer() {
	// no-op
}

func (s *GRPCServer) GetMessages(ctx context.Context, req *MessageRequestTelegram) (*MessageResponse, error) {
	response := &MessageResponse{
		Messages: []*MessageRequestTelegram{
			req, // Добавляем переданный запрос в массив сообщений
		},
	}

	return response, nil
}

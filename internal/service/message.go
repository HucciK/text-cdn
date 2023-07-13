package service

import (
	"context"
	"log"
	"text-cdn/internal/core"
)

type MessageRepository interface {
	GetMessage(ctx context.Context, msg core.Message) (core.Message, error)
}

type MessageService struct {
	logger *log.Logger
	MessageRepository
}

func NewMessageService(log *log.Logger, m MessageRepository) *MessageService {
	return &MessageService{
		logger:            log,
		MessageRepository: m,
	}
}

func (m MessageService) GetMessage(ctx context.Context, msg core.Message) (core.Message, error) {
	return m.MessageRepository.GetMessage(ctx, msg)
}

package repository

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"text-cdn/internal/core"
)

type MessageMongoDB struct {
	logger *log.Logger
	db     *mongo.Database
}

func NewMessageRepository(log *log.Logger, db *mongo.Database) *MessageMongoDB {
	return &MessageMongoDB{
		logger: log,
		db:     db,
	}
}

func (m MessageMongoDB) GetMessage(ctx context.Context, msg core.Message) (core.Message, error) {
	var resultMsg core.Message
	res := m.db.Collection("messages").FindOne(ctx, msg)
	if err := res.Decode(&resultMsg); err != nil {
		m.logger.Println(fmt.Errorf("error while trying to find message: %v", err))
		return core.Message{}, fmt.Errorf("error while trying to find message: %v", err)
	}
	return resultMsg, nil
}

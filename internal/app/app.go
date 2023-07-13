package app

import (
	"context"
	"fmt"
	"log"
	"text-cdn/config"
	"text-cdn/internal/repository"
	"text-cdn/internal/service"
	"text-cdn/internal/transport"
	"text-cdn/internal/transport/handler"
	"text-cdn/pkg/mongo"
)

type App struct {
	logger *log.Logger
	Config config.Config
}

func New(log *log.Logger, cfg config.Config) *App {
	return &App{
		logger: log,
		Config: cfg,
	}
}

func (a App) Run() error {
	ctx := context.TODO()

	db, err := mongo.NewMongoClient(ctx, a.Config.DBСonfig)
	if err != nil {
		return fmt.Errorf("error while creating new mongo client: %v", err)
	}

	messageRepository := repository.NewMessageRepository(a.logger, db)

	messageService := service.NewMessageService(a.logger, messageRepository)

	handler := handler.NewHandler(a.logger, messageService)

	server := transport.NewServer(a.Config.ServerConfig, handler.InitRoutes())
	if err := server.Start(); err != nil {
		return err
	}

	return nil
}

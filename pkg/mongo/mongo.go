package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"text-cdn/config"
)

func NewMongoClient(ctx context.Context, cfg config.DBСonfig) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(cfg.ToSourceString())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error while trying to connect to mongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("error while trying to ping mongoDB: %v", err)
	}
	db := client.Database(cfg.DbName)

	return db, nil
}

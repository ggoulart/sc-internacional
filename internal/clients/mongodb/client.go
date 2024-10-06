package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	config      Config
	MongoClient *mongo.Client
}

func NewMongoClient(config *Config) (*Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := config.MongoURI
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	return &Client{config: *config, MongoClient: client}, nil
}

func (c Client) Database() *mongo.Database {
	return c.MongoClient.Database(c.config.DBName)
}

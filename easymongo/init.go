package easymongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client   *mongo.Client
	Context  context.Context
	Database *mongo.Database
}

// Constructor
func Connect(options *options.ClientOptions, database string) (*MongoClient, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options)

	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		return nil, err
	}

	return &MongoClient{
		client,
		ctx,
		client.Database(database),
	}, nil
}

// Standard methods
func (client *MongoClient) Collection(name string, opts ...*options.CollectionOptions) *CollectionCursor {
	return &CollectionCursor{
		client.Context,
		client.Database.Collection(name, opts...),
	}
}

func (client *MongoClient) Disconnect() error {
	return client.Client.Disconnect(client.Context)
}

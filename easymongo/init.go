package easymongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	Client            *mongo.Client
	Context           context.Context
	Database          *mongo.Database
	LoadedCollections map[string]*CollectionCursor
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
		map[string]*CollectionCursor{},
	}, nil
}

// Standard methods
func (client *MongoClient) Collection(name string, opts ...*options.CollectionOptions) *CollectionCursor {
	if cursor, ok := client.LoadedCollections[name]; ok {
		return cursor
	}

	client.LoadedCollections[name] = &CollectionCursor{
		client.Context,
		client.Database.Collection(name, opts...),
	}

	return client.LoadedCollections[name]
}

func (client *MongoClient) UnloadCollection(name string) bool {
	_, ok := client.LoadedCollections[name]

	if ok {
		delete(client.LoadedCollections, name)
	}

	return ok
}

func (client *MongoClient) Disconnect() error {
	clear(client.LoadedCollections)
	return client.Client.Disconnect(client.Context)
}

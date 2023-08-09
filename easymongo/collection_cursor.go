package easymongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionCursor struct {
	Context    context.Context
	Collection *mongo.Collection
}

func (cursor *CollectionCursor) Count(filter any) (int64, error) {
	return cursor.Collection.CountDocuments(cursor.Context, filter)
}

func (cursor *CollectionCursor) Insert(data ...any) error {
	for _, document := range data {
		_, err := cursor.Collection.InsertOne(cursor.Context, document)

		if err != nil {
			return err
		}
	}

	return nil
}

func (cursor *CollectionCursor) FindOne(filter any, data any, opts ...*options.FindOneOptions) error {
	return cursor.Collection.FindOne(context.TODO(), filter, opts...).Decode(&data)
}

package repository

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoURL = os.Getenv("MONGODB_URL")
var biblionDatabase = os.Getenv("MONGODB_DATABASE")

const defaultTimeout = 30 * time.Second

func getCollection(name string) (*mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	collection := client.Database(biblionDatabase).Collection(name)
	return collection, err
}

type Repository struct {
	CollectionName string
}

func parseCursorDocuments(ctx context.Context, cursor *mongo.Cursor) ([]map[string]interface{}, error) {
	defer cursor.Close(ctx)

	var documents []map[string]interface{}
	for cursor.Next(ctx) {
		var document bson.M
		err := cursor.Decode(&document)
		if err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	return documents, nil
}

func (r Repository) GetAll() ([]map[string]interface{}, error) {
	collection, err := getCollection(r.CollectionName)
	if err != nil {
		log.Println("Error getting collection", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		log.Println("Error finding data", err)
		return nil, err
	}

	documents, err := parseCursorDocuments(ctx, cursor)
	return documents, err
}

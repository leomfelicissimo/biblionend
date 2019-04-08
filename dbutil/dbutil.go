package dbutil

import (
	"context"
	"errors"
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

// Repository represents a data repository
type Repository struct {
	CollectionName string
}

// Cursor represents a cursor of data
type Cursor interface {
	Next(context.Context) bool
	Decode(val interface{}) error
	Close(context.Context) error
}

func getCollection(name string) (*mongo.Collection, error) {
	if name == "" {
		return nil, errors.New("getCollection: name is required")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}

	collection := client.Database(biblionDatabase).Collection(name)
	return collection, err
}

func parseCursor(ctx context.Context, cursor Cursor) ([]map[string]interface{}, error) {
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

func ExecuteManyResult(collectionName string, filter interface{}) ([]map[string]interface{}, error) {
	collection, err := getCollection(collectionName)
	if err != nil {
		log.Println("Error getting collection", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)

	if err != nil {
		log.Println("Error finding data", err)
		return nil, err
	}

	documents, err := parseCursor(ctx, cursor)
	return documents, err
}

func ExecuteSingleResult(collectionName string, filter interface{}) (map[string]interface{}, error) {
	collection, err := getCollection(collectionName)
	if err != nil {
		log.Println("Error getting collection", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	result := collection.FindOne(ctx, filter)

	var document bson.D
	err = result.Decode(&document)

	return document.Map(), err
}

// FindAll method performs a find all in the given collection
func (r Repository) FindAll() ([]map[string]interface{}, error) {
	return ExecuteManyResult(r.CollectionName, bson.D{})
}

// FindBy method performs a find getting a single result using a specific field
func (r Repository) FindBy(field string, value interface{}) (map[string]interface{}, error) {
	return ExecuteSingleResult(r.CollectionName, bson.D{{field, value}})
}

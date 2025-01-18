package mongo

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration constants
const (
	mongoURI       = "mongodb://localhost:27017"
	dbName         = "golang_data"
	collectionName = "wecareit_audit_logs"
)

// Global variables
var (
	mongoClient *mongo.Client
	collection  *mongo.Collection
)

// ConnectMongoDB establishes connection with MongoDB
func ConnectMongoDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	mongoClient = client
	collection = client.Database(dbName).Collection(collectionName)
	return nil
}

// SaveToMongo saves data to MongoDB
func SaveToMongo(data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return fmt.Errorf("failed to insert data to MongoDB: %v", err)
	}

	return nil
}

func GetDataFromMongo(pageStr string) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return nil, fmt.Errorf("invalid page number: %v", err)
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(10 * (page - 1)))
	findOptions.SetLimit(10)

	// Use an empty filter instead of nil
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find data in MongoDB: %v", err)
	}

	var data []interface{}
	if err = cursor.All(ctx, &data); err != nil {
		return nil, fmt.Errorf("failed to decode data from MongoDB: %v", err)
	}

	return data, nil
}

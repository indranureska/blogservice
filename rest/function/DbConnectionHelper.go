package function

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	serviceConst "github.com/indranureska/service/rest/common"
)

/* Used to create a singleton object of MongoDB client.
Initialized and exposed through  GetMongoDbClient().*/
var clientInstance *mongo.Client

// Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

// Used to execute client creation procedure only once.
var mongoOnce sync.Once

// GetMongoClient - Return mongodb connection to work with
func GetMongoDbClient() (*mongo.Client, error) {
	// Perform connection creation operation only once.
	mongoOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Set client options
		clientOptions := options.Client().ApplyURI(serviceConst.BLOG_DB_URI)
		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(ctx, readpref.Primary())
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

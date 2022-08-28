package function

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	serviceConst "github.com/indranureska/service/rest/common"
	"github.com/indranureska/service/rest/utils"
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

		// Default Blog DB URI
		var blogDbURI string

		config, err := utils.LoadConfig("../..")
		if err != nil {
			log.Println("cannot load config: ", err)
			log.Println("using default configuration instead")
			blogDbURI = serviceConst.BLOG_DB_URI
		} else {
			log.Println("config loaded")
			blogDbURI = config.BlogDbURI
		}

		// Set client options
		clientOptions := options.Client().ApplyURI(blogDbURI)
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

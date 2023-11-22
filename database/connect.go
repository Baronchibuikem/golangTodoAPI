package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoInstance : MongoInstance Struct
type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}


// Database : An instance of MongoInstance Struct
var Database MongoInstance

func ConnectDB(){
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect to mongodb
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	// check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	Database = MongoInstance{
		Client: client,
		DB:     client.Database(os.Getenv("DATABASE_NAME")),
	}
}

func CloseDB() {
	if Database.Client != nil {
		err := Database.Client.Disconnect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
	}
}
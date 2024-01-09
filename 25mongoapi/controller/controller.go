package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://supriya01:Ghorai@2001@cluster0.bkmocl0.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {

	clientOption := options.Client().ApplyURI(connectionString)

	// connect with DB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	// collection instance
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collectin instance is ready")
}

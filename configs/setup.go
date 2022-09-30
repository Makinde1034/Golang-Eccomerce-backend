package configs

import (
	"context"
	"fmt"

	"log"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"


)

func ConnectDB()  *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://Makinde1034:Makinde1034@cluster0.hqnk9l6.mongodb.net/?retryWrites=true&w=majority"))

	if err != nil {
		log.Fatal(err)
	}

	// defer func() {
	// 	if err = client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err, "fffff")
	}

	fmt.Print("connected to mdb")

	// databases, err := client.ListDatabaseNames(ctx, bson.M{})

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(databases)
	return client
}

var DB *mongo.Client = ConnectDB()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection =  client.Database("Music-app").Collection(collectionName)
	return collection
}

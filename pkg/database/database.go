package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB(URL string) (*mongo.Collection, error) {
	log.Println("connecting database")
	option := options.Client().ApplyURI(URL)
	option.SetAuth(options.Credential{
		Username: "mongo",
		Password: "mongo",
	})

	client, err := mongo.Connect(context.TODO(), option)

	if err != nil {
		log.Println("failed to connect mongodb: ", err)
		return nil, err
	}

	collection := client.Database("user").Collection("user")
	log.Println("connected to db")
	return collection, nil
}

package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func FindMatch(id string) (IMatch, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	matchesCollection := client.Database("games").Collection("matches")

	var result bson.D
	if err = matchesCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result); err != nil {
		return nil, err
	}
	return bsonToMatch(result), nil
}

func InsertMatch(match IMatch) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	matchesCollection := client.Database("games").Collection("matches")

	_, err = matchesCollection.InsertOne(context.TODO(), match)
	if err != nil {
		return err
	}
	return nil
}

func bsonToMatch(record bson.D) IMatch {
	return QuoridorMatch{}
}

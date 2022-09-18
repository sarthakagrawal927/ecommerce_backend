package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectDatabase() (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	return client, ctx
}

func initDatabase() {
	database = dbClient.Database(DATABASE_NAME)
	userDetailsCollection = database.Collection("users")
	userDetailsCollection.Indexes().CreateOne( // create index on email
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: ""}},
			Options: options.Index().SetUnique(true),
		},
	)
}

func getAllUsersFromDB() []UserDetails {
	defer Recovery()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := userDetailsCollection.Find(ctx, bson.M{"active": true})
	if err != nil {
		log.Fatal(err)
	}
	var users []UserDetails
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result UserDetails
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, result)
	}
	return users
}

func getUserByMail(email string) UserDetails {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user UserDetails
	err := userDetailsCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	handleError(err)
	defer Recovery()
	return user
}

func createUserDB(name string, age string, email string) UserDetails {
	ageInt, err := strconv.ParseInt(age, 10, 16)
	handleError(err)
	user := UserDetails{
		Name:   name,
		Age:    int16(ageInt),
		Email:  email,
		Active: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = userDetailsCollection.InsertOne(ctx, user)
	handleError(err)
	return user
}

func deleteUserByEmailDB(email string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := userDetailsCollection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"active": false}})
	handleError(err)
	return res.UpsertedCount
}

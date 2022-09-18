package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectMongoDatabase() (*mongo.Client, context.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	handleError(err)
	return client, ctx
}

func initDatabase() {
	initMongoDB()
	initPostgresDB()
}

func initPostgresDB() {
	postgresDB.AutoMigrate(&User{})
}

func initMongoDB() {
	mongoDB = mongoDBClient.Database(MONGO_DATABASE_NAME)
	userDetailsCollection = mongoDB.Collection("users")
	userDetailsCollection.Indexes().CreateOne( // create index on email
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: ""}},
			Options: options.Index().SetUnique(true),
		},
	)
}

func connectPostgresDatabase() *gorm.DB {
	connStr := "host=localhost user=sarthak password=12345678 dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	postgresDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	handleError(err)
	return postgresDB
}

func clostPostgresDatabase() {
	dbInstance, err := postgresDB.DB()
	handleError(err)
	dbInstance.Close()
}

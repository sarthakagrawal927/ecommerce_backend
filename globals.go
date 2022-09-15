package main

import "go.mongodb.org/mongo-driver/mongo"

var dbClient *mongo.Client
var database *mongo.Database
var userCollection *mongo.Collection

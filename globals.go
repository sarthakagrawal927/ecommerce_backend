package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var dbClient *mongo.Client
var database *mongo.Database
var userDetailsCollection *mongo.Collection
var logger *zap.Logger

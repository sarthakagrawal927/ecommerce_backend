package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var mongoDBClient *mongo.Client
var mongoDB *mongo.Database
var userDetailsCollection *mongo.Collection
var logger *zap.Logger
var postgresDB *gorm.DB

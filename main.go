package main

import "context"

func main() {
	var ctx context.Context
	mongoDBClient, ctx = connectMongoDatabase()
	postgresDB = connectPostgresDatabase()
	defer mongoDBClient.Disconnect(ctx)
	defer clostPostgresDatabase()
	go initDatabase()
	go initLogger()
	startServer()
}

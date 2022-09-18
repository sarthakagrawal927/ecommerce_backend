package main

import "context"

func main() {
	var ctx context.Context
	dbClient, ctx = connectDatabase()
	defer dbClient.Disconnect(ctx)

	go initDatabase()
	go initLogger()
	startServer()
}

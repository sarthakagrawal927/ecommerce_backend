package main

import (
	"github.com/labstack/echo"
)

func startServer() {
	e := echo.New()
	e.GET("/", testServer)

	// user routes
	e.GET("/users/single", getUser)
	e.POST("/users/create", createUser)
	e.GET("/users/all", getUsers)
	e.POST("/users/delete", deleteUserByEmail)

	e.Logger.Fatal(e.Start(":1323"))
}

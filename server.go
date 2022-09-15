package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func startServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", testServer)

	// user routes
	e.GET("/users/single", getUser)
	e.POST("/users/create", createUser)
	e.GET("/users/all", getUsers)
	e.POST("/users/delete", deleteUserByEmail)

	e.Logger.Fatal(e.Start(":1323"))
}

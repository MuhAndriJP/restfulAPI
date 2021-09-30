package main

import (
	"restfulAPI/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/user/:id", controller.GetUserController)
	e.GET("/users", controller.GetUsersController)
	e.DELETE("/user/:id", controller.DeleteUserController)
	e.PUT("/user/:id", controller.UpdateUserController)
	e.POST("/user", controller.CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}

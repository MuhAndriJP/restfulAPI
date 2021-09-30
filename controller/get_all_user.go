package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get All Users",
		"users":   users,
	})
}

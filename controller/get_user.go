package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get user by id
func GetUserController(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "Invalid ID")
	}
	if id < len(users) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "Succes get user by id",
			"id":       users[id].Id,
			"name":     users[id].Name,
			"email":    users[id].Email,
			"password": users[id].Password,
		})
	}
	return c.String(http.StatusNotFound, "Not Found")
}

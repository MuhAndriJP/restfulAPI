package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// update user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id",
		})
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			if i == len(users)-1 {
				users = users[:len(users)-1]
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "success get all users",
					"users":    users,
				})
			}
			users = users[i+1:]
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get all users",
				"users":    users,
			})
		}

	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "invalid id",
	})
}

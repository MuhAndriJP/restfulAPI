package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func UpdateUserController(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	users[id].Email = u.Email
	users[id].Password = u.Password
	return c.JSON(http.StatusOK, users[id])
}

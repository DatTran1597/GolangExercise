package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type (
	user struct {
		ID   int    `json:"code"`
		Name string `json:"mgs"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)
func createUser(c echo.Context) error {
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("code"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("code"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("code"))
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()
	e.POST("/users", createUser)
	e.GET("/users/:code", getUser)
	e.PUT("/users/:code", updateUser)
	e.DELETE("/users/:code", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
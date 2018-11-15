package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	user struct {
		CODE   int    `json:"code"`
		Mgs string `json:"mgs"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
)
func createUser(c echo.Context) error {
	u := &user{
		CODE: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.CODE] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("code"))
	return c.JSON(http.StatusOK, users[code])
}

func updateUser(c echo.Context) error {
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	code, _ := strconv.Atoi(c.Param("code"))
	users[code].Mgs = u.Mgs
	return c.JSON(http.StatusOK, users[code])
}

func deleteUser(c echo.Context) error {
	code, _ := strconv.Atoi(c.Param("code"))
	delete(users, code)
	return c.NoContent(http.StatusNoContent)
}

func main() {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"host":"${host}","status":"${status}","method":"${method}"}`+"\n",
	}))
	e.POST("/users", createUser)
	e.GET("/users/:code", getUser)
	e.PUT("/users/:code", updateUser)
	e.DELETE("/users/:code", deleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
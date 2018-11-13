package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func show(h echo.Context) error {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	coll := session.DB("mydb").C("demo")

	// Find the number of games won by Dave
	player := "Dave"
	gamesWon, err := coll.Find(bson.M{"winner": player}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", player, gamesWon)
	return h.String(http.StatusOK, player+" has won "+strconv.Itoa(gamesWon)+" games\n")
}
func main() {
	e := echo.New()
	e.GET("/", show)
	e.Start(":8000")
}

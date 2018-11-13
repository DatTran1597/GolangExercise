package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

type Game struct {
	Winner       string    `bson:"winner"`
	OfficialGame bool      `bson:"official_game"`
	Location     string    `bson:"location"`
	StartTime    time.Time `bson:"start"`
	EndTime      time.Time `bson:"end"`
	Players      []Player  `bson:"players"`
}

type Player struct {
	Name   string    `bson:"name"`
	Decks  [2]string `bson:"decks"`
	Points uint8     `bson:"points"`
	Place  uint8     `bson:"place"`
}

func NewPlayer(name, firstDeck, secondDeck string, points, place uint8) Player {
	return Player{
		Name:   name,
		Decks:  [2]string{firstDeck, secondDeck},
		Points: points,
		Place:  place,
	}
}

func main() {
	game := Game{
		Winner:       "Dave",
		OfficialGame: true,
		Location:     "Austin",
		StartTime:    time.Date(2015, time.February, 12, 04, 11, 0, 0, time.UTC),
		EndTime:      time.Date(2015, time.February, 12, 05, 54, 0, 0, time.UTC),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}
	game2 := Game{
		Winner:       "Javi",
		OfficialGame: false,
		Location:     "Ohio",
		StartTime:    time.Date(2014, time.February, 05, 04, 40, 0, 0, time.UTC),
		EndTime:      time.Date(2015, time.September, 23, 24, 54, 0, 0, time.UTC),
		Players: []Player{
			NewPlayer("Dave", "Wizards", "Steampunk", 21, 1),
			NewPlayer("Javier", "Zombies", "Ghosts", 18, 2),
			NewPlayer("George", "Aliens", "Dinosaurs", 17, 3),
			NewPlayer("Seth", "Spies", "Leprechauns", 10, 4),
		},
	}
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Printf("Connected to %v\n", session.LiveServers())

	coll := session.DB("mydb").C("demo")
	if err := coll.Insert(game); err != nil {
		panic(err)
	}
	if err2 := coll.Insert(game2); err2 != nil {
		panic(err2)
	}
	fmt.Println("Document inserted successfully!")
}

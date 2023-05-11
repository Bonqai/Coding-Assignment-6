package main

import (
	"fmt"
	"math/rand"
	"time"
)

type room struct {
	name string
	description string
	exits map[string]*room
}
func rollRandNum(chance int) int {
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(chance)
	return randNum
}

func main() {
	fmt.Println("Go!")

// Setting up each room for the game loop
	room1 := &room{
		name: "Room 1",
		description: "",
		exits: map[string]*room{"north": nil, "east": nil, "south": room2, "west": nil},
	}
	room2 := &room{
		name: "Room 2",
		description: "",
		exits: map[string]*room{"north", room1, "east": room4, "south": room7, "west": room3},
	}
	room3 := &room{
		name: "Room 3",
		description: "",
		exits: map[string]*room{"north": nil, "east": room2, "south": hallway1, "west": nil, "sideroom": room3SideRoom},
	}
	room4 := &room{
		name: "Room 4",
		description: "",
		exits: map[string]*room{"north": nil, "east": nil, "south": hallway2, "west": room2, "sideroom": room6},
	}
	room5 := &room{
		name: "Room 5",
		description: "",
		exits: map[string]*room{"north": nil, "east": room7, "south": nil, "west": hallway1},
	}
	room6 := &room{
		name: "Room 6",
		description: "",
		exits: map[string]*room{"north": room4, "east": nil, "south": nil, "west": nil},
	}
	room7 := &room{
		name: "Room 7",
		description: "",
		exits: map[string]*room{"north": room2, "east": nil, "south": nil, "west": room5},
	}
	room8 := &room{
		name: "Room 8"
		description: "",
		exits: map[string]*room{"north": hallway1, "east": room13, "south": nil, "west": nil},
	}
	room9 := &room{
		name: "Room 9"
		description: "",
		exits: map[string]*room{"north": nil, "east": nil, "south": nil, "west": hallway2},
	}
	room10 := &room{
		name: "Room 10"
		description: "",
		exits: map[string]*room{"north": hallway2, "east": nil, "south": nil, "west": room12},
	}
	room11 := &room{
		name: "Room 11"
		description: "",
		exits: map[string]*room{"north": nil, "east": room12, "south": nil, "west": room13},
	}
	room12 := &room{
		name: "Room 12"
		description: "",
		exits: map[string]*room{"north": nil, "east": room10, "south": nil, "west": room11},
	}
	room13 := &room{
		name: "Room 13"
		description: "",
		exits: map[string]*room{"north": nil, "east": room11, "south": nil, "west": room8},
	}
	hallway1 := &room{
		name: "Hallway 1"
		description: "",
		exits: map[string]*room{"north": room3, "east": room5, "south": room8, "west": nil},
	}
	hallway2 := &room{
		name: "Hallway 2"
		description: "",
		exits: map[string]*room{"north": room4, "east": room9, "south": room10, "west": nil},
	}
	room3SideRoom := &room{
		name: "Room 3 Side Room",
		description: "",
		exits: map[string]*room{"north": nil, "east": nil, "south": nil, "west": room3}
	}
}
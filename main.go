package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

type room struct {
	name string
	description string
	options string
	north *room
	east *room
	south *room
	west *room
	xtraRoom *room
}

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")

	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (r *room) move(direction rune) *room {
	switch direction {
	case 'N', 'n':
		return r.north
	case 'E', 'e':
		return r.east
	case 'S', 's':
		return r.south
	case 'W', 'w':
		return r.west
	default:
		fmt.Println("I'm sorry, please try another direction.")	
		return r
	}
}

func currentOptions(r *room) {
	fmt.Println("----------------------------------------------")
	fmt.Println("You are currently in: ", r.name)
	fmt.Println(r.options)

	if r.north != nil {
		fmt.Println("(N) - North")
	} 
	
	if r.east != nil {
		fmt.Println("(E) - East")
	} 

	if r.south != nil {
		fmt.Println("(S) - South")
	} 

	if r.west != nil {
		fmt.Println("(W) - West")
	}
	fmt.Println("----------------------------------------------")
}

func rollRandNum(chance int) int {
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(chance)
	return randNum
}

func gameLoop() {
	reader := NewCinReader()

	// Initializing each of the rooms 
	r1 := &room {name: "Room 1", description: ""}
	r2 := &room {name: "Room 2", description: ""}
	r3 := &room {name: "Room 3", description: ""}
	r4 := &room {name: "Room 4", description: ""}
	r5 := &room {name: "Room 5", description: ""}
	r6 := &room {name: "Room 6", description: ""}
	r7 := &room {name: "Room 7", description: ""}
	r8 := &room {name: "Room 8", description: ""}
	r9 := &room {name: "Room 9", description: ""}
	r10 := &room {name: "Room 10", description: ""}
	r11 := &room {name: "Room 11", description: ""}
	r12 := &room {name: "Room 12", description: ""}
	r13 := &room {name: "Room 13", description: ""}
	h1 := &room {name: "Hallway 1", description: ""}
	h2 := &room {name: "Hallway 2", description: ""}
	r3sr := &room {name: "Room 3 Side Room", description: ""}

	// Initiate all possible movement options based on the map
	r1.south = r2
	r2.north = r1
	r2.east = r4
	r2.south = r7
	r2.west = r3
	r3.east = r2
	r3.south = h1
	r3.xtraRoom = r3sr
	r4.south = h2
	r4.west = r2
	r4.xtraRoom = r6 
	r5.east = r7
	r5.west = h1
	r6.north = r4
	r7.north = r2
	r7.west = r5
	r8.north = h1
	r8.east = r13
	r9.west = h2
	r10.north = h2
	r10.west = r12
	r11.east = r12
	r11.west = r13
	r12.east = r10
	r12.west = r11
	r13.east = r11
	r13.west = r8
	h1.north = r3
	h1.east = r5
	h1.south = r8
	h2.north = r4
	h2.east = r9
	h2.south = r10
	r3sr.west = r3

	// Create the current room placeholder variable
	currentRoom := r1
	gameEnd := false

// Variables ^^^ ----------------------------------------------------------------------------------------

for !gameEnd {
	ClearScreen()
	currentOptions(currentRoom)

	fmt.Println("Which direction would you like to move?")
	directionMoved := reader.ReadCharacterSet([]rune("NnEeSsWw"))
	currentRoom = currentRoom.move(directionMoved)

	if currentRoom == r11 {
		fmt.Println("You won!")
		gameEnd = true
	
	}
	}
}
func main() {
	gameLoop()
}

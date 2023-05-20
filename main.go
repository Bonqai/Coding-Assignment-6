/*
// AUTHOR: Ian Anderson
// Sources:
// https://freesound.org/
//
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Room struct {
	name string
	description string
	options string
	north *Room
	east *Room
	south *Room
	west *Room
	xtraRoom *Room
}

type Player struct {
	name *string
	health *int
}

type Item struct {
	name string
	description string
}

type Inventory struct {
    items []Item
}

func ClearScreen() {
	// Function to clear the screen on the terminal, very self-explanatory
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")

	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (r *Room) Move(direction rune) *Room {
	// Switch statement I don't wanna write again for telling the computer where the player is going to move
	switch direction {
	case 'N', 'n':
		return r.north
	case 'E', 'e':
		return r.east
	case 'S', 's':
		return r.south
	case 'W', 'w':
		return r.west
	case 'D', 'd':
		return r.xtraRoom
	default:
		fmt.Println("I'm sorry, please try another direction.")	
		return r
	}
}

func currentOptions(r *Room) {
	// Function to display the current options the player has for movement based on the room that they're in
	fmt.Println("\n----------------------------------------------")
	fmt.Println("Which direction would you like to go? Please select the key corresponding to the direction you would like to move and press [ENTER]")
	fmt.Println(r.options)

	if r.north != nil {
		fmt.Println("[N] - North")
	} 
	if r.east != nil {
		fmt.Println("[E] - East")
	}
	if r.south != nil {
		fmt.Println("[S] - South")
	} 
	if r.west != nil {
		fmt.Println("[W] - West")
	}
	if r.xtraRoom != nil {
		fmt.Println("[D] - Side Door")
	}
	fmt.Println("----------------------------------------------")
}

func (i *Inventory) AddItem(item Item) {
	// Function to add an item to the players inventory
    i.items = append(i.items, item)
}

func (inv *Inventory) RemoveItem(itemName string) bool {
	// Function to remove an item from the players inventory
    for i, item := range inv.items {
        if item.name == itemName {
            inv.items = append(inv.items[:i], inv.items[i + 1:]...)
            return true
        }
    }
    return false
}

func CharcaterCreation(name string, health int) Player {
	// Basic return statement for the users character
	return Player {
		name: &name,
		health: &health,
	}
}

func rollRandNum(chance int) int {
	// Random Number Generator for combat scenarios
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(chance)
	return randNum
}

func playAudio(location string, duration int) { 
	// This is a function that opens and audio file, decodes it, and plays it through a streamer
	done := make(chan bool)

	f, err := os.Open(location) 
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	select {
	case <- done:
		streamer.Close()
		return
	case <- time.After(time.Duration(duration) * time.Millisecond):
		speaker.Lock()

		speaker.Unlock()
	}
	streamer.Close()
}

func loopPrint(str string, interval int) {
	// This function loops through the length of a string and prints it out one character at a time
	for _, c := range str {
		fmt.Print(string(c))
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

// Room Functions
func room1() {
	loopPrint("Ahhhhh!", 50)
	fmt.Println()
	time.Sleep(time.Second)
	loopPrint("Oh no!", 50)
	time.Sleep(500 * time.Millisecond)
	loopPrint(" You've fallen down a well!", 50)
	fmt.Println()
	time.Sleep(time.Second)
	loopPrint("You look around to see a dimly lit torch, a bed with a cracked frame, and a pile of wood in the corner", 50)
	time.Sleep(3 * time.Second)
	ClearScreen()
	loopPrint("You hear a mysterious voice to the south of you", 50)
	loopPrint("...", 200)
}


func gameLoop() {
	reader := NewCinReader()

	// Initializing each of the Rooms 
	r1 := &Room {name: "Well"}
	r2 := &Room {name: "Room 2"}
	r3 := &Room {name: "Room 3", description: ""}
	r4 := &Room {name: "Room 4", description: ""}
	r5 := &Room {name: "Room 5", description: ""}
	r6 := &Room {name: "Room 6", description: ""}
	r7 := &Room {name: "Room 7", description: ""}
	r8 := &Room {name: "Room 8", description: ""}
	r9 := &Room {name: "Room 9", description: ""}
	r10 := &Room {name: "Room 10", description: ""}
	r11 := &Room {name: "Room 11", description: ""}
	r12 := &Room {name: "Room 12", description: ""}
	r13 := &Room {name: "Room 13", description: ""}
	h1 := &Room {name: "Hallway 1", description: ""}
	h2 := &Room {name: "Hallway 2", description: ""}
	r3sr := &Room {name: "Room 3 Side Room", description: ""}

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

	// Create the current Room placeholder variable
	gameEnd := false

// Variables ^^^ ----------------------------------------------------------------------------------------

for !gameEnd {
	ClearScreen()
	currentRoom := r1
	
	switch {
	case currentRoom == r1:
		room1()
	}

	currentOptions(currentRoom)
	fmt.Println("Which direction would you like to move?")
	directionMoved := reader.ReadCharacterSet([]rune("NnEeSsWwXx"))
	if directionMoved == 'X' || directionMoved == 'x' {
		exit()
	}
	currentRoom = (*Room).Move(currentRoom, directionMoved)

	if currentRoom == r11 {
		fmt.Println("You won!")
		gameEnd = true
	
	}
	}
}


func exit() {
	fmt.Println("Shutting down...")
	time.Sleep(250 * time.Millisecond)
	os.Exit(0)
}
func main() {
	// reader := NewCinReader()
	// fmt.Print("Enter Name: ")
	// playerName := reader.ReadString(true)
	// p1 := CharcaterCreation(playerName, 250)

	gameLoop()
}

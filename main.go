package main

import (
	"fmt"
)

func room1() {
	// You can go to room 2
}

func room2() {
	// You can go to room 3, 4, 7
}

func room3() {
	// You fight a skeleton wielding an axe and a flail
	// Once defeating the skeleton, there is a chest you can loot that contains a lockpick and a longsword
	// You can go to room 8, or through a trapdoor into a secret room
	// If you go to room 8, you can also go to room 5 through the hallway
}

func room4() {
	// You fight a goblin wielding a one handed sword and no shield
	// You can go to room 6, room 10, and room 10 if you go to room 10 you can room 9 as well through the hallway
}

func room5() {
	// You fight a godlin holding a two handed staff
	// You can either head into room 7, or the hallway towards room 3 or 8.
	// There is a locked chest in the room as well, in which you need a lockpick to open
}

func room6() {
	// There is a goblin holding a one handed sword and armor
	// There is one unlocked chest in the room with a better shield and 2 health potions
	// Room 6 is a room secluded, and can only go back to room 4.
}

func room7() {
	// Room seven is a safe room with no enemies
	// It contains a cupboard in the corner with 1 health potion
	// It leads to room 5 to the west and room 2 to the north.
}

func room8() {
	// Room 8 is a safe room with a bed
	// Leads to room 13 to the east and room 3 to the north
}

func room9() {
	// Room 9 is a trap room containing a chest and a lit campfire
}

func room10() {
	// Room 10 
}

func main() {
	fmt.Println("Go!")
}
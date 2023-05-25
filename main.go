/*
// AUTHOR: Ian Anderson
// Sources:
// https://freesound.org/
//
*/

/*

BUG LIST:
	The View Inventory Function goes back to the start of the game loop, rather than continuing
	You can only play audio files once before the whole system glitches out
	You can get infinite health potions from the chest in R3SR

*/

package main

import (
	"bufio"
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
	complete bool
}


type Entity struct {
	name string
	health int
	damage int
}

type Item struct {
	name string
	description string
	damage int
	field string
	healing int
	used bool
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

func waitForKeyPress() {
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadByte()
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

func (inventory *Inventory) currentOptions(r *Room) {
	// Function to display the current options the player has for movement based on the room that they're in
	if len(inventory.items) > 0 {
		fmt.Println("If you would like to view your inventory, press [I] and then press [ENTER]")
		fmt.Println("\n----------------------------------------------")
		fmt.Println("Which direction would you like to go? Please select the key corresponding to the direction you would like to move and press [ENTER]")
	} else {
		fmt.Println("Which direction would you like to go? Please select the key corresponding to the direction you would like to move and press [ENTER]")
		fmt.Println("\n-------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println(r.options)
	}

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
	fmt.Println("\n-------------------------------------------------------------------------------------------------------------------------------------------")
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

func (inv *Inventory) ViewInventory() int {
	reader :=  NewCinReader()
	ClearScreen()
	if len(inv.items) == 0 {
		fmt.Println("Your inventory is currently empty.")
		fmt.Println()
		fmt.Println("When finished, press [ENTER].")
		waitForKeyPress()
		return 0
	}

	healing := 0
	fmt.Println("Inventory:")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	for _, item := range inv.items {
		fmt.Printf("%s:\n\nType: %s\nDescription: %s", item.name, item.field, item.description)
		if item.healing != 0 {
			fmt.Println("\nWould you like to use this item?")
			yesOrNo := reader.ReadCharacterSet([]rune("YyNn"))
			if (yesOrNo == 'Y') || (yesOrNo == 'y') {
				inv.RemoveItem(item.name)
				healing = item.healing
				item.used = true
			} 
		}
		fmt.Println("\n- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	}
		fmt.Println()
		fmt.Println("When finished, press [ENTER].")
		waitForKeyPress()
		return healing
}

func (i *Inventory) chestPrompt(item Item) {
	fmt.Println("\nWould you like to open the chest?")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	fmt.Println("[Y] - Yes")
	fmt.Println("[N] - No")

	var yesOrNo rune
	fmt.Scanf("%c", &yesOrNo)
	switch {
	case (yesOrNo == 'y') || (yesOrNo == 'Y'):
		ClearScreen()
		i.AddItem(item)
		fmt.Printf("You found a %s!", item.name)
		time.Sleep(1500 * time.Millisecond)
	case (yesOrNo == 'n') || (yesOrNo == 'N'):
	default:
		fmt.Println("I'm sorry, please enter something else")
		time.Sleep(time.Second)
		i.chestPrompt(item)
}
}

func (inv *Inventory) checkInventoryContainsItem(item Item) bool {
	for i := range inv.items {
		if inv.items[i] == item {
			return true
		} 
	}
	return false
}


func rollRandNum(chance int) int {
	// Random Number Generator for combat scenarios
	rand.NewSource(time.Now().UnixNano())
	randNum := rand.Intn(chance) + 1
	return randNum
}


func combat(enemy Entity, userHealth int, userDamage int, enemyDamage int, userWeapon Item) int {
	// Function to handle all possible outcomes of combat, as well as random generation
	// Variable creation
	reader := NewCinReader()
	enemyAttackRandomNumberGen := 0
	userCriticalHit := 0
	userCriticalHitBool := false
	enemyCriticalHit := 0
	enemyCriticalHitBool := false
	userMissChance := 0
	enemyMissChance := 0
	combatEnd := false

	// Start of Combat For Loop
	for !combatEnd {
	ClearScreen()
	
	// Prompt for users action
	fmt.Printf("Your current health: %d", userHealth)
	fmt.Printf("\nThe %s's current health: %d", enemy.name, enemy.health)
	fmt.Println("\n-------------------------------------------------------------------------------------------------")
	fmt.Println("How would you like to attack?")
	fmt.Println("Press [A] to Attack")
	fmt.Println("Press [B] to Block")
	fmt.Println("Press [P] to Parry")
	fmt.Println("-------------------------------------------------------------------------------------------------")
	attackOption := reader.ReadCharacterSet([]rune("AaBbPpXx"))

	// Random number generation for combat outcomes
	userCriticalHit = rollRandNum(10)
		if userCriticalHit == 10 {
			userCriticalHitBool = true
		}
	enemyCriticalHit = rollRandNum(10)
		if enemyCriticalHit == 10 {
			enemyCriticalHitBool = true
		}
	userMissChance = rollRandNum(5)
	enemyMissChance = rollRandNum(5)
	enemyAttackRandomNumberGen = rollRandNum(2)
	ClearScreen()

	switch attackOption {
	case 'A', 'a':

		// Both players attacking
		if enemyAttackRandomNumberGen == 1 {
			if userMissChance != 5 { 						// Player lands their attack
				fmt.Printf("You swung your %s at the %s!", userWeapon.name, enemy.name)
				time.Sleep(time.Second)
				fmt.Println()
				if userCriticalHitBool  {					// Player got a critical hit
					fmt.Println("You got a critical hit!")
					time.Sleep(time.Second)
					fmt.Printf("\nYou attacked the %s for %d damage.", enemy.name, userDamage * 2)
					enemy.health -= userDamage * 2
					time.Sleep(time.Second)
				} else {
					fmt.Printf("\nYou attacked the %s for %d damage.", enemy.name, userDamage)
					enemy.health -= userDamage
					time.Sleep(time.Second)
				}
			} else {										// Player missed their attack
				fmt.Println("You missed your attack!")
				time.Sleep(time.Second)
			}
			if enemyMissChance != 5 {						// Enemy lands their attack
				if enemyCriticalHitBool {					// Enemy got a critical hit
					fmt.Printf("\nThe %s got a critical hit!", enemy.name)
					time.Sleep(time.Second)
					fmt.Printf("\nThe %s attacked you for %d damage.", enemy.name, enemyDamage * 2)
					userHealth -= enemyDamage * 2
					time.Sleep(time.Second)
				} else {
					fmt.Printf("\nThe %s attacked you for %d damage.", enemy.name, enemyDamage)
					userHealth -= enemyDamage
					time.Sleep(time.Second)
				}
			} else {
				fmt.Printf("\nThe %s missed it's attack!", enemy.name)
				time.Sleep(time.Second)
			}

		// Player attacks, Enemy blocks
		} else if enemyAttackRandomNumberGen == 2 {
			if userMissChance != 5 {						// Player lands their attack
				fmt.Printf("You swung your %s at the %s!", userWeapon.name, enemy.name)
				time.Sleep(time.Second)
				fmt.Printf("\nbut the %s blocked it!", enemy.name)
				time.Sleep(time.Second)
			} else {										// Player misses their attack
				fmt.Println("You missed!")
				time.Sleep(time.Second)
				fmt.Printf("\nbut the %s foolishly tried to block your attack", enemy.name)
				time.Sleep(time.Second)
			}
		}
	case 'B', 'b':

		// User Blocks, Enemy Attacks
		if enemyAttackRandomNumberGen == 1 {
			fmt.Printf("The %s attacked you...", enemy.name)
			fmt.Println("...but you blocked it!")
			time.Sleep(time.Second)
		// User Blocks, Enemy Blocks
		} else {
			fmt.Printf("Both you and the %s raise your shields!", enemy.name)
			time.Sleep(time.Second)
			fmt.Println("\nNothing happens!")
			time.Sleep(time.Second)
		}
	case 'P', 'p':
		//  User Parries, Enemy Attacks
		if enemyAttackRandomNumberGen == 1 {
			fmt.Printf("\nThe %s went in for an attack!", enemy.name)
			time.Sleep(time.Second)
			fmt.Println("...but you successfully parried it!")
			time.Sleep(time.Second)
			fmt.Println("You go in for a critical hit!")
			if userMissChance != 5 {
				time.Sleep(time.Second)
				fmt.Printf("\nYou attack for %d damage to the %s!", userDamage * 2, enemy.name)
				enemy.health -= userDamage * 2
				time.Sleep(time.Second)
			} else {
				time.Sleep(time.Second)
				fmt.Println("But you missed your attack!")
				time.Sleep(time.Second)
			}
		// User Parries, Enemy Blocks
		} else {
			if enemyMissChance <= 4 { 
				fmt.Printf("The %s tried to block an attack while you parried", enemy.name)
				time.Sleep(time.Second)
				fmt.Printf("\nThe %s now saw an opportunity and went in for a critical attack!", enemy.name)
				time.Sleep(time.Second)
				fmt.Printf("\nThe %s attacked you for %d damage!", enemy.name, enemyDamage * 2)
				userHealth -= enemyDamage * 2
				time.Sleep(time.Second)
			} else {
				fmt.Printf("The %s tried to block an attack while you parried", enemy.name)
				time.Sleep(time.Second)
				fmt.Printf("\nThe %s now saw an opportunity and went in for a critical attack!", enemy.name)
				time.Sleep(time.Second)
				fmt.Printf("\n...but the %s missed!", enemy.name)
				time.Sleep(time.Second)
			}
		}
	case 'X', 'x':
		exit()
	}
	if userHealth <= 0 {
		fmt.Println("\nYou've died!")
		time.Sleep(2 * time.Second)
		combatEnd = true
	} else if enemy.health <= 0 {
		fmt.Printf("\nYou've defeated the %s!", enemy.name)
		time.Sleep(2 * time.Second)
		combatEnd = true
	}
}
	return userHealth
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

func introduction() {
}

// Room Functions
func room1() {
	fmt.Println("Ahhhhh!")
	fmt.Println()
	fmt.Print("Oh no!")
	fmt.Print(" You've fallen down a well!")
	fmt.Println()
	fmt.Print("You look around to see a dimly lit torch, a bed with a cracked frame, and a pile of wood in the corner")
	ClearScreen()
	fmt.Print("You hear a mysterious voice to the south of you")
	time.Sleep(time.Second)
}

func (i *Inventory) room2() {
	sword := &Item{name: "Sword", description: "A strong straight blade with a tattered handle intended for manual cutting or thrusting.", damage: 25, field: "Weapon"}
	if i.checkInventoryContainsItem(*sword) {
		return
	} else {
		i.chestPrompt(*sword)
	}
}

func room2Description() {
	fmt.Println("You enter a dark room filled with several piles of bones and a dilapidated barrel in the corner")
	fmt.Println("A faint light appears from a torch in the southwest corner of the room")
	fmt.Println("Next to it you see a chest")
	time.Sleep(time.Second)
}

func (inventory *Inventory) room3(playerInstance *Entity, weapon Item) int {
	skeleton := &Entity{name: "Skeleton", health: 50, damage: 10}
	lockpick := &Item{name: "Lockpick", description: "A sharp instrument used to pick locks and enable entry to things such of small chests and doors.", field: "Utility"}
	fmt.Println("You've entered a fight with this skeleton!")
	time.Sleep(time.Second)
	playerInstance.health = combat(*skeleton, playerInstance.health, weapon.damage, skeleton.damage, weapon)
	if playerInstance.health != 0 {
		ClearScreen()
		fmt.Println("You see a chest out of the corner of your eye on the southwest corner of the room.")
		time.Sleep(time.Second)
		inventory.chestPrompt(*lockpick)
	}

	return playerInstance.health
}

func room3Description() bool {
	reader := NewCinReader()
	fmt.Println("You hear the sound of bones clanking together ahead of you as well as a dim glow coming from the room in front of you.")
	fmt.Println("As you walk closer you hear something that resembles a metal chain and a fire.")
	fmt.Println("...")
	ClearScreen()
	fmt.Println("You see a skeleton wielding a chain flail.")
	fmt.Println("Would you like to fight it?")
	fmt.Println("Press [Y] to fight the skeleton, or Press [N] to turn back")
	fmt.Println("-------------------------------------------------------------------------------------------------")

	fightCheck := false
	fightCheckRune := reader.ReadCharacterSet([]rune("YyNn"))
	if (fightCheckRune == 'Y') || (fightCheckRune == 'y') {
		fightCheck = true
	}
	
	return fightCheck
}

func (i *Inventory) r3sr() {
	healthPotion := &Item{name: "Health Potion", description: "A vial of red serum that heals the wounds of those who drink them. [Heals 20 Health]", field: "Healing", healing: 20, used: false}
	lockpick := &Item{name: "Lockpick", description: "A sharp instrument used to pick locks and enable entry to things such of small chests and doors.", field: "Utility"}
	
	ClearScreen()
	if !healthPotion.used {
	fmt.Println("You only a door to the southeast side of the room")
	time.Sleep(time.Second)
	fmt.Println("You see a locked chest to your left as you enter the room?")
	fmt.Println("Do you try to open it?")
	if i.checkInventoryContainsItem(*lockpick) {
		i.chestPrompt(*healthPotion) 
	} else {
		ClearScreen()
		fmt.Println("\nWould you like to open the chest?")
		fmt.Println("-------------------------------------------------------------------------------------------------")
		fmt.Println("[Y] - Yes")
		fmt.Println("[N] - No")

		var yesOrNo rune
		fmt.Scanf("%c", &yesOrNo)
		if (yesOrNo == 'Y') || (yesOrNo == 'y') {
			ClearScreen()
			fmt.Println("You couldn't break the lock")
			time.Sleep(time.Second)
		}
	}
}
}

func (inv *Inventory) room4(playerInstance *Entity, weapon Item) int {
	troll := &Entity{name: "Troll", health: 50, damage: 15}
	largeHealthPotion := &Item{name: "Large Health Potion", description: "A large vial of red serum that heals the wounds of those who drink them. [Heals 50 Health]", field: "Healing", healing: 50, used: false}
	trollSword := &Item{name: "Troll Sword", description: "A sword commonly held by trolls for generations. They've become very hard to come by in recent years."}
	sword := &Item{name: "Sword", description: "A strong straight blade with a tattered handle intended for manual cutting or thrusting.", damage: 25, field: "Weapon"}
	reader := NewCinReader()

	fmt.Println("You walk into the room to see a 10 foot tall troll wielding a large one handed sword.")
	time.Sleep(time.Second)
	playerInstance.health = combat(*troll, playerInstance.health, weapon.damage, troll.damage, weapon)
	if playerInstance.health != 0 {
		ClearScreen()
		fmt.Println("Would you like to pick up his sword?")
		fmt.Println("Enter [Y] for Yes, and [N] for No, then press [ENTER]")
		yesOrNo := reader.ReadCharacterSet([]rune("YyNn"))
		if (yesOrNo == 'Y') || (yesOrNo == 'y') {
			inv.AddItem(*trollSword)
			inv.RemoveItem(sword.name)
			fmt.Printf("\nYou dropped your %s in favor of the %s.", weapon.name, trollSword.name)
			time.Sleep(1500 * time.Millisecond)
		}
		if !inv.checkInventoryContainsItem(*largeHealthPotion) {
			inv.chestPrompt(*largeHealthPotion)
		}
	}
	return playerInstance.health
}

func room4Description() bool {
	reader := NewCinReader()
	ClearScreen()

	fmt.Println("You walk along a dark hallway and see a glow at the end")
	time.Sleep(1500 * time.Millisecond)
	fmt.Println("As you approach you start to hear the groans of something taking a deep breath")
	time.Sleep(2 * time.Second)
	fmt.Println("Do you dare to enter?")
	fmt.Println("Press [Y] for Yes or [N] for No, then press [ENTER].")
	yesOrNo := reader.ReadCharacterSet([]rune("YyNn"))
	switch yesOrNo {
	case 'Y', 'y':
		return true
	case 'N', 'n':
		return false
	default:
		fmt.Println("Please enter something else.")
		room4Description()
	}
	return false
}

func gameLoop() {
	reader := NewCinReader()

	// Initializing each of the Rooms 
	r1 := &Room {name: "Well Entrance"}
	r2 := &Room {name: "South of Well Entrance"}
	r3 := &Room {name: "Northwest Corner"}
	r4 := &Room {name: "Northeast Corner"}
	r5 := &Room {name: "Room 5"}
	r6 := &Room {name: "Room 6"}
	r7 := &Room {name: "Room 7"}
	r8 := &Room {name: "Room 8"}
	r9 := &Room {name: "Room 9"}
	r10 := &Room {name: "Room 10"}
	r11 := &Room {name: "Room 11"}
	r12 := &Room {name: "Room 12"}
	r13 := &Room {name: "Room 13"}
	h1 := &Room {name: "Hallway 1"}
	h2 := &Room {name: "Hallway 2"}
	r3sr := &Room {name: "Room 3 Side Room"}

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
	currentRoom := r1
	lastRoom := r1
	inventory := &Inventory{}
	var equippedWeapon Item


	player := &Entity{name: "Player", health: 100, damage: 10}
	sword := &Item{name: "Sword", description: "A strong straight blade with a tattered handle intended for manual cutting or thrusting.", damage: 25, field: "Weapon"}
	healthPotion := &Item{name: "Health Potion", description: "A vial of red serum that heals the wounds of those who drink them. [Heals 20 Health]", field: "Healing", healing: 20, used: false}
	// trollSword := &Item{name: "Troll Sword", description: "A sword commonly held by trolls for generations. They've become very hard to come by since the surface extinction of trolls."}
	// lockpick := &Item{name: "Lockpick", description: "A sharp instrument used to pick locks and enable entry to things such of small chests and doors.", field: "Utility"}



// Variables ^^^ ----------------------------------------------------------------------------------------


for !gameEnd {
	ClearScreen()

	if len(inventory.items) > 0 {
		equippedWeapon = inventory.items[0]
	}
		
	if !currentRoom.complete {
	switch currentRoom {
	case r1:
		room1()
		r1.complete = true
	case r2:
		if !inventory.checkInventoryContainsItem(*sword) {
			room2Description()
			inventory.room2()
			r2.complete = true
		}
	case r3:
		fightCheck := room3Description()
		if fightCheck {
			player.health = inventory.room3(player, equippedWeapon)
			if player.health == 0 {
				gameEnd = true
			}
			r3.complete = true
		} else {
			currentRoom = lastRoom
		}
	case r3sr:
		if !inventory.checkInventoryContainsItem(*healthPotion) {
			inventory.r3sr()
			r3sr.complete = true
		}
	case r4:
		tOrF := room4Description()
		if tOrF {
			player.health = inventory.room4(player, equippedWeapon)
			r4.complete = true
		} else {
			currentRoom = lastRoom
		}
	}
	}


	ClearScreen()
	// Displaying the current health of the player
	fmt.Printf("You are currently in: %s", currentRoom.name)
	fmt.Printf("Health: %d", player.health)
	fmt.Println()
	fmt.Println("\n-------------------------------------------------------------------------------------------------")


	var directionMoved rune
	// Displaying current options for movement and inventory usage
	inventory.currentOptions(currentRoom)
	fmt.Println()
	fmt.Println("Which direction would you like to move?")

	// Checking to see if the player has anything in their inventory
	if len(inventory.items) > 0 {
		directionMoved = reader.ReadCharacterSet([]rune("NnEeSsWwDdXxIi"))
		if directionMoved == 'X' || directionMoved == 'x' {
			exit()
		}
		if directionMoved == 'I' || directionMoved == 'i' {
			player.health += inventory.ViewInventory()
		}
	} else {
		directionMoved = reader.ReadCharacterSet([]rune("NnEeSsWwDdXx"))
		if directionMoved == 'X' || directionMoved == 'x' {
			exit()
		}
	}

	// Telling the computer which room to move the player into
	lastRoom = currentRoom
	currentRoom = (*Room).Move(currentRoom, directionMoved)

	// Create a winning situation for the player
	if currentRoom == r11 {
		ClearScreen()
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
	gameLoop()
}
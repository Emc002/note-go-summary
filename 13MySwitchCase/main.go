package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Switch Case in GO languages")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Number is 1 and you can open")
	case 2:
		fmt.Println("you can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("you can move 3 spot")
		fallthrough
	case 4:
		fmt.Println("you can move 4 spot")
	case 5:
		fmt.Println("you can move 5 spot")
	case 6:
		fmt.Println("LUCKY you, you can move 6 spot roll dice again")
	default:
		fmt.Println("you can't open")
	}
}

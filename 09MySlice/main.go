package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slice in GOLANG")

	var carList = []string{"Morchelago", "Lancer", "Gallardo"}
	fmt.Printf("type is: %T\n", carList)

	carList = append(carList, "CarreraGT")
	fmt.Println(carList)
	carList = append(carList[1:3])
	fmt.Println(carList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 223
	highScores[2] = 323
	highScores[3] = 423
	// highScores[4] = 523

	highScores = append(highScores, 111, 222, 333)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slice based on index

	var roomNumber = []string{"one", "two", "three", "four", "five"}
	fmt.Println(roomNumber)
	var index int = 2
	roomNumber = append(roomNumber[:index], roomNumber[index+1:]...)
	fmt.Println(roomNumber)

}

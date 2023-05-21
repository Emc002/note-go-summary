package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Rating in our movie")

	// comma ok || err ok syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for watching, ", input)
	fmt.Printf("the Type, %T ", input)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome Welcome")
	fmt.Println("Thanks and give me the rating")

	readers := bufio.NewReader(os.Stdin)

	input, _ := readers.ReadString('\n')

	fmt.Println("Thanks for the rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("added 1:  ", numRating+1)
	}
}

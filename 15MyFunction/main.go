package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Function in GOLANG")
	greeter()
	greeterTwo()
	result := add(3, 3)
	fmt.Println("Result is: ", result)
	resultPro := proAdd(7, 7, 9, 8, 6)
	fmt.Println("Result from proAdd is: ", resultPro)
	numberBack, messageBack := proAddTwo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Result from proAddTwo is: ", numberBack)
	fmt.Println("and the message: ", messageBack)

}
func greeterTwo() {
	fmt.Println("Another Method")
}

func greeter() {
	fmt.Println("Hello World")
}

func add(a int, b int) int {
	c := a + b
	return c
}

func proAdd(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}
	return total
}

func proAddTwo(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result Is SUcces"
}

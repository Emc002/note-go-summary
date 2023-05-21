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

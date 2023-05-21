package main

import "fmt"

func main() {
	fmt.Println("Welcome to IF and ELSE in GOLANG")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "guest"
	} else if loginCount > 10 {
		result = "admin"
	} else {
		result = " so the loggin count is 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := -11; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}
}

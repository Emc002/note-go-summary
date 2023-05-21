package main

import "fmt"

func main() {
	fmt.Println("Welcome Welcome")

	// var ptr *int
	// fmt.Println("value of ptr : ", ptr)

	myNumber := 99

	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)

	var nana = &myNumber

	test := *nana + 2

	fmt.Println(test)
	fmt.Println(&test)

	fmt.Println(myNumber)
	fmt.Println(nana)

}

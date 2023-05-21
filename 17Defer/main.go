package main

import "fmt"

func main() {
	// defer is like stacking {LIFO} something and whatever the
	// first that init with the defer its gonna be in
	// bottom of the stacking like we safe our plate
	// (start from the bottom)
	// and when it come to use that we actualy use
	// the plate that we save
	// the last in other meaining is the plate in
	// a top of stacking
	// so that whats defer behave (like stack not queueu)
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Helo")
	myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

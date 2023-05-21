package main

import "fmt"

func main() {
	fmt.Println("Welcome to LOOPS in GO languages")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougeValue := 1

	for rougeValue < 10 {

		// if rougeValue == 5 {
		// 	break
		// }

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeOnline.in")
}

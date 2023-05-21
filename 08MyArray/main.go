package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in GO languages")

	var carList [5]string

	carList[0] = "Gallardo"
	carList[1] = "Morchelago"
	carList[3] = "DB9"
	carList[4] = "Lancer"

	fmt.Println(carList)
	fmt.Println(len(carList))

	var laptopList = [5]string{"ideapad", "msi", "acer", "lenovo", "asus rog"}
	fmt.Println(laptopList)
	fmt.Println(len(laptopList))

}

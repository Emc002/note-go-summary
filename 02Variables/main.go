package main

import "fmt"

//var JwtToken = 3000

// public
const LoginToken = "efko4ij32490fdij"

func main() {
	var username string = "ALIFNAFISALGHANY"
	fmt.Println(username)
	fmt.Printf("type of Varibale: %T \n", username)

	var isLoggin bool = true
	fmt.Println(isLoggin)
	fmt.Printf("type of Varibale: %T \n", isLoggin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("type of Varibale: %T \n", smallVal)

	var smallFLoat float32 = 255.7654678976543567
	fmt.Println(smallFLoat)
	fmt.Printf("type of Varibale: %T \n", smallFLoat)

	// some default values and some aliases
	var anotherVariable bool
	fmt.Println(anotherVariable)
	fmt.Printf("type of Varibale: %T \n", anotherVariable)

	// implicit type
	var website = 1233334455557767857
	fmt.Println(website)
	fmt.Printf("type of Varibale: %T \n", website)

	// no var style
	numberOfUser := 1234325345
	fmt.Println(numberOfUser)
	fmt.Printf("type of Varibale: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("type of Varibale: %T \n", LoginToken)

}

package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list all of languages: ", languages)
	fmt.Println("list of languages: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("list all of languages: ", languages)

	//Loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

	for _, value := range languages {
		fmt.Printf("Value is %v\n", value)
	}

}

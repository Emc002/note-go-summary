package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to Files in golang")

	content := "This needs to go in a file - learnEverting.in"

	file, err := os.Create("./myfile.txt")

	// if err != nil {
	// 	panic(err)
	// }

	errHandler(err)

	length, err := io.WriteString(file, content)

	errHandler(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./myfile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	errHandler(err)
	fmt.Println("Text data inside the file is: \n", string(dataByte))
}

func errHandler(err error) {
	if err != nil {
		panic(err)
	}
}

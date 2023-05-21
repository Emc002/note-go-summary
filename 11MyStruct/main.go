package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to My Struct in GOlang")
	//no inheritance in GOLANG and no super or parent or child in the GOLANG
	alifnafis := User{"ALIF", "alifnafisalghany@gmail.com", true, 21}
	fmt.Println(alifnafis)
	fmt.Printf("alif nafis details are: %+v\n", alifnafis)
	fmt.Printf("Name is %v and email is %v", alifnafis.Name, alifnafis.Email)

}

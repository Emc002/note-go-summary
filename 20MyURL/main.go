package main

import (
	"fmt"
	"net/url"
)

const MyURL = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to handling URL in GOLANG")

	result, _ := url.Parse(MyURL)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Fragment)
	fmt.Println(result.Path)
	fmt.Println(result.RawFragment)
	fmt.Println(result.RawPath)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Println(result.Query())

	qparams := result.Query()

	fmt.Println(qparams["list"])

	for _, val := range qparams {
		fmt.Println("Value of params: ", val)
	}

	partsofURL := &url.URL{
		Scheme:  "https",
		Host:    "www.youtube.com",
		Path:    "/watch",
		RawPath: "user=alif",
	}

	checkresult := partsofURL.String()

	fmt.Println(checkresult)

}

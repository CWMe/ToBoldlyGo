package main

import "fmt"

func main() {
	var stringA = "first var"
	var stringB = "second var"
	fmt.Println(stringA,stringB)

	var intA = 20
	var intB int = 25
	fmt.Println(intA,intB)

	var boolA, boolB bool = true, false
	fmt.Println(boolA,boolB)

	stringC := "third string"
	fmt.Println(stringC)
}
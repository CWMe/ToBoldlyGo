package main

import "fmt"

func main (){
	var stringA = "8"
	for i := 1; i <= 10; i++ {
		stringA += "="
		fmt.Println(stringA + "D")
    }
}
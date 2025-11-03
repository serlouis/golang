package main

import "fmt"

func main(){

	message := "Hello World"

	for i, v := range message {
		//fmt.Println(i,v) // returns the index and unicode value of "Hello World"
		fmt.Printf("Index: %d, Rune: %c\n", i, v) // this will print out the char values
	}
	// range operates on a copy so it won't change the values inside the loop

	for _, v := range message {
		fmt.Printf(" Rune: %c\n", v)
	}

	
}
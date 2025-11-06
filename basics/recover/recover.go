package main

import "fmt"

func main() {

	process()
	fmt.Println("Returned from process")

}

func process() {
	defer func() { // this is an anonymous function its being called inline which confused the hell out of me at first but its defining a function and calling it in the same block
		if r := recover(); r != nil { // if recover happens then r is not equal nil thus a panic happened and we return the recovered value?
			fmt.Println("Recovered", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}

package main

import "fmt"

func main(){

	// panic(interface()) you can input any value of any type as an argument for this function
	// interface is like type any. Lets you use any type of value

	// example of a valid input
	process(10)

	// example of invalid input
	process(-3)

}

func process(input int) {

	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number") // will only run after the deferred functions
		//fmt.Println("After Panic") // this will not be executed
		//defer fmt.Println("Deferred 3") // still unreachable
	}
	fmt.Println("Processing input: ", input)
}
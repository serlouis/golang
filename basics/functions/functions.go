package main

import "fmt"

func main(){

	// func <name>(parameters list)(returnType) {}
	// code block
	// return value

	// public functions start with capital letter
	// private functions start with lowercase letter
	// for example fmt.Println() starts with uppercase

	/*sum := add(1,2)
	fmt.Println(sum)

	fmt.Println(add(2,3))

	// anonymous function
	greet := func(){
		fmt.Println("Hello Anonymous function")
	}

	greet()

	operation := add

	result := operation(3,5)
	fmt.Println(result)
	*/

	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// return and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))
}
// first class objects refers to entities that have no restrictiosn on their use and can be treated uniformly throughout the language

func add(a,b int)(int){
	return a + b
}

// FUnction that takes a function as an argument
func applyOperation(x int, y int, operation func(int,int) int) int {
	return operation(x,y)
}

// Function that returns a function

func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
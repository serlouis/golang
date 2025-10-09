package main

import "fmt"

func main() {
	// variable declaration
	var a, b int = 10, 3
	var result int

	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplation:", result)

	result = a / b // this is going to return an int result cast to a float to get accurate result
	fmt.Println("Division:", result)

	result = a % b
	fmt.Println("Modulus:", result)

	const p float64 = 22 / 7.0 // one of the operands is in float type so it returns a float type. Even if p is declared as a float but if you still use integers it returns integer
	fmt.Println(p)
}

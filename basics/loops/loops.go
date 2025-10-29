package main

import "fmt"

func main() {
	// simple iteration over a range
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// iterate over collection
	numbers := []int {1,2,3,4,5,6}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %v\n", index, value) //%d is specific to numbers %v is generic values
	}

	for j:=1; j<=10; j++ {
		if j%2 == 0 {
			continue // continue the loop but skip the rest of the lines/statements
		}
		fmt.Println("Odd Number: ", j)
		if j == 5 {
			break // break out of the loop
		}
	}

	rows := 5

	// Outer loop
	for k:=1; k<=rows; k++{
		// inner loop for spaces before stars
		for l:=1; l<=rows-k; l++{
			fmt.Print(" ")
		}

		// inner loop for stars
		for m:=1; m<=2*k-1; m++ {
			fmt.Print("*")
		}
		fmt.Println() // move to the next line
	}

	for u:= range 10{
		fmt.Println(10-u)
	}
	fmt.Println("We have a lift off!")

	
}

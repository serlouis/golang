package main

import "fmt"

func main() {

	/*
	fruit := "pineapple"

	switch fruit {
	case "apple":
		fmt.Println("It's an apple.")
	case "banana":
		fmt.Println("It's a banana")
	default:
		fmt.Println("Unknown fruit")
	}

	// Multiple Conditions
	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It is a weekday")
	case "Sunday":
		fmt.Println("Its a weekend day")
	default:
		fmt.Println("invalid day.")
	}

	// Expressions in case statements
	number := 15
	
	switch {
	case number < 10:
		fmt.Println("Number is less than 10")
	case number >= 10 && number <20:
		fmt.Println("Number is between 10 and 19")
	default:
		fmt.Println("Number is 20 or more")
	}

	// fallthrough example

	num := 2

	switch {
	case num > 1:
		fmt.Println("Greater than 1") // this condition is true
		fallthrough
	case num == 2:
		fmt.Println("Number is equal to 2") // this is also true so we need to fallthrough to evaluate both conditions we know to be true
	default:
		fmt.Println("Not 2")
	}
	*/

	for i := 1; i <= 100; i++ {
		switch {
		case i % 5 == 0 && i % 3 == 0:
			fmt.Println("fizzbuzz")
		case i % 5 == 0:
			fmt.Println("buzz")
		case i % 3 == 0:
			fmt.Println("fizz")
		default:
			fmt.Println(i)
		}
	}
}

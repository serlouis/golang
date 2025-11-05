package main

import (
	"errors"
	"fmt"
)

func main(){

	//func funtionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2,....){
	// code block
	// return returnValue1, returnValue2,...
	//}

	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r)

	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a/b // don't use gopher here because gopher is only for new variable delcarations, we named them in the return type
	remainder = a%b
	return // this will return both quotient and remainder because we named them as return parameters
}

// how to send error as a return value

func compare(a, b int) (string, error){
	if ( a > b ){
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
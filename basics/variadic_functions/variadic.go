package main

import "fmt"

func main(){

	// func functionName(param1 type1, param2 type2, param3 ...type3) returnType {} variadic function can accept 0 or more paraemters of a certain type

	// variadic function in action
	//fmt.Println("sum of 1, 2, 3: ", sum(1,2,3))
	statement, total := sum("The sum of 1, 2, 3 is:", 1,2,3)
	fmt.Println(statement, total)

	// create a slice and pass it as variadric parameter
	numbers := []int{1,2,3,4,5,9}

	sequence3, total3 := sum("3", numbers...)
	fmt.Println("Sequence: ", sequence3, "Total", total3)
}
// the variadic parameter must always come last in the function declaration
func sum(returnString string, nums ...int) (string,int) {
	total := 0
	for _, v := range nums {
	total += v
	}
	return returnString, total
}
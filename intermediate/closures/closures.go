package main

import "fmt"

func main(){

	//sequence := adder() // this is only called once so that is why i is initialized to zero

	// we return an anonymous function that returns the updated value of i when we call it
	//fmt.Println(sequence()) // i is going to be updated with every run of sequence
	//fmt.Println(sequence()) // i is scoped to the adder function not sequence
	//fmt.Println(sequence()) // the return function is stored in the sequence value  so its not getting initialized again
	//fmt.Println(sequence())

	//sequence2 := adder() // reinitializes i to zero
	//fmt.Println(sequence2()) 
	
	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int{
			countdown -= x
			return countdown
		}
	}()

	// using the closure subtracter
	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))
	fmt.Println(subtracter(4))
	fmt.Println(subtracter(5))
}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
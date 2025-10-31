package main

import "fmt"

func main(){

	//var arrayName [size]elementTYpe

	var numbers[5]int
	// initialized to zero

	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third Element", fruits[2])

/*	originalArray := [3]int{1,2,3}
	copiedArray := originalArray

	copiedArray[0] = 100 // this won't effect the original array

	fmt.Println("Original Array", originalArray)
	fmt.Println("Copied Array", copiedArray)
*/

	originalArray := [3]int{1,2,3}
	var copiedArray *[3]int

	// this is going to update the original array because we're modifying the address of originalArray instead of a copy
	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Original array", originalArray)
	fmt.Println("Copied array: ", *copiedArray)

	// can determien the length of the array with len function
	for i :=0; i < len(numbers); i++ {
		fmt.Println("ELement at index", i, ":", numbers[i])
	}

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// can accomplish the same thing using the blank identifier, used to store unused values
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	fmt.Println("The length of numbers array is", len(numbers))

	// comparing arrays
	array1 := [3]int{1,2,3}
	array2 := [3]int{10,2,3}

	fmt.Println("Array 1 is equal to array2", array1 == array2)

	// multi dimensional array
	var matrix [3][3]int = [3][3] int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(matrix)

	// someFunction has 2 return values but if we only want to use one we need to use the blank identifier to discard the second return value
	a, _ := someFunction()
	fmt.Println(a)
	//fmt.Println(b)

	b := 2
	_ =  b // if not using the variable immediately we can get rid of the error by assigning it the blank identifier
}

func someFunction() (int, int) {
	return 1, 2
}
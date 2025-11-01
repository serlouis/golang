package main

import (
	"fmt"
	"slices"
)

func main(){
	// declare a slice, a series of numbers (array without fixed length)
	//var numbers []int
	//var numbers1 = []int{1,2,3}

	//numbers2 := []int{9.8,7}

	// these slices are references to underlying arrays but provide a window into the arrays elements. Slices can grow and shrink dynamically

	//slice := make([]int,5)

	a := [5]int{1,2,3,4,5}

	slice1 := a[1:4] // this will not include the element at index 4 {2,3,4} will be included in the slice
	fmt.Println(slice1)

	slice1 = append(slice1, 6,7)
	fmt.Println("Slice1:", slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Slice Copy: ",sliceCopy)

	//var nilSlice []int // does not reference any array its nil

	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	fmt.Println("Element at index 3 of slice1", slice1[3])

	//slice1[3] = 50
	//fmt.Println("Element at index 3 of slice1", slice1[3])

	if slices.Equal(slice1, sliceCopy) {
		fmt.Println("slice1 is equal to sliceCopy")
	}

	// this will make a 2d slice that is initially set to zero but the second loop populates the first slice with sub slices
	twoD := make([][]int, 3)
	for i  := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
			fmt.Printf("Adding value %d in outer slice at index %d, and inner slice index of %d\n", i+j, i, j)
		}
	}

	fmt.Println(twoD)
	
	// slice[low:high]
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	fmt.Println("The capacity of slice2 is", cap(slice2))
	fmt.Println("The capacity of slice2 is", len(slice2))

}
package main

import "fmt"

func main(){

	// var ptr *int creates a pointer to the type int

	var a int = 10
	var ptr *int
	ptr = &a // reference

	fmt.Println(a) // value of a
	fmt.Println(ptr) // memory address of a
	fmt.Println(&a) // should be the same as ptr
	fmt.Println(*ptr) // this will dereference the pointer and return the value of what ptr is pointing to which is a

	// the zero value of a pointer is nil used to indicate the pointer is not valid or initialized

	var ptr1 *int // unitialized pointer should be nil
	if ptr1 == nil {
		fmt.Println("Pointer is nil")
	}

	// go does not support pointer arithmatic limited to referencing and dereferencing 

	// we can pass pointers to functions

	modifyValue(ptr)
	fmt.Println(a)
}

func modifyValue(ptr *int) { // normally a copy is passed but since we pass in a pointer we can perform the operation on the actual object
	*ptr++

	// can't do &ptr++ because go doesn't allow pointer arithmatic like c/c++
}
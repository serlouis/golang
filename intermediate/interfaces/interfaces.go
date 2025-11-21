package main

import (
	"fmt"
	"math"
)

func main() {

	// create a rectanble with width of 3 and height of 4
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//r1 := rect1{width: 3, height: 4}

	measure(&r)
	measure(&c)
	//measure(&r1) // works after rect1 implements the methods of geometry interface

	myPrinter(1, "John", 45.9, true)
	printType(9)
	printType("John")
	printType(false)
}

// declare an interface with 2 functions
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (r rect) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// circle has an additional function that is not part of the interface
func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

// my explanation of whats happening
// the gemotry interface declares 2 method signatures, not implementations
// both rect and circle implement those methods. Matching methods = implementation of the interface, the struct doesn't necessarily have to implment an interface
// the measure function accepts anything that satisfies geometry. In this case both rect and circle do because they impleement both interface methods

// create a new rect struct not an instance but a new one
type rect1 struct {
	width, height float64
}

func (r rect1) area() float64 {
	return r.height * r.width
}

//func (r rect1) perim() float64 {
//	return 2 * (r.height + r.width)
//}

// consider interfaces like a contract. If you want to be considered this "type" then you need to have these methods
// in this example the geometry interface requires  area and perim methods with the same return values as definedh
// this enables polymorphism you can define similar objects but implement different behavior
// every geometric shape has an area, perimiter etc. but they have different ways of deriving those values


// another use of interface. when we want to accept any type of value in our function
// we use a variadic parameter so we can accept infinite values
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

// this function is ready to accept any type of value but it can only accept 1
func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Unkonwn type")
	}
}
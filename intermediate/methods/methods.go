package main

import "fmt"

func main() {

	rect := Rectangle{length: 10, width: 9}
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is: ", area)

	rect.Scale(2)
	area = rect.Area() // have to use assignment operator, this was throwing me off when I tried to use := it already existed though so just reassign it
	fmt.Println("Area of rect scaled by factor of 2 is now: ", area)

	// since we're using a custom receiver type we have to create instances of the object first
	num := MyInt(-5)
	num1 := MyInt(9)

	// now that we have an instance of the object MyInt we can now call functions
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())
	fmt.Println(num.welcomeMessage())

	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	// because we have embedded a struct within a struct the inner struct will be promoted to the outer struct
	// Rectangle's methods become available to shape in other words
	fmt.Println(s.Area())

	// this works as well but easier to just access it directly
	fmt.Println(s.Rectangle.Area())
}

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver use when you need to modify the underlying struct or avoid copying large struct to function
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

// method on user defined type
// MyInt is the receiver type most of the time this is an object like a struct but it can be any user defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

// this time we didn't create an instance but we can associate the type with the method of that type
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}

// promotion of struct
type Shape struct {
	Rectangle
}

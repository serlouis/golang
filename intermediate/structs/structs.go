package main

import "fmt"

func main() {

	// structs can be initilized with a struct literal

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K",
		},
		PHoneHomeCell: PHoneHomeCell{
			home: "5555555555",
			cell: "4444444444",
		},
	}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	//p1.address.city = "New York"
	//p1.address.country = "USA"

	// fields are accessed with dot notation

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p.fullName())
	fmt.Println(p.address)
	fmt.Println(p1.address.city)
	fmt.Println(p1.address.country)
	fmt.Println(p.cell)

	// anonymous structs don't have a pre-defined type name useful for tmep data structures

	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}

	fmt.Println(user.username)

	// go supports attaching methods to structs

	// method in action

	// pointer receiver
	fmt.Println("Before increment", p.age)
	p.incrementAgeByOne()
	fmt.Println("After increment", p.age)

	// structs can be compared
	fmt.Println("Are p and p1 equal?", p == p1) // they are not equal

	fmt.Println("Are p1 and p2 equal?", p1 == p2) // they should be equal?

}

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PHoneHomeCell // this is an embedded struct looks similar but address is a field the embedded struct can be accessed directly
}

type PHoneHomeCell struct {
	home string
	cell string
}

type Address struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++ // this works because we're sending a pointer to the struct

	// if you pass it by value it copies the struct and then after the modify the struct is out of scope and the original is unmodified
}

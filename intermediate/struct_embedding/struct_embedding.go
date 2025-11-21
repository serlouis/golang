package main

import "fmt"

func main() {

	emp := Employee{
		person: person{name: "John", age: 30},
		empId:  "E001", salary: 50000,
	}

	fmt.Println("Name:", emp.name) // accessing the embedded struct field emp.person.name directly without having to go through the person struct
	fmt.Println("Age", emp.age)    // same as above
	fmt.Println("Employee ID:", emp.empId)
	fmt.Println("Salary: ", emp.salary)

	emp.introduce()

}

// not exporting the struct or fields to other packages
type person struct {
	name string
	age  int
}

type Employee struct {
	person // anonymous field
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi, I am %s and I am %d years old\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi, I am %s, with employee ID %s and I earn %.2f\n", e.name, e.empId, e.salary)
}

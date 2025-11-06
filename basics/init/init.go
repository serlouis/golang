package main

import "fmt"

// init functions are called sequentialy in the order they are created

// practical use case
// setup tasks
// configuration
// registering components
// database initialization

func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {

	fmt.Println("Inside the main function")

}

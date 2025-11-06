package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Deferred statement") // this shouldn't run

	// os.exit stops the program immediately without any cleanup/deferred functions
	fmt.Println("Starting the main function")

	// Exit with status code of 1
	os.Exit(1)

	// This will never be executed
	fmt.Println("End of main fucntion")

}

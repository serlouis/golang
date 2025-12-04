package main

import "fmt"

func main() {

	num := 42454564
	fmt.Printf("%05d\n", num) // adds zeros if the digit doesn't have 5 digits

	message := "Hello"
	// fix the minimum width to 10
	fmt.Printf("|%10s|\n", message) // right aligned output use pipes to see the alignment of the spaces

	// left aligned minimum width is 10
	fmt.Printf("|%-10s|\n", message)

	// string interpolation
	// if we use ` it makes a string literal` so for example if you add a \n it doesn't use the escape sequence it just prints \n
	message1 := "Hello \nWorld"
	message2 := `Hello \nWorld` // string literal will print the \n escape sequence
	fmt.Println(message1)
	fmt.Println(message2)

	// string literal is useful for situations like sql query
	//sqlQuery := `SELECT * FROM users WHERE age > 30`
}

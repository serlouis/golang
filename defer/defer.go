package main

import "fmt"

func main() {

	process(10)

}

func process(i int) {
	defer fmt.Println("Deferred i value: ", i)             // this will print 10 because this defer is encountered before we increment i
	defer fmt.Println("First Deferred statement executed") // this will be deferred until the end of the function they are processed in LIFO order
	defer fmt.Println("Second deferred statement excecuted")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i: ", i)
}

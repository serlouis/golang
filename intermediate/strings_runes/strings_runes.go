package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!" // \t tab character
	message2 := "Hello, \rGo!" // carriage return 
	rawMessage := `Hello\nGo` // these escape sequences wont work with ``

	fmt.Println(message)
	fmt.Println(rawMessage)
	fmt.Println(message1)
	fmt.Println(message2)

	// strings are made up of runes in Go an array of characters essentially
	fmt.Println("Lenght of message variable is: ", len(message))
	fmt.Println("Length of raw message variable is: ", len(rawMessage))

	// accessing a character in a string by index zero based
	fmt.Println("The first character in message var is ", message[0]) // ASCII Value of H

	// can also do string concatenation
	greeting := "Hello " // have to manually add a space if needed
	name := "Alice"
	fmt.Println(greeting + name)

	// strings can be compared

	str1 := "Apple" // A has an ASCII value of 65
	str := "apple" // ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app" // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1) // the compiler is comparing ASCII value of the characters 97 > 65 so this is false
	fmt.Println(str > str1) // true
	fmt.Println(str > str3) // true compiler checks app vs app but str has le to evaluate

	// string iteration just like iterating over a slice or an array it has index, value

	for _, char := range message {
		//fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%x\n", char) // will print the hexidecimal value of each character
		fmt.Printf("%v\n", char) // will give the ASCII value
		fmt.Printf("%c\n", char) // will print the rune
	}

	// count the number of runes in the string
	fmt.Println("Rune count: ", utf8.RuneCountInString(greeting))

	// strings are immutable

	greetingWithName := greeting + name // have to create new string when trying to operate on them once the string is created
	fmt.Println(greetingWithName)

	var ch rune = 'a' // runes are declared using single quotes double quotes are for strings
	fmt.Println(ch) // this prints the integer value associated with the value of the rune 'a'

	fmt.Printf("%c\n", ch) // we should get the actual character this time

	// convert runes to strings
	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr is %T\n", cstr) // this prints out type string proving we converted ch from rune to string

	// iterating over runes
	jhello := "hello" 
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue) // using %c gives actual runes if we use %v we would get the int values of each rune
	}

	// you can also print emojis but I am not sure how to do that but its possible. Look it up if needed

	// recursive palindrome attempt let it be said I hate recursion
	fmt.Println(isPalindrome("poop"))
	fmt.Println(isPalindrome("nope"))
	fmt.Println(isPalindrome("nijyuyonshoku"))
	fmt.Println(isPalindrome("a"))
}

func isPalindrome(s string) bool {
	return checkString(s, 0, len(s)-1)
}

func checkString(s string, leftIndex int, rightIndex int) bool {
	// base case
	if leftIndex >= rightIndex {
		return true
	}

	if s[leftIndex] != s[rightIndex] {
		return false
	}

	return checkString(s, leftIndex +1, rightIndex - 1)
}
package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\"") // could also wrap this in single ticks ` and that will create a string literal`

	// compile a regular expression pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`) // the hyphen - needs to be at the end if its not used as a range
	
	// test some strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email1:", re.MatchString(email1))
	fmt.Println("Email2:", re.MatchString(email2))

	// capturing groups
	// compile a regex pattern to capture date components
	re = regexp.MustCompile(`(\d){4}-(\d{2})-(\d{2})`)

	// test string
	date := "2024-07-30"

	// find all submatches
	submatches := re.FindStringSubmatch(date)
	fmt.Println(submatches)
	fmt.Println(submatches[0])
	fmt.Println(submatches[1])
	fmt.Println(submatches[2])
	fmt.Println(submatches[3])

	// target string
	str := "Hello World"
	re = regexp.MustCompile(`[aeiou]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// flags 
	// i - case insensitive
	// m - multi line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`) // ?i means case insensitive flag

	// test string
	text := "Golang is great"

	// match

	fmt.Println("Match:", re.MatchString(text)) // go is present in the test string so this returns true
}

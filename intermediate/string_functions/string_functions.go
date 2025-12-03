package main

import (
	"fmt"
	"strings"
)

func main() {
	/*str := "Hello Go!"

	// get the length of str
	fmt.Println(len(str))

	// string concatenation
	str1 := "Hello"
	str2 := "World"
	result := str1 + " " + str2
	fmt.Println(result)

	// get the index of any rune in the string this will give us the ascii value
	fmt.Println(str[0])
	fmt.Println(str[1:7]) // this prints ello G

	// string conversion
	num := 18
	str3 := strconv.Itoa(num) //convert int to string
	fmt.Println(len(str3)) // length of num is 2

	// string splitting
	fruits := "apple, orange, banana"
	fruits1 := "apple-orange-banana"
	parts := strings.Split(fruits, ",") //splitting the fruits string on the comma seperator
	fmt.Println(parts)
	parts1 := strings.Split(fruits1, "-") // splitting on -
	fmt.Println(parts1)

	// joining strings
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ") // if you want to space them out you have to add the space after the comma
	fmt.Println(joined)

	// checks if a string contains a subset of chracters
	fmt.Println(strings.Contains(str, "Go?")) // The string doesn't contain Go? so it will return false

	// replace a substring within one string with another substring
	replaced := strings.Replace(str, "Go", "Universe", 1) // 1 is the number of occurences to replace I think?
	fmt.Println(replaced)

	// can also remove leading white space
	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// to lower and to upper
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	// repeat function
	fmt.Println(strings.Repeat("foo ", 3))

	// count occurence of substring
	fmt.Println(strings.Count("Hello", "l")) // this will return 2 because weh ave 2 ls

	fmt.Println(strings.HasPrefix("Hello", "He")) // true
	fmt.Println(strings.HasSuffix("Hello", "lo"))

	// advanced techniques
	// regex pattern matching and manipulation

	str5 := "Hel1lo, 123 Go! 11"
	re := regexp.MustCompile(`\d+`) // uses backticks the regex we're using is looking for 1 or more numbers
	matches := re.FindAllString(str5, -1) // -1 means all the matches of the regular expression in the source string
	fmt.Println(matches)

	// unicode UTF-8 package
	str6 := "Hello, 𫖯"
	fmt.Println(utf8.RuneCountInString(str6)) // 8 characters in the string (it counts spaces)

	// efficient string building in memory efficient way since strings in Go are immutable allows you to build final string incrementally
	*/

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Make the final string from the builder
	// Convert the builder to a string
	result := builder.String() // now we have the final string stored in result
	fmt.Println(result)

	// Using write rune to add a character
	builder.WriteRune(' ') // runes are enclosed in single quotes
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// in order to start a new string we have to reset the builder
	builder.Reset()
	builder.WriteString("Starting Fresh")
	result = builder.String()
	fmt.Println(result)
}

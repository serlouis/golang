package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// seed the random number generator
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the Guessing Game")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess the number?")

	var guess int
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)

		// check if the guess == target
		if guess == target {
			fmt.Println("🎉 Congratulations, you guessed right!")
			break
		} else if guess < target {
			fmt.Println("Too low, try guessing a higher number.")
		} else {
			fmt.Println("Too high, try guessing a lower number.")
		}
	}
}

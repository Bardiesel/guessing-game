package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate a random number between 0 and 10
	secret := getRandomNumber()
	// fmt.Println(secret)
	for matching := false; !matching; {
		// Get User Guess
		guess := getUserGuess()
		// fmt.Println(secret, guess)

		// Check if guess is correct
		matching = isMatching(secret, guess)
		// fmt.Println(matching)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Too high")
		return false
	} else if guess < secret {
		fmt.Println("Too low")
		return false
	} else {
		fmt.Println("You win!")
		return true
	}
}

func getUserGuess() int {
	fmt.Print("Guess the number: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Failed to read input", err)
	} else {
		fmt.Println("You entered", input)
	}
	return input
}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 11
}

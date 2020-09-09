package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var startChoice int
	choice := [3]string{"rock", "paper", "scissor"}
	ai := choice[rand.Intn(3)]
	var playerChoice string
	win := false
	numGames := 0
	numWon := 0

	for {
		fmt.Println("Choose from the following below: \n1. Rock, Paper, Scissor \n2. Guessing Game \n3. Quit")
		fmt.Scanln(&startChoice)
		if startChoice < 1 || startChoice > 3 {
			fmt.Println("Invalid Option. Please try again.")
		} else {
			break
		}
	}

	if startChoice == 2 {
		//rand.Seed(time.Now().UnixNano())
		num := rand.Intn(100) + 1
		var guessNum int
		fmt.Println("Guess a number between 1 to 100.")
		guess := 1
		for {
			fmt.Scanln(&guessNum)
			if guessNum < 1 || guessNum > 100 {
				fmt.Println("Number is out of range! Please try again.")
				continue
			} else if guessNum > num {
				fmt.Println("Your guess is larger than the number. Please try again.")
			} else if guessNum < num {
				fmt.Println("Your guess is smaller than the number. Please try again.")
			} else {
				fmt.Println("Your guess is correct! Congrats. You guessed right after", guess, "tries.")
				break
			}
			guess++
		}
	} else if startChoice == 1 {
		for win == false {
			fmt.Println("Rock, Paper, Scissor or Stop?")
			fmt.Scanln(&playerChoice)
			playerChoice = strings.TrimSpace(playerChoice)
			playerChoice = strings.ToLower(playerChoice)

			if ai == playerChoice {
				fmt.Println("Tie!")
			} else if playerChoice == "rock" {
				if ai == "paper" {
					fmt.Println("You lose!", ai, "covers", playerChoice)
				} else {
					numWon++
					fmt.Println("You win!", playerChoice, "smashes", ai)
				}
			} else if playerChoice == "paper" {
				if ai == "scissor" {
					fmt.Println("You lose!", ai, "cuts", playerChoice)
				} else {
					numWon++
					fmt.Println("You win!", playerChoice, "covers", ai)
				}
			} else if playerChoice == "scissor" {
				if ai == "rock" {
					fmt.Println("You lose!", ai, "smashes", playerChoice)
				} else {
					numWon++
					fmt.Println("You win!", playerChoice, "cut", ai)
				}
			} else if playerChoice == "stop" {
				fmt.Println("You won", numWon, "out of", numGames, "games! Goodbye.")
				break
			} else {
				fmt.Println("That's not a valid play. Check your spelling!")
			}
			numGames++
			win = false
			ai = choice[rand.Intn(2)]
		}
	} else {
		os.Exit(0)
	}
}

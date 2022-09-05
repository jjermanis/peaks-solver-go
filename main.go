package main

import (
	"fmt"
	"strings"
)

func main() {
	var gameState GameState
	gameState.Initialize("possibleAnswers.txt")
	PlayGame(gameState)
}

func PlayGame(state GameState) {
	fmt.Println("Welcome to peaks-solver-go. This program will help you solve Wordle Peaks puzzles.")
	fmt.Println("To use, guess what this program suggests. Then, let this program know the result.")
	fmt.Println("Enter the five colors from the result. G for green, B for blue, and O for orange.")

	var currGuess, result string

	for i := 0; i < 6; i++ {
		// TODO implement the logic for best guess
		currGuess = state.possibleAnswers[0]
		fmt.Println("Please guess:", currGuess)
		result = PromptResult()
		if result == "GGGGG" {
			fmt.Println("Congrats!")
			return
		}
		state.UpdateAfterGuess(currGuess, result)
	}
	fmt.Println("Looks like you ran out of guesses. My fault.")
}

func PromptResult() string {
	var result string
	for {
		fmt.Print("Result? ")
		fmt.Scan(&result)
		result = strings.ToUpper(result)
		if !IsValidResult(result) {
			fmt.Println("Incorrect format. Results should be five letters long. G for green, B for blue, and O for orange.")
		} else {
			return result
		}
	}
}

func IsValidResult(r string) bool {
	var VALID_RESULT_CHARS string = "GBO"
	if len(r) != 5 {
		return false
	} else {
		for i := 0; i < 5; i++ {
			loc := strings.Index(VALID_RESULT_CHARS, string(r[i]))
			if loc < 0 {
				return false
			}
		}
	}
	return true
}

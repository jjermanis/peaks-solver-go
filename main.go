// Peaks Solver helps play the Wordle Peaks game with effective strategy.
// It provides optimal guesses, and asks for the result of each guess.
// This approach provides an effective way to get to a solution.
//
// This application is interactive - user must read proposed guesses from
// the commandline, and then enter the results of each guess.
//
// The game is available here: [Wordle Peaks]: https://vegeta897.github.io/wordle-peaks/
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

// PlayGame helps the user get through one game of Wordle Peaks.
// state should contain an initialized GameState with all possible answers.
func PlayGame(state GameState) {
	fmt.Println("Welcome to peaks-solver-go. This program will help you solve Wordle Peaks puzzles.")
	fmt.Println("To use, guess what this program suggests. Then, let this program know the result.")
	fmt.Println("Enter the five colors from the result. G for green, B for blue, and O for orange.")

	var currGuess, result string

	for i := 0; i < 6; i++ {
		currGuess = state.bestGuess
		fmt.Println("Please guess:", currGuess)
		result = promptResult()
		if result == "GGGGG" {
			fmt.Println("Congrats!")
			return
		}
		state.UpdateAfterGuess(currGuess, result)
		displayStatus(state)
	}
	fmt.Println("Looks like you ran out of guesses. My fault.")
}

func promptResult() string {
	var result string
	for {
		fmt.Print("Result? ")
		fmt.Scan(&result)
		result = strings.ToUpper(result)
		if !isValidResult(result) {
			fmt.Println("Incorrect format. Results should be five letters long. G for green, B for blue, and O for orange.")
		} else {
			return result
		}
	}
}

func isValidResult(r string) bool {
	// A valid result must be exactly 5 characters, with each representing one
	// of the three valid color (green, blue, and orange).
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

func displayStatus(state GameState) {
	possibleAnswerCount := len(state.possibleAnswers)
	if possibleAnswerCount > 1 {
		fmt.Println(possibleAnswerCount, "options remain.")
	} else if possibleAnswerCount == 1 {
		fmt.Println("Only 1 possible answer remains.")
	} else {
		fmt.Println("Error. No possible answers known. Sorry - I cannot help anymore.")
	}
}

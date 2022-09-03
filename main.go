package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	possibleAnswers := InitPossibleAnswers()
	PlayGame(possibleAnswers)
}

func PlayGame(possibleAnswers []string) {
	fmt.Println("Welcome to peaks-solver-go. This program will help you solve Wordle Peaks puzzles.")
	fmt.Println("To use, guess what this program suggests. Then, let this program know the result.")
	fmt.Println("Enter the five colors from the result. G for green, B for blue, and O for orange.")

	var currGuess, result string

	for i := 0; i < 6; i++ {
		// TODO implement the logic for best guess
		currGuess = possibleAnswers[0]
		fmt.Println("Please guess:", currGuess)
		result = PromptResult()
		if result == "GGGGG" {
			fmt.Println("Congrats!")
			return
		}
		possibleAnswers = UpdateFromResult(possibleAnswers, currGuess, result)
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

func InitPossibleAnswers() []string {
	bytes, err := os.ReadFile("possibleAnswers.txt")
	if err != nil {
		fmt.Print(err)
		return nil
	}
	result := strings.Split(string(bytes), ",")
	return result
}

func UpdateFromResult(currAnswers []string, guess string, result string) []string {
	var newAnswers []string
	currAnswerCount := len(currAnswers)
	for i := 0; i < currAnswerCount; i++ {
		var currWord = currAnswers[i]
		if WordFollowsResult(currWord, guess, result) {
			newAnswers = append(newAnswers, currWord)
		}
	}
	return newAnswers
}

func WordFollowsResult(word string, guess string, result string) bool {
	for i := 0; i < 5; i++ {
		currResult := result[i]
		currWord := word[i]
		currGuess := guess[i]

		if currResult == 'G' {
			if currWord != currGuess {
				return false
			}
		} else if currResult == 'B' {
			if currWord >= currGuess {
				return false
			}
		} else if currResult == 'O' {
			if currWord <= currGuess {
				return false
			}
		}
	}
	return true
}

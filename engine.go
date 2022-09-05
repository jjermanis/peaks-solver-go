package main

import (
	"fmt"
	"os"
	"strings"
)

type GameState struct {
	possibleAnswers []string
}

func (state *GameState) initPossibleAnswers(filename string) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
		return
	}
	answers := strings.Split(string(bytes), ",")
	state.possibleAnswers = answers
}

func isWordPossibleAnswer(word string, guess string, result string) bool {
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

func (state *GameState) Initialize(filename string) {
	state.initPossibleAnswers(filename)
}

func (state *GameState) UpdateAfterGuess(guess string, result string) {
	var newAnswers []string
	currAnswerCount := len(state.possibleAnswers)
	for i := 0; i < currAnswerCount; i++ {
		var currWord = state.possibleAnswers[i]
		if isWordPossibleAnswer(currWord, guess, result) {
			newAnswers = append(newAnswers, currWord)
		}
	}
	state.possibleAnswers = newAnswers
}

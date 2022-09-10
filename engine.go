package main

import (
	"fmt"
	"os"
	"strings"
)

// GameState contains the data used to help play Peaks.
// Recommend calling Initialize() to set up with all potential answers.
type GameState struct {
	possibleAnswers []string // remaining possible correct answers
	bestGuess       string   // the optimal next guess
	lowestLetters   [5]byte  // the lowest possible letter for each position; starts at A
	highestLetters  [5]byte  // the highest possible letter for each position; starts at Z
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

func (state *GameState) updateBestGuess() {
	bestScore := 999999
	var bestGuess string
	currAnswerCount := len(state.possibleAnswers)
	for i := 0; i < currAnswerCount; i++ {
		var currGuess = state.possibleAnswers[i]
		var currScore = state.scoreGuess(currGuess)
		if currScore < bestScore {
			bestScore = currScore
			bestGuess = currGuess
		}
	}
	state.bestGuess = bestGuess
}

func (state GameState) scoreGuess(guess string) int {
	var score = 0
	for i := 0; i < 5; i++ {
		currTarget := (state.highestLetters[i]-state.lowestLetters[i])/2 + state.lowestLetters[i]
		currValue := int(currTarget) - int(guess[i])
		if currValue < 0 {
			currValue = -currValue
		}
		score += currValue
	}
	return score
}

func (state *GameState) Initialize(filename string) {
	state.initPossibleAnswers(filename)
	for i := 0; i < 5; i++ {
		state.lowestLetters[i] = 'a'
		state.highestLetters[i] = 'z'
	}
	state.updateBestGuess()
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
	for i := 0; i < 5; i++ {
		currResult := result[i]
		currGuess := guess[i]
		if currResult == 'G' {
			state.lowestLetters[i] = currGuess
			state.highestLetters[i] = currGuess
		} else if currResult == 'B' {
			state.highestLetters[i] = currGuess - 1
		} else if currResult == 'O' {
			state.lowestLetters[i] = currGuess + 1
		}
	}
	state.possibleAnswers = newAnswers
	state.updateBestGuess()
}

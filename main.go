package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Game struct {
	secretWord     string
	guesses        []byte
	chancesLeft    uint
	correctGuesses []byte
}

func NewGame(secretWord string) Game {
	return Game{
		secretWord:     secretWord,
		guesses:        []byte{},
		chancesLeft:    7,
		correctGuesses: []byte{},
	}
}

func hasPunctuations(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return true
		}
	}
	return false
}

func getSecretWord(wordFileName string) string {

	var allowedWords []string

	WordFile, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Error in %v cause of %v", WordFile, err)
		panic(errMessage)
	}

	defer WordFile.Close()

	scanner := bufio.NewScanner(WordFile)

	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && !hasPunctuations(word) {
			allowedWords = append(allowedWords, word)
		}
	}

	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]

}

func checkGuess(state Game, guess byte) Game {
	if state.chancesLeft > 1 && !bytes.Contains(state.guesses, []byte{guess}) {
		if strings.ContainsRune(state.secretWord, rune(guess)) {
			state.correctGuesses = append(state.correctGuesses, guess)
			state.guesses = append(state.guesses, guess)
		} else {
			state.guesses = append(state.guesses, guess)
			state.chancesLeft--
		}
	}
	return state
}

func hasWon(state Game) bool {
	for _, ch := range state.secretWord {
		if !bytes.Contains(state.correctGuesses, []byte{byte(ch)}) {
			return false
		}
	}
	return true

}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}

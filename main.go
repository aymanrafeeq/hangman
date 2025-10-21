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
	if state.chancesLeft > 0 && !bytes.Contains(state.guesses, []byte{guess}) {
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

func hasLose(state Game) bool {
	return state.chancesLeft == 0 && !hasWon(state)
}

func displayWord(state Game) {
	for _, ch := range state.secretWord {
		if bytes.Contains(state.correctGuesses, []byte{byte(ch)}) {
			fmt.Printf("%c ", ch)
		} else {
			fmt.Print("_ ")
		}
	}
	fmt.Println()
}

func main() {
	word := getSecretWord("/usr/share/dict/words")
	game := NewGame(word)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Try to guess the word! You have 7 chances.")
	fmt.Println()

	for game.chancesLeft > 0 {
		displayWord(game)
		fmt.Printf("Chances left: %d\n", game.chancesLeft)
		fmt.Print("Enter a letter: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if len(input) != 1 {
			fmt.Println("Please enter only one letter.")
			continue
		}

		guess := input[0]

		if guess < 'a' || guess > 'z' {
			fmt.Println("Please enter a valid character....!!!")
			fmt.Println()
			continue
		}

		if bytes.Contains(game.guesses, []byte{guess}) {
			fmt.Printf("You already guessed '%c'.Try a different letter.\n\n", guess)
			continue
		}

		game = checkGuess(game, input[0])

		if hasWon(game) {
			fmt.Println("\nCongratulations! You guessed the word:", game.secretWord)
			return
		}
		fmt.Println()
	}

	fmt.Println("\nGame Over! The word was:", game.secretWord)

}

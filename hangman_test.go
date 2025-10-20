package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func createDictFile(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "hangman-dict")
	if err != nil {
		fmt.Println("Couldn't create temp file.")
	}
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func TestSecretWordNocapital(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if secretWord != strings.ToLower(secretWord) {
		t.Errorf("Should not get words with Capital letters. Got %s", secretWord)
	}
}

func TestSecreWordtLength(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if len(secretWord) < 6 {
		t.Errorf("Words should contain a minimum of 6 letters. Got %v", len(secretWord))
	}
}

func TestSecreWordNoPunctuation(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	for _, ch := range secretWord {
		if !(ch >= 'a' && ch <= 'z') {
			t.Errorf("Words should not contain a punctuation Got %v", secretWord)
		}
	}
}

func TestCorrectGuess(t *testing.T) {
	secretWord := "elephant"
	guess := 'a'
	currentState := NewGame(secretWord)
	newState := checkGuess(currentState, byte(guess))

	expected := Game{
		secretWord:     currentState.secretWord,
		guesses:        append(currentState.guesses, byte(guess)),
		correctGuesses: append(currentState.correctGuesses, byte(guess)),
		chancesLeft:    7,
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guesses, expected.guesses) {
		t.Errorf("Guess should be %q but got %q", expected.guesses, newState.guesses)
	}
	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}
	if newState.chancesLeft != expected.chancesLeft {
		t.Errorf("chances left has modified")
	}
}

func TestCorrectGuess2(t *testing.T) {
	secretWord := "elephant"
	guess := 'l'
	currentState := Game{
		secretWord:     secretWord,
		guesses:        []byte{'c', 'a'},
		correctGuesses: []byte{'a'},
		chancesLeft:    6,
	}
	newState := checkGuess(currentState, byte(guess))

	expected := Game{
		secretWord:     currentState.secretWord,
		guesses:        append(currentState.guesses, byte(guess)),
		correctGuesses: append(currentState.correctGuesses, byte(guess)),
		chancesLeft:    currentState.chancesLeft,
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guesses, expected.guesses) {
		t.Errorf("Guess should be %q but got %q", expected.guesses, newState.guesses)
	}
	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}
	if newState.chancesLeft != expected.chancesLeft {
		t.Errorf("chances left has modified")
	}
}

func TestWrongGuess(t *testing.T) {
	secretWord := "soldier"
	guess := 'b'
	currentState := Game{
		secretWord:     secretWord,
		guesses:        []byte{'a'},
		correctGuesses: []byte{},
		chancesLeft:    6,
	}

	newState := checkGuess(currentState, byte(guess))

	expected := Game{
		secretWord:     secretWord,
		guesses:        []byte{'a', 'b'},
		correctGuesses: []byte{},
		chancesLeft:    5,
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guesses, expected.guesses) {
		t.Errorf("Guess should be %q but got %q", expected.guesses, newState.guesses)
	}

	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

	if !(newState.chancesLeft == expected.chancesLeft) {
		t.Errorf("Chances left is modified")
	}
}

func TestAlreadyGuess(t *testing.T) {
	secretWord := "soldier"
	guess := 'a'
	currentState := Game{
		secretWord:     secretWord,
		guesses:        []byte{'a'},
		correctGuesses: []byte{},
		chancesLeft:    6,
	}

	newState := checkGuess(currentState, byte(guess))
	expected := Game{
		secretWord:     secretWord,
		guesses:        []byte{'a'},
		correctGuesses: []byte{},
		chancesLeft:    6,
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guesses, expected.guesses) {
		t.Errorf("Guess should be %q but got %q", expected.guesses, newState.guesses)
	}
	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}
	if !(newState.chancesLeft == expected.chancesLeft) {
		t.Errorf("Remaining chances is modified")
	}
}

func TestGameOverWin(t *testing.T) {
	state := Game{
		secretWord:     "apple",
		correctGuesses: []byte{'a', 'p', 'l', 'e'},
		chancesLeft:    5,
	}

	if !hasWon(state) {
		t.Errorf("Expected win, but got not won")
	}
}

func TestGameOverLose(t *testing.T) {
	state := Game{
		secretWord:     "apple",
		correctGuesses: []byte{'a', 'e'},
		chancesLeft:    0,
	}
	if !hasLose(state) {
		t.Errorf("Expected lose, but got not lose")
	}
}

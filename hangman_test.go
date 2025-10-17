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
	userInput := "a"
	state := NewGame(secretWord)
	newState := checkGuess(state, userInput)

	expected := Game{
		secretWord:     state.secretWord,
		guesses:        append(state.guesses, userInput[0]),
		correctGuesses: append(state.correctGuesses, userInput[0]),
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

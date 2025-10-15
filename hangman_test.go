package main

import (
	"strings"
	"testing"
)

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

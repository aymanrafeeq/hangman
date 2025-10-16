package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Game struct {
	secrerWord     string
	guesses        []byte
	chancesLeft    uint
	correctGuesses []byte
}

func NewGame(secrerWord string) Game {
	return Game{
		secrerWord:     secrerWord,
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

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}

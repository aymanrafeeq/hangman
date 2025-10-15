package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func getSecretWord(wordFileName string) string {
	WordFile, err := os.Open(wordFileName)
	if err != nil {
		errMessage := fmt.Sprintf("Error in %v cause of %v", WordFile, err)
		panic(errMessage)
	}

	defer WordFile.Close()

	scanner := bufio.NewScanner(WordFile)

	var wordList []string
	for scanner.Scan() {
		wordList = append(wordList, scanner.Text())
	}
	randomNum := rand.Intn(len(wordList))

	randWord := wordList[randomNum]

	return randWord

}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}

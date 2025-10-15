package main

import (
	"bufio"
	"fmt"
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return "elephant"

}
func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))
}

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func ReadFileByLine(path string) []string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	return lines
}

func extractNumberFromWord(word string) int {
	number := ""
	var positionFirstNumber int

	// For example, if you found the first digit on position 9 you just need to use len(word) - 1 until get at pos 9.
	for i, char := range word {
		if !unicode.IsLetter(char) {
			number = number + string(char)
			positionFirstNumber = i
			break
		}
	}

	// Edge case: Just one number in the word.
	for i := len(word) - 1; i != positionFirstNumber-1; i-- {
		if !unicode.IsLetter(rune(word[i])) {
			number = number + string(word[i])
			break
		}
	}

	numberParsed, err := strconv.Atoi(number)

	if err != nil {
		panic("Oupsie")
	}

	return numberParsed
}

func main() {

	var counter int
	listOfWords := ReadFileByLine("./input.csv")

	for _, word := range listOfWords {
		// fmt.Println(word, extractNumberFromWord(word))
		counter += extractNumberFromWord(word)
	}

	fmt.Println(counter)
}

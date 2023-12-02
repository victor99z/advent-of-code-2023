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

	posFirstDigit, firstDigit := verifySpelledFirstDigits(word)

	// fmt.Printf("%d %d \n", posFirstDigit, firstDigit)

	if posFirstDigit < positionFirstNumber {
		positionFirstNumber = posFirstDigit
		number = strconv.Itoa(firstDigit)
	}

	posSecondDigit, secondDigit := verifySpelledSecondDigits(word)

	var positionSecondNumber int
	number2 := ""

	// Edge case: Just one number in the word.
	for i := len(word) - 1; i != positionFirstNumber-1; i-- {
		if !unicode.IsLetter(rune(word[i])) {
			positionSecondNumber = i
			number2 = number2 + string(word[i])
			break
		}
	}

	if posSecondDigit > positionSecondNumber {
		number = number + strconv.Itoa(secondDigit)
	} else {
		number = number + number2
	}

	fmt.Println(word, number)
	numberParsed, err := strconv.Atoi(number)

	if err != nil {
		panic("Oupsie")
	}

	return numberParsed
}

// 65fivetnsseight8lcvgkkglslcjtjssxmgtvk

// Part 2
func verifySpelledFirstDigits(word string) (int, int) {

	digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}
	// digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	appearPos := 9999999
	numberFilled := -1

	//fmt.Printf("word: %s", word)
	for spelled, digit := range digits {
		idx := strings.Index(word, spelled)
		if idx != -1 {
			if idx < appearPos {
				appearPos = idx
				numberFilled = digit
			}
		}
	}
	// fmt.Printf(" pos: %d number: %d", appearPos, numberFilled)
	return appearPos, numberFilled
}

func verifySpelledSecondDigits(word string) (int, int) {

	digits := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9}

	appearPos := -1
	numberFilled := -1

	//fmt.Printf("word: %s", word)
	for spelled, digit := range digits {
		idx := strings.LastIndex(word, spelled)
		if idx != -1 {
			if idx > appearPos {
				appearPos = idx
				numberFilled = digit
			}
		}
	}
	// fmt.Printf(" pos: %d number: %d \n", appearPos, numberFilled)
	return appearPos, numberFilled
}

func main() {

	var counter int
	listOfWords := ReadFileByLine("./input.csv")

	for _, word := range listOfWords {
		counter += extractNumberFromWord(word)
	}

	fmt.Println(counter)
}

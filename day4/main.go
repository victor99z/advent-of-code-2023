package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	var matches1, matches2 []string

	data := map[string]int{}

	allAcc := 0
	for _, v := range lines {
		sepCards := strings.Split(v, "|")

		regExp, _ := regexp.Compile(`(\d+)`)

		matches1 = regExp.FindAllString(sepCards[0], -1)
		matches2 = regExp.FindAllString(sepCards[1], -1)

		cardNumber := matches1[0]

		//		data[cardNumber] =

		cardAcc := 0

		for _, v := range matches1[1:] {
			for _, k := range matches2 {
				if v == k {
					cardAcc = cardAcc * 2
				}
				if v == k && cardAcc == 0 {
					cardAcc = 1
				}
			}
		}
		allAcc += cardAcc
		data[cardNumber] = cardAcc
	}
	fmt.Println(data, allAcc)
}

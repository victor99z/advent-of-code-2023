package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseData() map[int]int {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	var matches1, matches2 []string

	data := map[int]int{}

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
					cardAcc++
				}
			}
		}

		cardNumberInt, _ := strconv.Atoi(cardNumber)

		data[cardNumberInt] = cardAcc

	}
	return data
}

type Card struct {
	number, matches int
}

// func getMatches(ca []Card) int {
// 	if len(ca) == 0 {
// 		return 1
// 	}
// 	return 1 + getMatches(ca[])
// }

/*

{
	1: 4
	2: 2
	3: 1
	4: 1
	5: 0
	6: 0
}


card 1 : [2,3,4,5]


*/

func checkCards(data *map[int]int, card int, acc int) int {
	if v, _ := (*data)[card]; v == 0 {
		return 1
	}
	return checkCards(data, card+1, acc)
}

// map[1:4 2:2 3:2 4:1 5:0 6:0]

// map[1:1 2:2 3:3 4:1 5:1]

func main() {

	data := parseData()

	anotherData := map[int]int{}

	// cards := []Card

	for k, v := range data {
		for i := k; i <= v+1; i++ {
			if v, ok := anotherData[i]; ok {
				anotherData[i] = v + 1
			} else {
				anotherData[i] = 1
			}
		}
	}

	fmt.Println(data, anotherData)

}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

var number_inputs int

func parseData() []int {
	file, _ := os.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	var matches1, matches2 []string

	//data := map[int]int{}
	data := []int{}

	for _, v := range lines {
		number_inputs++
		sepCards := strings.Split(v, "|")

		regExp, _ := regexp.Compile(`(\d+)`)

		matches1 = regExp.FindAllString(sepCards[0], -1)
		matches2 = regExp.FindAllString(sepCards[1], -1)

		// cardNumberInt, _ := strconv.Atoi(matches1[0])

		cardAcc := 0

		for _, v := range matches1[1:] {
			for _, k := range matches2 {
				if v == k {
					cardAcc++
				}
			}
		}

		data = append(data, cardAcc)
	}
	return data
}

func main() {

	data := parseData()

	copies := make([]int, number_inputs)

	for i := range copies {
		copies[i] = 1
	}

	fmt.Println(copies, data)

	for idx, matches := range data {
		for j := 1; j < matches+1; j++ {
			copies[idx+j] += copies[idx]
		}
	}

	acc := 0
	for _, v := range copies {
		acc += v
	}
	// Sdds de um reduce agora nessas horas :sadge:
	fmt.Println(acc)
}

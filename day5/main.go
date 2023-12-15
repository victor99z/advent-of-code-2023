package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func part1() {
	data, _ := os.ReadFile("./day5/test.txt")

	lines := strings.Split(string(data), "\n")

	re := regexp.MustCompile(`\d+`)

	fmt.Println(data)

	// Find all matches in the input string
	matches := re.FindAllString(lines[0], -1)

	fmt.Println(matches)

	for _, v := range lines {

		switch {
		case strings.Contains(v, "seed-to-soil"):
		case strings.Contains(v, "soil-to-fertilizer"):
		case strings.Contains(v, "fertilizer-to-water"):
		case strings.Contains(v, "water-to-light"):
		case strings.Contains(v, "light-to-temperature"):
		case strings.Contains(v, "temperature-to-humidity"):
		case strings.Contains(v, "humidity-to-location"):

		}
	}

}

func main() {
	part1()
}

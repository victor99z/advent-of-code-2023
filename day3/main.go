package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func read(filepath string) [][]string {
	file, err := os.ReadFile(filepath)

	matriz := make([][]string, 140)

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(file), "\n")

	for i, line := range lines {
		matriz[i] = strings.Split(line, "")
	}

	return matriz
}

type PosOffSet struct {
	posX int
	posY int
}

func generateOffsets() []PosOffSet {
	// x,y
	// x + 1, y
	// x - 1, y
	// x - 1 , y + 1
	// x - 1 , y - 1
	// x + 1, y + 1
	// x + 1, y -1
	// x , y - 1
	// x, y + 1

	var offsets []PosOffSet

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			offsets = append(offsets, PosOffSet{posX: i, posY: j})
		}
	}
	return offsets
}

func getNeighboursPos(x, y int, matriz *[][]string) []int {
	return []int{}
}

func main() {
	// matriz := read("input")

	// getNeighboursPos()
	fmt.Println(len(generateOffsets()), generateOffsets())

	// // Compile the regex pattern
	// re := regexp.MustCompile(`(\d+)`)

	// // Find all matches in the input string
	// matches := re.FindAllString(inputString, -1)

	// // Print the matches
	// fmt.Println(matches)

}

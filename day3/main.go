package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const MATRIZ_SIZE = 140

func read(filepath string) [][]string {
	file, err := os.ReadFile(filepath)

	matriz := make([][]string, MATRIZ_SIZE)

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

func getNeighboursPos(x, y int) []PosOffSet {

	offsets := generateOffsets()

	coordsOffset := []PosOffSet{}

	for _, offset := range offsets {
		if (x+offset.posX >= 0 && x+offset.posX < MATRIZ_SIZE) && ((y+offset.posY) >= 0 && (y+offset.posY) < MATRIZ_SIZE) {

			coordsOffset = append(coordsOffset, PosOffSet{posX: x + offset.posX, posY: y + offset.posY})

		}
	}

	return coordsOffset
}

func verifyNearPos(x, y int, matriz *[][]string) bool {
	list := getNeighboursPos(x, y)

	// fmt.Println(list)

	for _, v := range list {
		stringToRune := []rune((*matriz)[v.posX][v.posY])[0]
		if (*matriz)[v.posX][v.posY] != "." && !unicode.IsDigit(stringToRune) {
			return true
		}
	}
	return false
}

// Fazer com regex em cada linha, pega os numeros e depois usa strings.contains para pegar os index, iterava ao redor usando as funcoes de vizinhos :x

type Digit struct {
	number   int
	posStart []int
	posEnd   []int
}

func main() {

	matriz := read("input")

	// file, _ := os.ReadFile("input")

	// lines := strings.Split(string(file), "\n")

	// re := regexp.MustCompile(`(\d+)`)

	// allcases := [][]string{}

	// acc := 0

	// for _, v := range lines {

	// 	matches := re.FindAllString(v, -1)

	// 	acc += len(matches)

	// 	allcases = append(allcases, matches)
	// }

	// fmt.Println(allcases, acc)

	// // Compile the regex pattern
	// re := regexp.MustCompile(`(\d+)`)

	// // Find all matches in the input string
	// matches := re.FindAllString(w1, -1)

	// // Print the matches
	// fmt.Println(matches)

	// getNeighboursPos()

	// maperson := map[string]int{}

	// listOfNumbers := []string{}
	// validNumbers := []string{}

	// if verifyNearPos(0, 25, &matriz) {
	// 	fmt.Println("yay", matriz[0][25])
	// }

	digit := ""
	valid := false
	mapeamento := map[string]bool{}
	acc := 0
	for i := 0; i < MATRIZ_SIZE; i++ {
		for j := 0; j < MATRIZ_SIZE; j++ {
			x := matriz[i][j]
			yayFodase := "."
			if j > 0 {
				yayFodase = matriz[i][j-1]
			}
			anterior := []rune(yayFodase)[0]
			y := []rune(x)[0]
			if unicode.IsDigit(y) {
				digit += x

				if verifyNearPos(i, j, &matriz) {
					valid = true
				}

			} else if unicode.IsDigit(anterior) {
				mapeamento[digit] = valid
				valid = false
				digit = ""
			}
		}
		if len(digit) > 0 && valid {
			mapeamento[digit] = valid
		}
		for k, v := range mapeamento {
			if v {
				number, _ := strconv.Atoi(k)
				acc += number
			}
		}
		mapeamento = map[string]bool{}
	}

	fmt.Println(acc)

}

// 525394
// 446 + 524948 + 941
// 2268225

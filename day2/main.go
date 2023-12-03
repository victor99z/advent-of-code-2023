package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	red   int
	green int
	blue  int
}

type Game struct {
	number int
	sets   []Set
}

func parseAllGames(path string) []Game {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	var listGames []Game

	for _, v := range lines {
		gameRetrieved := parseGameByLine(v)

		if gameRetrieved.number == 0 {
			continue
		}
		listGames = append(listGames, gameRetrieved)
	}

	return listGames
}

func parseGameByLine(gameString string) Game {

	var game1 Game

	temp := strings.Split(gameString, ":")

	gameNumber := strings.Split(temp[0], " ")[1]
	sets := strings.Split(temp[1], ";")

	var err error
	game1.number, err = strconv.Atoi(gameNumber)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range sets {

		cube := strings.Split(v, ",")
		set := Set{}

		for _, c := range cube {

			colorNumber, err := strconv.Atoi(strings.Split(c, " ")[1])

			if err != nil {
				log.Fatal(err)
			}

			switch {
			case strings.Contains(c, "red"):
				if colorNumber > 12 {
					return Game{}
				}
				set.red = colorNumber
			case strings.Contains(c, "green"):
				if colorNumber > 13 {
					return Game{}
				}
				set.green = colorNumber
			case strings.Contains(c, "blue"):
				if colorNumber > 14 {
					return Game{}
				}
				set.blue = colorNumber
			}
		}
		game1.sets = append(game1.sets, set)
	}
	return game1
}

func main() {

	games := parseAllGames("./input")

	var counter int

	for _, v := range games {
		counter += v.number
	}
	fmt.Println(counter)
}

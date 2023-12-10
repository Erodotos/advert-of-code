package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	id    int
	red   int
	green int
	blue  int
}

type Game struct {
	id     string
	rounds []Round
}

func main() {

	configuration := []int{12, 13, 14}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading file.")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	games := []Game{}

	// Parse data
	for fileScanner.Scan() {
		line := fileScanner.Text()

		g := new(Game)

		elements := strings.Split(line, ":")
		g.id = strings.Split(elements[0], " ")[1]

		for index, round := range strings.Split(elements[1], ";") {
			r := new(Round)
			r.id = index + 1
			for _, cubes := range strings.Split(strings.TrimSpace(round), ", ") {
				switch c := strings.Split(strings.TrimSpace(cubes), " ")[1]; c {
				case "red":
					r.red, _ = strconv.Atoi(strings.Split(cubes, " ")[0])
				case "green":
					r.green, _ = strconv.Atoi(strings.Split(cubes, " ")[0])
				case "blue":
					r.blue, _ = strconv.Atoi(strings.Split(cubes, " ")[0])
				default:
					break
				}
			}
			g.rounds = append(g.rounds, *r)
		}
		games = append(games, *g)
	}

	// Process data
	sumOfIds := 0
	invalidGame := false
	for _, game := range games {
		for _, round := range game.rounds {
			if round.red > configuration[0] || round.green > configuration[1] || round.blue > configuration[2] {
				invalidGame = true
				break
			}
		}

		if invalidGame == false {
			gameId, _ := strconv.Atoi(game.id)
			sumOfIds += gameId
		}

		invalidGame = false
	}

	fmt.Println(sumOfIds)

}

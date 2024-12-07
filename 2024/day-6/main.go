package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error reading the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	// Read the map
	labMap := [][]byte{}
	guardI := 0
	guardJ := 0
	re := regexp.MustCompile(`[\.\#\>\<\^v]*`)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if re.MatchString(str) {
			labMap = append(labMap, []byte(str))
		}
	}

	// find guard coordinates
	for i := 0; i < len(labMap); i++ {
		for j := 0; j < len(labMap[i]); j++ {
			if labMap[i][j] == '^' || labMap[i][j] == '>' || labMap[i][j] == 'v' || labMap[i][j] == '<' {
				guardI = i
				guardJ = j
				break
			}
		}
	}

	visitedCoordinates := map[string]int{}
	visitedCoordinates[string(guardI)+string(guardJ)] += 1

	// Simulate guard movement
	for guardI < len(labMap)-1 && guardI > 0 && guardJ < len(labMap[0])-1 && guardJ > 0 {
		switch labMap[guardI][guardJ] {
		case '^':
			if labMap[guardI-1][guardJ] == '.' {
				labMap[guardI][guardJ] = '.'
				guardI -= 1
				labMap[guardI][guardJ] = '^'
			} else {
				labMap[guardI][guardJ] = '>'
			}
		case '>':
			if labMap[guardI][guardJ+1] == '.' {
				labMap[guardI][guardJ] = '.'
				guardJ += 1
				labMap[guardI][guardJ] = '>'
			} else {
				labMap[guardI][guardJ] = 'v'
			}
		case 'v':
			if labMap[guardI+1][guardJ] == '.' {
				labMap[guardI][guardJ] = '.'
				guardI += 1
				labMap[guardI][guardJ] = 'v'
			} else {
				labMap[guardI][guardJ] = '<'
			}
		case '<':
			if labMap[guardI][guardJ-1] == '.' {
				labMap[guardI][guardJ] = '.'
				guardJ -= 1
				labMap[guardI][guardJ] = '<'
			} else {
				labMap[guardI][guardJ] = '^'
			}
		}
		visitedCoordinates[string(guardI)+string(guardJ)] += 1
	}

	fmt.Println(len(visitedCoordinates))

}

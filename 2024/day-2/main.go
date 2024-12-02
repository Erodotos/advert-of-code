package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	re := regexp.MustCompile(`^\d+(?:\s\d+)*$`)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error reading the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	safeReportsPartOne := 0
	safeReportsPartTwo := 0

	for fileScanner.Scan() {

		if !re.MatchString(fileScanner.Text()) {
			continue
		}

		levelsStr := strings.Fields(fileScanner.Text())
		levelsInt := make([]int, len(levelsStr))
		for i, number := range levelsStr {
			intNum, err := strconv.Atoi(number)
			if err != nil {
				log.Fatalln("Error converting to integer:", err)
			}
			levelsInt[i] = intNum
		}

		if checkConstraints(levelsInt) {
			safeReportsPartOne++
		}

		if checkConstraints(levelsInt) || canBeFixed(levelsInt) {
			safeReportsPartTwo++
		}
	}

	fmt.Printf("Part1: %d\nPart2: %d\n", safeReportsPartOne, safeReportsPartTwo)
	// output:
	// Part1: 326
	// Part2: 381
}

func canBeFixed(data []int) bool {
	// Check if removing the middle element will
	// satisfy the constraints described in
	// part1 of the problem
	for i := 0; i < len(data); i++ {
		data1 := append([]int{}, data[:i]...)
		data1 = append(data1, data[i+1:]...)
		if checkConstraints(data1) {
			return true
		}
	}
	return false
}

func checkConstraints(data []int) bool {

	prevLvl := data[0]
	currentLvl := data[1]

	// isUnsafe := false
	asc := false
	if currentLvl-prevLvl > 0 {
		asc = true
	}

	for i := 1; i < len(data); i++ {
		currentLvl = data[i]
		if asc {
			if currentLvl-prevLvl < 1 || currentLvl-prevLvl > 3 {
				return false
			}
		} else {
			if currentLvl-prevLvl > -1 || currentLvl-prevLvl < -3 {
				return false
			}
		}
		prevLvl = currentLvl
	}

	return true
}

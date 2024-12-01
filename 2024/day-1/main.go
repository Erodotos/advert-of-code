package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	listA := []int{}
	listB := []int{}
	idFrequencies := map[int]int{}

	re := regexp.MustCompile(`^[\d]{5}[\s]{3}[\d]{5}$`)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error reading the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		if !re.MatchString(fileScanner.Text()) {
			continue
		}

		data := strings.Split(fileScanner.Text(), "   ")

		intA, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatalln("Parsing Error: ", err)
		}
		intB, err := strconv.Atoi(data[1])
		if err != nil {
			log.Fatalln("Parsing Error: ", fileScanner.Text(), err)
		}

		listA = append(listA, intA)
		listB = append(listB, intB)
		idFrequencies[intB] += 1
	}

	sort.Slice(listA, func(i, j int) bool {
		return listA[i] < listA[j]
	})

	sort.Slice(listB, func(i, j int) bool {
		return listB[i] < listB[j]
	})

	distance := 0
	similarity := 0
	for i := 0; i < len(listA); i++ {
		distance += int(math.Abs(float64(listA[i]) - float64(listB[i])))
		similarity += listA[i] * idFrequencies[listA[i]]
	}

	fmt.Println(distance)
	fmt.Println(similarity)
}

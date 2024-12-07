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
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error reading the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	// Read rules
	reRules := regexp.MustCompile(`\d{1,2}\|\d{1,2}`)
	rulesMap := map[int][]int{}
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if reRules.MatchString(str) {
			pages := strings.Split(str, "|")
			page1, _ := strconv.Atoi(pages[0])
			page2, _ := strconv.Atoi(pages[1])
			rulesMap[page1] = append(rulesMap[page1], page2)
		}
	}

	// Reset file scanner
	f.Seek(0, 0)
	fileScanner = bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	// Check for valid updates
	validMiddleValuesSum := 0
	fixedMiddleValuesSum := 0
	reUpdates := regexp.MustCompile(`(\d{1,2}\,)+`)
	for fileScanner.Scan() {
		str := fileScanner.Text()
		if reUpdates.MatchString(str) {
			updateArray := []int{}
			for _, page := range strings.Split(str, ",") {
				pageNum, _ := strconv.Atoi(page)
				updateArray = append(updateArray, pageNum)
			}

			valid, depI, pI := isValidUpdate(updateArray, rulesMap)
			if valid {
				validMiddleValuesSum += updateArray[len(updateArray)/2]
			} else {
				// swap invalid numbers until the array is valid
				for !valid {
					updateArray[depI], updateArray[pI] = updateArray[pI], updateArray[depI]
					valid, depI, pI = isValidUpdate(updateArray, rulesMap)
				}
				fixedMiddleValuesSum += updateArray[len(updateArray)/2]
			}
		}
	}

	fmt.Println(validMiddleValuesSum, fixedMiddleValuesSum)

}

// isValidUpdate checks if the update array satisfies the rules
func isValidUpdate(updateArray []int, rulesMap map[int][]int) (bool, int, int) {

	pageIndex := map[int]int{}

	// Create a map of page positions in the updateArray
	for i, page := range updateArray {
		pageIndex[page] = i
	}

	// Validate rules
	for page, dependencies := range rulesMap {
		for _, dependentPage := range dependencies {
			// Check if dependentPage exists and appears before the current page
			if _, exists := pageIndex[dependentPage]; exists {
				if pageIndex[dependentPage] < pageIndex[page] {
					return false, pageIndex[dependentPage], pageIndex[page]
				}
			}

		}
	}
	return true, 0, 0
}

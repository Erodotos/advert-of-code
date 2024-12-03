package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// Running this function alone gets you results for part1.
func getMulSum(text string) int {

	rePairs := regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	reNumbers := regexp.MustCompile(`\d{1,3}`)

	pairs := rePairs.FindAllString(text, -1)

	sum := 0

	for _, pair := range pairs {
		nums := reNumbers.FindAllString(pair, -1)

		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			log.Fatalln("Error converting string to int: ", err)
		}

		sum += num1 * num2
	}

	return sum
}

// Running this function alone gets you results for part2.
// If you want result for part1 take a look on function getMulSum()
func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error reading the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	total := 0
	str := ""

	// Concat all lines to form a valid program
	// concatenating line might be form a new
	// multiplication
	for fileScanner.Scan() {
		str += fileScanner.Text()
	}

	// Remove any multiplication pais between don't() and do()
	re := regexp.MustCompile(`don't\(\).*?do\(\)`)
	txt := re.ReplaceAllLiteralString(str, " ")

	// culculate result
	total += getMulSum(txt)

	fmt.Println(total)

}

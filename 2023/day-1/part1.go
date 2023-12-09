package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error reading input file.")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	sum := 0

	for fileScanner.Scan() {
		chars := []byte(fileScanner.Text())
		var num []byte
		for _, c := range chars {
			if c >= '0' && c <= '9' {
				num = append(num, c)
			}
		}
		rowResult, err := strconv.Atoi(fmt.Sprintf("%c%c", num[0], num[len(num)-1]))
		if err != nil {
			log.Fatal("Error converting string to integer.")
		}
		sum += rowResult
	}

	fmt.Println(sum)
}

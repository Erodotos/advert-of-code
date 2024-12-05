package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln("Error readininputData the file", err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)

	// Dump input data in byte array
	inputData := [][]byte{}
	for fileScanner.Scan() {
		inputData = append(inputData, []byte(fileScanner.Text()))
	}

	rows := len(inputData)
	columns := len(inputData[0])

	occurencesP1 := 0
	occurencesP2 := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < columns; c++ {
			if c+3 < columns && inputData[r][c] == 'X' && inputData[r][c+1] == 'M' && inputData[r][c+2] == 'A' && inputData[r][c+3] == 'S' {
				occurencesP1++
			}
			if r+3 < rows && inputData[r][c] == 'X' && inputData[r+1][c] == 'M' && inputData[r+2][c] == 'A' && inputData[r+3][c] == 'S' {
				occurencesP1++
			}
			if r+3 < rows && c+3 < columns && inputData[r][c] == 'X' && inputData[r+1][c+1] == 'M' && inputData[r+2][c+2] == 'A' && inputData[r+3][c+3] == 'S' {
				occurencesP1++
			}
			if c+3 < columns && inputData[r][c] == 'S' && inputData[r][c+1] == 'A' && inputData[r][c+2] == 'M' && inputData[r][c+3] == 'X' {
				occurencesP1++
			}
			if r+3 < rows && inputData[r][c] == 'S' && inputData[r+1][c] == 'A' && inputData[r+2][c] == 'M' && inputData[r+3][c] == 'X' {
				occurencesP1++
			}
			if r+3 < rows && c+3 < columns && inputData[r][c] == 'S' && inputData[r+1][c+1] == 'A' && inputData[r+2][c+2] == 'M' && inputData[r+3][c+3] == 'X' {
				occurencesP1++
			}
			if r-3 >= 0 && c+3 < columns && inputData[r][c] == 'S' && inputData[r-1][c+1] == 'A' && inputData[r-2][c+2] == 'M' && inputData[r-3][c+3] == 'X' {
				occurencesP1++
			}
			if r-3 >= 0 && c+3 < columns && inputData[r][c] == 'X' && inputData[r-1][c+1] == 'M' && inputData[r-2][c+2] == 'A' && inputData[r-3][c+3] == 'S' {
				occurencesP1++
			}

			if r+2 < rows && c+2 < columns && inputData[r][c] == 'M' && inputData[r+1][c+1] == 'A' && inputData[r+2][c+2] == 'S' && inputData[r+2][c] == 'M' && inputData[r][c+2] == 'S' {
				occurencesP2++
			}
			if r+2 < rows && c+2 < columns && inputData[r][c] == 'M' && inputData[r+1][c+1] == 'A' && inputData[r+2][c+2] == 'S' && inputData[r+2][c] == 'S' && inputData[r][c+2] == 'M' {
				occurencesP2++
			}
			if r+2 < rows && c+2 < columns && inputData[r][c] == 'S' && inputData[r+1][c+1] == 'A' && inputData[r+2][c+2] == 'M' && inputData[r+2][c] == 'M' && inputData[r][c+2] == 'S' {
				occurencesP2++
			}
			if r+2 < rows && c+2 < columns && inputData[r][c] == 'S' && inputData[r+1][c+1] == 'A' && inputData[r+2][c+2] == 'M' && inputData[r+2][c] == 'S' && inputData[r][c+2] == 'M' {
				occurencesP2++
			}

		}
	}

	fmt.Println(occurencesP1, occurencesP2)
}

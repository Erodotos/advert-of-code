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
		for index, c := range chars {
			if c >= '0' && c <= '9' {
				num = append(num, c)
			}
			if index+2 < len(chars) && c == 'o' {
				if chars[index+1] == 'n' && chars[index+2] == 'e' {
					num = append(num, '1')
				}
			}
			if index+2 < len(chars) && c == 't' {
				if chars[index+1] == 'w' && chars[index+2] == 'o' {
					num = append(num, '2')
				}
			}
			if index+4 < len(chars) && c == 't' {
				if chars[index+1] == 'h' && chars[index+2] == 'r' && chars[index+3] == 'e' && chars[index+4] == 'e' {
					num = append(num, '3')
				}
			}
			if index+3 < len(chars) && c == 'f' {
				if chars[index+1] == 'o' && chars[index+2] == 'u' && chars[index+3] == 'r' {
					num = append(num, '4')
				}
			}
			if index+3 < len(chars) && c == 'f' {
				if chars[index+1] == 'i' && chars[index+2] == 'v' && chars[index+3] == 'e' {
					num = append(num, '5')
				}
			}
			if index+2 < len(chars) && c == 's' {
				if chars[index+1] == 'i' && chars[index+2] == 'x' {
					num = append(num, '6')
				}
			}
			if index+4 < len(chars) && c == 's' {
				if chars[index+1] == 'e' && chars[index+2] == 'v' && chars[index+3] == 'e' && chars[index+4] == 'n' {
					num = append(num, '7')
				}
			}
			if index+4 < len(chars) && c == 'e' {
				if chars[index+1] == 'i' && chars[index+2] == 'g' && chars[index+3] == 'h' && chars[index+4] == 't' {
					num = append(num, '8')
				}
			}
			if index+3 < len(chars) && c == 'n' {
				if chars[index+1] == 'i' && chars[index+2] == 'n' && chars[index+3] == 'e' {
					num = append(num, '9')
				}
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

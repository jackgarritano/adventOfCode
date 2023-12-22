package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isDigit(c uint8) bool {
	return c >= '0' && c <= '9'
}

func getValuePart1(line string) int {
	front := ' '
	back := ' '
	for i := 0; i < len(line); i++ {
		if front == ' ' && isDigit(line[i]) {
			front = rune(line[i])
		}
		if back == ' ' && isDigit(line[len(line)-i-1]) {
			back = rune(line[len(line)-i-1])
		}
	}
	value := 10*int(front-'0') + int(back-'0')
	return value
}

var numStrs = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func parseNumStrEndAdd(s string) (rune, bool) {
	for i := 0; i < len(s); i++ {
		if numChar, exists := numStrs[s[i:len(s)]]; exists {
			return numChar, true
		}
	}
	return ' ', false
}

func parseNumStrBeginAdd(s string) (rune, bool) {
	for i := 1; i <= len(s); i++ {
		if numChar, exists := numStrs[s[0:i]]; exists {
			return numChar, true
		}
	}
	return ' ', false
}

func getValuePart2(line string) int {
	front := ' '
	back := ' '
	frontStr := ""
	endStr := ""
	for i := 0; i < len(line); i++ {
		frontStr += string(line[i])
		endStr = string(line[len(line)-i-1]) + endStr
		if front == ' ' && isDigit(line[i]) {
			front = rune(line[i])
			frontStr = ""

		} else if numChar, exists := parseNumStrEndAdd(frontStr); front == ' ' && exists {
			front = numChar
		}
		if back == ' ' && isDigit(line[len(line)-i-1]) {
			back = rune(line[len(line)-i-1])
			endStr = ""
		} else if numChar, exists := parseNumStrBeginAdd(endStr); back == ' ' && exists {
			back = numChar
		}
	}
	value := 10*int(front-'0') + int(back-'0')
	return value
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	values := make([]int, 0)

	for scanner.Scan() {
		//scanner.Scan()
		line := scanner.Text()
		value := getValuePart2(line)
		values = append(values, value)
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("sum: ", sum)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func isDigit(c uint8) bool {
	fmt.Println("isDigit: ", c >= '0' && c <= '9')
	return c >= '0' && c <= '9'
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
		line := scanner.Text()
		fmt.Println("line: ", line)
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
		fmt.Println("front: ", front)
		fmt.Println("back: ", back)
		value := 10*int(front-'0') + int(back-'0')
		fmt.Println("value: ", value)
		values = append(values, value)
	}

	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Println("sum: ", sum)
}

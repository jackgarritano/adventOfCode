package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func isSymbol(r rune) bool {
	return r != '.' && (r < '0' || r > '9')
}

// calling this where [row][col] is a symbol and row or col is 0 will incorrectly return true
// (shouldn't affect correctness of solution)
func hasAdjSymbol(schem *[]string, row, col int) bool {
	l := len((*schem)[row]) - 1
	return isSymbol(rune((*schem)[max(row-1, 0)][col])) ||
		isSymbol(rune((*schem)[min(row+1, l)][col])) ||
		isSymbol(rune((*schem)[row][max(col-1, 0)])) ||
		isSymbol(rune((*schem)[row][min(col+1, l)])) ||
		isSymbol(rune((*schem)[max(row-1, 0)][max(col-1, 0)])) ||
		isSymbol(rune((*schem)[max(row-1, 0)][min(col+1, l)])) ||
		isSymbol(rune((*schem)[min(row+1, l)][max(col-1, 0)])) ||
		isSymbol(rune((*schem)[min(row+1, l)][min(col+1, l)]))
}

func getFullNum(schemRow string, ind int) (int, int) {
	frontInd := ind
	for frontInd > 0 && schemRow[frontInd-1] >= '0' && schemRow[frontInd-1] <= '9' {
		frontInd--
	}

	backInd := ind
	for backInd < len(schemRow)-1 && schemRow[backInd+1] >= '0' && schemRow[backInd+1] <= '9' {
		backInd++
	}

	fullNum, _ := strconv.ParseInt(schemRow[frontInd:backInd+1], 10, strconv.IntSize)

	return int(fullNum), backInd
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	schem := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		schem = append(schem, line)
	}

	sum := 0
	for rowInd, schemRow := range schem {
		for i := 0; i < len(schemRow); i++ {
			charNum := schemRow[i]
			if charNum >= '0' && charNum <= '9' && hasAdjSymbol(&schem, rowInd, i) {
				num, newInd := getFullNum(schemRow, i)
				sum += num
				i = newInd
			}
		}
	}
	fmt.Println("sum: ", sum)
}

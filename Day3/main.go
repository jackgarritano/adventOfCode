package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type star struct {
	row int
	col int
}

func isSymbol(r rune) bool {
	return r != '.' && (r < '0' || r > '9')
}

// calling this where [row][col] is a symbol and row or col += 1 is out of bounds will
// incorrectly return true (shouldn't affect correctness of solution)
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

func findAdjStars(schem *[]string, row, col int) []star {
	isInRange := func(row, col int) bool {
		return row >= 0 && row < len(*schem) && col >= 0 && col < len((*schem)[row])
	}

	stars := make([]star, 0)
	for newRow := -1; newRow <= 1; newRow++ {
		for newCol := -1; newCol <= 1; newCol++ {
			if row+newRow == 3 && col+newCol == 24 {
				fmt.Println("it's in range: ", isInRange(row+newRow, col+newCol))
				fmt.Println("it's a star: ", (*schem)[row+newRow][col+newCol] == '*')
			}
			if isInRange(row+newRow, col+newCol) && (*schem)[row+newRow][col+newCol] == '*' {
				stars = append(stars, star{row: row + newRow, col: col + newCol})
			}
		}
	}

	return stars
}

func getFullNum(schemRow string, ind int) (int, int, int) {
	frontInd := ind
	for frontInd > 0 && schemRow[frontInd-1] >= '0' && schemRow[frontInd-1] <= '9' {
		frontInd--
	}

	backInd := ind
	for backInd < len(schemRow)-1 && schemRow[backInd+1] >= '0' && schemRow[backInd+1] <= '9' {
		backInd++
	}

	fullNum, _ := strconv.ParseInt(schemRow[frontInd:backInd+1], 10, strconv.IntSize)

	return int(fullNum), frontInd, backInd
}

// part 2 idea: for each number, find all adj stars and hold on to each star's coordinates
// and the number that was adj to it, get all the stars w/ 2 numbers in the end
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

	starsToNums := make(map[star][]int)
	//iterate through each row
	for rowInd, schemRow := range schem {
		//iterate through each char in a row
		for i := 0; i < len(schemRow); i++ {
			charNum := schemRow[i]
			if charNum >= '0' && charNum <= '9' {
				num, frontInd, backInd := getFullNum(schemRow, i)
				set := make(map[star]struct{})
				//get all stars adjacent to any part of the found number
				for j := frontInd; j <= backInd; j++ {
					if j == 31 && rowInd == 74 {
						fmt.Println("74 31")
					}
					stars := findAdjStars(&schem, rowInd, j)
					//add each star to set to remove repeated stars
					for _, foundStar := range stars {
						set[foundStar] = struct{}{}
					}
				}
				for foundStar := range set {
					starsToNums[foundStar] = append(starsToNums[foundStar], num)
					// } else {
					//	starsToNums[foundStar] =
					//}
				}
				i = backInd
			}
		}
	}
	fmt.Println("starsToNums: ", starsToNums)
	starsToDel := make([]star, 0)
	for foundStar := range starsToNums {
		if len(starsToNums[foundStar]) != 2 {
			starsToDel = append(starsToDel, foundStar)
		}
	}
	for _, starToDel := range starsToDel {
		delete(starsToNums, starToDel)
	}
	fmt.Println("filtered starsToNums: ", starsToNums)
	sum := 0
	for foundStar := range starsToNums {
		if len(starsToNums[foundStar]) == 2 {
			product := 1
			for _, num := range starsToNums[foundStar] {
				product *= num
			}
			sum += product
		}
	}

	fmt.Println("sum: ", sum)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getScanner(path string) (*bufio.Scanner, func() error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file), file.Close
}

type card struct {
	cardNum     int
	nums        []int
	winningNums []int
}

func parseCard(line string) card {
	parseInt := func(s string) (int, bool) {
		i, err := strconv.ParseInt(strings.Trim(s, " "), 10, strconv.IntSize)
		if err != nil {
			return 0, false
		}
		return int(i), true
	}
	cardParts := strings.Split(line, ": ")
	cardNum, _ := parseInt(cardParts[0][strings.Index(cardParts[0], " ")+1:])
	numParts := strings.Split(cardParts[1], " | ")
	winningNumStrs := strings.Split(numParts[0], " ")
	numStrs := strings.Split(numParts[1], " ")

	winningNums := make([]int, 0)
	for _, winningNumStr := range winningNumStrs {
		if n, ok := parseInt(winningNumStr); ok {
			winningNums = append(winningNums, n)
		}
	}

	nums := make([]int, 0)
	for _, numStr := range numStrs {
		if n, ok := parseInt(numStr); ok {
			nums = append(nums, n)
		}
	}

	return card{
		cardNum,
		nums,
		winningNums,
	}
}

func (c *card) getPoints() int {
	points := 0
	for _, winningNum := range c.winningNums {
		if slices.Contains(c.nums, winningNum) {
			points = max(1, points*2)
		}
	}
	return points
}

func (c *card) getNumMatching() int {
	numMatching := 0
	for _, winningNum := range c.winningNums {
		if slices.Contains(c.nums, winningNum) {
			numMatching++
		}
	}
	return numMatching
}

func main() {
	scanner, closeF := getScanner("./input.txt")
	defer closeF()
	cards := make([]card, 0)
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseCard(line))
	}

	//part 1
	//totalPoints := 0
	//for _, iCard := range cards {
	//	totalPoints += iCard.getPoints()
	//}

	matchingMap := make(map[int]int)
	for _, iCard := range cards {
		matchingMap[iCard.cardNum] = iCard.getNumMatching()
	}

	getWonCopies := func(gameNum, copiesWon int) []int {
		wonCopies := make([]int, 0)
		for i := gameNum + 1; i <= gameNum+copiesWon; i++ {
			wonCopies = append(wonCopies, i)
		}
		return wonCopies
	}

	totalCopies := 0
	copiesWon := make([]int, 0)
	for gameNum := range matchingMap {
		totalCopies++
		copiesWon = append(copiesWon, getWonCopies(gameNum, matchingMap[gameNum])...)
	}

	for len(copiesWon) > 0 {
		totalCopies++
		curCopy := copiesWon[0]
		copiesWon = copiesWon[1:]
		copiesWon = append(copiesWon, getWonCopies(curCopy, matchingMap[curCopy])...)
	}

	fmt.Println("totalCopies: ", totalCopies)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type draw map[string]uint

var cubeLimits = draw{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	draws   []draw
	gameNum uint
}

// implements part 1
func (g *game) isValid() bool {
	for _, drawMap := range g.draws {
		for color := range drawMap {
			if limit, _ := cubeLimits[color]; limit < drawMap[color] {
				return false
			}
		}
	}
	return true
}

func (g *game) getMinimumValidPower() uint {
	minValid := make(draw)
	for _, drawMap := range g.draws {
		for color := range drawMap {
			minValid[color] = max(drawMap[color], minValid[color])
		}
	}

	power := uint(1)
	for _, minNum := range minValid {
		power *= minNum
	}

	return power
}

func parseGame(s string) (uint, []draw) {
	gameNumAndData := strings.Split(s, ": ")

	gameNumStr := gameNumAndData[0][strings.Index(gameNumAndData[0], " ")+1:]
	gameNum, _ := strconv.ParseUint(gameNumStr, 10, strconv.IntSize)

	drawStrs := strings.Split(gameNumAndData[1], "; ")
	var draws []draw
	for _, drawStr := range drawStrs {
		drawMap := make(draw)
		cubeStrs := strings.Split(drawStr, ", ")
		for _, cube := range cubeStrs {
			colorAndNum := strings.Split(cube, " ")
			numColor, _ := strconv.ParseUint(colorAndNum[0], 10, strconv.IntSize)
			drawMap[colorAndNum[1]] = uint(numColor)
		}
		draws = append(draws, drawMap)
	}

	return uint(gameNum), draws
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := uint(0)
	for scanner.Scan() {
		line := scanner.Text()
		currentGame := game{}
		currentGame.gameNum, currentGame.draws = parseGame(line)
		//if currentGame.isValid() {
		//	sum += currentGame.gameNum
		//}
		sum += currentGame.getMinimumValidPower()
	}
	fmt.Println("sum: ", sum)
}

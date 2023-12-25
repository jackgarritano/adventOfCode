package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func parseInt(s string) (int, bool) {
	s = strings.Trim(s, " ")
	result, err := strconv.ParseInt(s, 10, strconv.IntSize)
	if err != nil {
		return 0, false
	}
	return int(result), true
}

func parseInput() []race {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	times := make([]int, 0)

	scanner.Scan()
	timeLine := scanner.Text()
	timeData := timeLine[strings.Index(timeLine, ":")+1:]
	//this line is for part 2
	timeData = strings.ReplaceAll(timeData, " ", "")
	for _, timeStr := range strings.Split(timeData, " ") {
		if time, ok := parseInt(timeStr); ok {
			times = append(times, time)
		}
	}

	dists := make([]int, 0)

	scanner.Scan()
	distLine := scanner.Text()
	distData := distLine[strings.Index(distLine, ":")+1:]
	//this line is for part 2
	distData = strings.ReplaceAll(distData, " ", "")
	for _, distStr := range strings.Split(distData, " ") {
		if dist, ok := parseInt(distStr); ok {
			dists = append(dists, dist)
		}
	}

	races := make([]race, 0, len(times))
	for ind, time := range times {
		races = append(races, race{time: time, dist: dists[ind]})
	}

	return races
}

func (r *race) getWinningPresses() []int {
	presses := make([]int, 0)
	counter := 0
	for p := 1; p < r.time; p++ {
		if distance := p * (r.time - p); distance > r.dist {
			presses = append(presses, p)
		}
		counter++
		if counter%1000000 == 0 {
			fmt.Println(counter/1000000, " million")
		}
	}
	return presses
}

func main() {
	races := parseInput()

	fmt.Println("races: ", races)

	product := 1
	for _, raceI := range races {
		winningPresses := raceI.getWinningPresses()
		product *= len(winningPresses)
	}

	fmt.Println("product: ", product)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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

type alm struct {
	from   string
	to     string
	ranges []struct {
		fromStart int
		toStart   int
		offset    int
	}
}

func parseInt(s string) (int, bool) {
	i, err := strconv.ParseInt(strings.Trim(s, " "), 10, strconv.IntSize)
	if err != nil {
		return 0, false
	}
	return int(i), true
}

func parseInput() (map[string][]int, []alm) {
	parseInitialList := func(s *bufio.Scanner) map[string][]int {
		line := s.Text()
		categoryAndParts := strings.Split(line, ": ")
		category := categoryAndParts[0]
		partStrs := strings.Split(categoryAndParts[1], " ")
		parts := make([]int, 0)
		for _, partStr := range partStrs {
			if i, ok := parseInt(partStr); ok {
				parts = append(parts, i)
			}
		}
		return map[string][]int{
			category: parts,
		}
	}
	parseMap := func(s *bufio.Scanner) alm {
		toAndFrom := strings.Split(strings.Split(s.Text(), " ")[0], "-")
		to := toAndFrom[2]
		from := toAndFrom[0]
		ranges := make([]struct {
			fromStart int
			toStart   int
			offset    int
		}, 0)
		for s.Scan() && s.Text() != "" {
			data := strings.Split(s.Text(), " ")
			toStart, _ := parseInt(data[0])
			fromStart, _ := parseInt(data[1])
			offset, _ := parseInt(data[2])
			ranges = append(ranges, struct {
				fromStart int
				toStart   int
				offset    int
			}{
				fromStart,
				toStart,
				offset,
			})
		}
		return alm{
			from,
			to,
			ranges,
		}
	}

	scanner, closeF := getScanner("./input.txt")
	defer closeF()

	initialVals := make(map[string][]int)
	alms := make([]alm, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Index(line, "-") >= 0 {
			alms = append(alms, parseMap(scanner))
		} else if strings.Index(line, ":") >= 0 {
			initialVals = parseInitialList(scanner)
		}
	}
	return initialVals, alms
}

func (a *alm) calcNext(initial int) (int, string) {
	nextAlm := a.to
	if nextAlm == "location" {
		nextAlm = ""
	}
	for _, r := range a.ranges {
		if initial >= r.fromStart && initial < r.fromStart+r.offset {
			return r.toStart + (initial - r.fromStart), nextAlm
		}
	}
	return initial, nextAlm
}

func main() {
	initialVals, alms := parseInput()
	almMap := make(map[string]*alm)
	for ind := range alms {
		almMap[alms[ind].from] = &alms[ind]
	}

	minVal := math.MaxInt
	for _, val := range initialVals["seeds"] {
		nextAlm := ""
		val, nextAlm = almMap["seed"].calcNext(val)
		for nextAlm != "" {
			val, nextAlm = almMap[nextAlm].calcNext(val)
		}
		minVal = min(minVal, val)
	}

	fmt.Println("minVal: ", minVal)
}

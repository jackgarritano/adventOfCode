package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func parseInt(s string) (int, bool) {
	s = strings.Trim(s, " ")
	result, err := strconv.ParseInt(s, 10, strconv.IntSize)
	if err != nil {
		return 0, false
	}
	return int(result), true
}

type hand struct {
	cards    []int
	strength int
	bid      int
}

func getNumDuplicates(item int, list *[]int) int {
	numDuplicates := 0
	for _, x := range *list {
		if item == x {
			numDuplicates++
		}
	}
	return numDuplicates
}

func (c *hand) addStrength() {
	tagToOccs := make(map[int]int)
	numJokers := 0
	for _, card := range c.cards {
		if card == 1 {
			numJokers++
		} else {
			tagToOccs[card] = getNumDuplicates(card, &c.cards)
		}
	}
	orderedOccs := make([]int, 0, len(tagToOccs))
	for tag := range tagToOccs {
		orderedOccs = append(orderedOccs, tagToOccs[tag])
	}
	slices.SortFunc(orderedOccs, func(a, b int) int {
		return b - a
	})
	if len(orderedOccs) > 0 {
		orderedOccs[0] += numJokers
	} else {
		orderedOccs = append(orderedOccs, numJokers)
	}

	if len(orderedOccs) >= 1 && orderedOccs[0] == 5 {
		c.strength = 7
	} else if len(orderedOccs) >= 1 && orderedOccs[0] == 4 {
		c.strength = 6
	} else if len(orderedOccs) >= 2 && orderedOccs[0] == 3 && orderedOccs[1] == 2 {
		c.strength = 5
	} else if len(orderedOccs) >= 1 && orderedOccs[0] == 3 {
		c.strength = 4
	} else if len(orderedOccs) >= 2 && orderedOccs[0] == 2 && orderedOccs[1] == 2 {
		c.strength = 3
	} else if len(orderedOccs) >= 1 && orderedOccs[0] == 2 {
		c.strength = 2
	} else {
		c.strength = 1
	}
}

func newHand(cardStrs string, bidStr string) *hand {
	cards := make([]int, 0, 5)
	for _, cardStr := range cardStrs {
		if result, ok := parseInt(string(cardStr)); ok {
			cards = append(cards, result)
		} else {
			switch string(cardStr) {
			case "T":
				cards = append(cards, 10)
			case "J":
				cards = append(cards, 1)
			case "Q":
				cards = append(cards, 12)
			case "K":
				cards = append(cards, 13)
			case "A":
				cards = append(cards, 14)
			}
		}
	}

	bid, _ := parseInt(bidStr)

	c := &hand{
		cards: cards,
		bid:   bid,
	}
	c.addStrength()

	return c
}

func parseInput() handList {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make(handList, 0)

	for scanner.Scan() {
		line := scanner.Text()
		handParts := strings.Split(line, " ")
		hands = append(hands, *newHand(handParts[0], handParts[1]))
	}
	return hands
}

type handList []hand

// Implement the Len method required by sort.Interface
func (h handList) Len() int {
	return len(h)
}

// Implement the Less method required by sort.Interface
func (h handList) Less(i, j int) bool {
	if h[i].strength != h[j].strength {
		return h[i].strength < h[j].strength
	}
	for c := 0; c < 5; c++ {
		if h[i].cards[c] != h[j].cards[c] {
			return h[i].cards[c] < h[j].cards[c]
		}
	}
	return false
}

// Implement the Swap method required by sort.Interface
func (h handList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	hands := parseInput()
	sort.Sort(hands)
	sum := 0
	for ind, handI := range hands {
		sum += (ind + 1) * handI.bid
	}
	fmt.Println("sum: ", sum)
}

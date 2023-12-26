package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type node struct {
	left  string
	right string
}

type nodes struct {
	dirMap map[string]node
}

func parseInput() (string, nodes) {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parsedDirs := scanner.Text()
	parsedNodes := nodes{dirMap: make(map[string]node)}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		parsedNode := node{
			left:  line[7:10],
			right: line[12:15],
		}
		parsedNodes.dirMap[line[0:3]] = parsedNode
	}
	return parsedDirs, parsedNodes
}

func (ns *nodes) getNext(n string, d uint8) string {
	if d == 'L' {
		return ns.dirMap[n].left
	} else {
		return ns.dirMap[n].right
	}
}

func main() {
	directions, parsedNodes := parseInput()

	counter := 0
	for curNode, i := "AAA", 0; curNode != "ZZZ"; i = (i + 1) % len(directions) {
		curNode = parsedNodes.getNext(curNode, directions[i])
		counter++
	}
	fmt.Println("counter: ", counter)
}

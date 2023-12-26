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
	directions string
	dirMap     map[string]node
}

func parseInput() nodes {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	parsedDirs := scanner.Text()
	parsedNodes := nodes{directions: parsedDirs,
		dirMap: make(map[string]node)}

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		parsedNode := node{
			left:  line[7:10],
			right: line[12:15],
		}
		parsedNodes.dirMap[line[0:3]] = parsedNode
	}
	return parsedNodes
}

func (ns *nodes) getNext(n string, d uint8) string {
	if d == 'L' {
		return ns.dirMap[n].left
	} else {
		return ns.dirMap[n].right
	}
}

// return length to first z and length of cycle
func (ns *nodes) getCycleInfo(aNode string) (int, int) {
	initialLength := 0
	i := 0
	for ; aNode[2] != 'Z'; i = (i + 1) % len(ns.directions) {
		aNode = ns.getNext(aNode, ns.directions[i])
		initialLength++
	}

	zNode := aNode
	aNode = ns.getNext(aNode, ns.directions[i])
	cycleLength := 1
	for i = (i + 1) % len(ns.directions); aNode != zNode; i = (i + 1) % len(ns.directions) {
		if aNode[2] == 'Z' {
			fmt.Println("different z in cycle: ", aNode)
		}
		aNode = ns.getNext(aNode, ns.directions[i])
		cycleLength++
	}
	return initialLength, cycleLength
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	parsedNodes := parseInput()

	aNodes := make([]string, 0)
	for key := range parsedNodes.dirMap {
		if key[2] == 'A' {
			aNodes = append(aNodes, key)
		}
	}

	cycleLens := make([]int, 0)
	for i := 0; i < len(aNodes); i++ {
		_, cycleLength := parsedNodes.getCycleInfo(aNodes[i])
		cycleLens = append(cycleLens, cycleLength)
	}
	fmt.Println("LCM: ", LCM(cycleLens[0], cycleLens[1], cycleLens[2:]...))
}

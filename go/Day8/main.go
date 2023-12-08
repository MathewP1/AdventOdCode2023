package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type solution struct {
	name     string
	filename string
	f        func(filename string) int
}

type NodeId [3]byte

func newNodeId(s string) NodeId {
	return NodeId([]byte(s))
}

type NodeTree struct {
	nodeMap    map[NodeId][2]NodeId
	start, end []NodeId
}

func newNodeTree(lines []string) *NodeTree {
	nodeTree := &NodeTree{}
	nodeTree.nodeMap = make(map[NodeId][2]NodeId, len(lines))
	for _, l := range lines {
		nodeName := NodeId([]byte(l[:3]))
		leftChild := NodeId([]byte(l[7:10]))
		rightChild := NodeId([]byte(l[12:15]))
		nodeTree.nodeMap[nodeName] = [2]NodeId{leftChild, rightChild}
	}
	return nodeTree
}

type Instructions struct {
	turns []byte
}

func newInstructions(s string) *Instructions {
	return &Instructions{[]byte(s)}
}

func part1(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)

	}

	instr := newInstructions(lines[0])
	fd := newNodeTree(lines[1:])
	if fd == nil {
		panic("tree creation failed")
	}

	currentId := newNodeId("AAA")
	endId := newNodeId("ZZZ")

	steps := 0
	for currentId != endId {
		for _, i := range instr.turns {
			turn := 0
			if i == 'R' {
				turn = 1
			}
			currentId = fd.nodeMap[currentId][turn]
			steps += 1
		}
	}
	return steps
}

func part2(filename string) int {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}

	instr := newInstructions(lines[0])
	nodeLines := lines[1:]
	nodes := &NodeTree{}
	nodes.nodeMap = make(map[NodeId][2]NodeId)
	for _, l := range nodeLines {
		if l == "" {
			continue
		}
		nodeName := NodeId([]byte(l[:3]))
		leftChild := NodeId([]byte(l[7:10]))
		rightChild := NodeId([]byte(l[12:15]))

		n := l[2]
		if n == 'A' {
			nodes.start = append(nodes.start, nodeName)
		} else if n == 'Z' {
			nodes.end = append(nodes.end, nodeName)
		}
		nodes.nodeMap[nodeName] = [2]NodeId{leftChild, rightChild}
	}

	current := nodes.start
	steps := 0
    fmt.Println("Starting nodes len: ", len(current))
    
	for true {
		nextTurn := instr.turns[steps%len(instr.turns)]
		steps += 1
		turn := 0
		if nextTurn == 'R' {
			turn = 1
		}
        z := 0
		for i, c := range current {
			nextNode := nodes.nodeMap[c][turn]
            if nextNode[2] == 'Z' {
                z += 1
            }
			current[i] = nextNode
		}
        if z == len(current) {
            break
        }
	}

	return steps
}

func main() {

	solutions := []solution{
		{"Part1 test input", "test.txt", part1},
		{"Part1 real input", "input.txt", part1},
		{"Part2 test input", "test2.txt", part2},
        {"part 2 real input", "input.txt", part2},
	}

	for _, s := range solutions {
		fmt.Println("Starting ", s.name)
		timeStart := time.Now()
		fmt.Println("Result: ", s.f(s.filename))
		fmt.Println("Took: ", time.Since(timeStart).Microseconds())
	}
}

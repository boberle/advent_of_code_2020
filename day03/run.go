package day03

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Point struct {
	x, y int
}

type GameBoard struct {
	trees         map[Point]struct{}
	width, height int
}

type slopeList []struct {
	right, down int
}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalf("can't open file '%s'", infile)
	}
	defer fh.Close()

	board := parseFile(fh)

	fmt.Printf("Part 1: nb of trees encountered: %d\n", board.traverseBoardAndCountTrees(3, 1))

	slopes := slopeList{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	fmt.Printf("Part 2: nb of trees encountered: %d\n", board.checkAllSlopes(slopes))
}

func parseFile(reader io.Reader) GameBoard {
	scanner := bufio.NewScanner(reader)

	var board GameBoard
	board.trees = make(map[Point]struct{})

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, r := range line {
			if r == '#' {
				board.trees[Point{x, y}] = struct{}{}
			}
			if y == 0 {
				board.width++
			}
		}
		y++
	}
	board.height = y
	return board
}

func (board *GameBoard) traverseBoardAndCountTrees(right, down int) int {
	var counter, x, y int
	for {
		x = (x + right) % board.width
		y += down
		if y >= board.height {
			return counter
		}
		if _, found := board.trees[Point{x, y}]; found {
			counter++
		}
	}
}

func (board *GameBoard) checkAllSlopes(slopes slopeList) int {
	rv := 1
	for _, slope := range slopes {
		rv *= board.traverseBoardAndCountTrees(slope.right, slope.down)
	}
	return rv
}

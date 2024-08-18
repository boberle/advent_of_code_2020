package day12

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Pos struct {
	x int
	y int
}

type Direction int

const (
	N Direction = 0
	E Direction = 1
	S Direction = 2
	W Direction = 3
	R Direction = iota
	L Direction = iota
	F Direction = iota
)

type Instruction struct {
	action Direction
	value  int
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	instructions := parseFile(fh)

	pos := walk(instructions)
	distance := computeManhattanDistance(pos)
	fmt.Printf("Part 1: Manhattan distance: %d (x=%d, y=%d)\n", distance, pos.x, pos.y)

	pos = walkWithWaypoint(instructions)
	distance = computeManhattanDistance(pos)
	fmt.Printf("Part 2: Manhattan distance: %d (x=%d, y=%d)\n", distance, pos.x, pos.y)

}

func parseFile(reader io.Reader) []Instruction {
	instructions := []Instruction{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		var action Direction
		switch scanner.Text()[0] {
		case 'N':
			action = N
		case 'E':
			action = E
		case 'S':
			action = S
		case 'W':
			action = W
		case 'R':
			action = R
		case 'L':
			action = L
		case 'F':
			action = F
		}
		value, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, Instruction{action, value})
	}

	return instructions
}

func move(direction Direction, value int, pos Pos) Pos {
	switch direction {
	case N:
		pos.y -= value
	case E:
		pos.x += value
	case S:
		pos.y += value
	case W:
		pos.x -= value
	}
	return pos
}

func walk(instructions []Instruction) Pos {
	var currentPos Pos
	var currentDir Direction
	currentDir = E

	for _, instr := range instructions {
		switch instr.action {
		case N:
			fallthrough
		case E:
			fallthrough
		case S:
			fallthrough
		case W:
			currentPos = move(instr.action, instr.value, currentPos)
		case F:
			currentPos = move(currentDir, instr.value, currentPos)
		case R:
			currentDir = Direction(modulo(int(currentDir)+instr.value/90, 4))
		case L:
			currentDir = Direction(modulo(int(currentDir)-instr.value/90, 4))
		}
	}
	return currentPos
}

func walkWithWaypoint(instructions []Instruction) Pos {
	var shipPos Pos
	var waypointPos Pos

	waypointPos = Pos{10, -1}

	for _, instr := range instructions {
		switch instr.action {
		case N:
			fallthrough
		case E:
			fallthrough
		case S:
			fallthrough
		case W:
			waypointPos = move(instr.action, instr.value, waypointPos)
		case F:
			shipPos = move(E, waypointPos.x*instr.value, shipPos)
			shipPos = move(S, waypointPos.y*instr.value, shipPos)
		case R:
			waypointPos = computeCoordinatesAfterRotation(waypointPos, float64(instr.value)*-1)
		case L:
			waypointPos = computeCoordinatesAfterRotation(waypointPos, float64(instr.value))
		}
	}
	return shipPos
}

func modulo(a, b int) int {
	return (a%b + b) % b
}

func computeManhattanDistance(pos Pos) int {
	return int(math.Round(math.Abs(float64(pos.x)) + math.Abs(float64(pos.y))))
}

func computeCoordinatesAfterRotation(point Pos, rot float64) Pos {
	rotRad := rot * math.Pi / 180

	x := int(math.Round(float64(point.x)*math.Cos(rotRad) + float64(point.y)*math.Sin(rotRad)))
	y := int(math.Round(float64(point.y)*math.Cos(rotRad) - float64(point.x)*math.Sin(rotRad)))

	return Pos{x, y}
}

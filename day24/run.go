package day24

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type Direction int

const (
	East      Direction = iota
	West      Direction = iota
	SouthEast Direction = iota
	SouthWest Direction = iota
	NorthEast Direction = iota
	NorthWest Direction = iota
)

type Path []Direction

type Position struct {
	x float32
	y float32
}

type Offset struct {
	x float32
	y float32
}

type TileList map[Position]struct{}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	paths := parseFile(fh)
	blackTiles := computeBlackTiles(paths)
	fmt.Printf("Number of black tiles (part 1): %d\n", len(blackTiles))

	blackTilesAfter100days := flipTilesNTimes(blackTiles, 100)
	fmt.Printf("Number of black tiles after 100 days (part 2): %d\n", len(blackTilesAfter100days))
}

func parseFile(reader io.Reader) []Path {
	paths := make([]Path, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		path := make(Path, 0)
		i := 0
		for i < len(line) {
			if line[i] == 'e' {
				path = append(path, East)
				i++
			} else if line[i] == 'w' {
				path = append(path, West)
				i++
			} else if line[i] == 's' {
				if line[i+1] == 'e' {
					path = append(path, SouthEast)
				} else if line[i+1] == 'w' {
					path = append(path, SouthWest)
				}
				i += 2
			} else if line[i] == 'n' {
				if line[i+1] == 'e' {
					path = append(path, NorthEast)
				} else if line[i+1] == 'w' {
					path = append(path, NorthWest)
				}
				i += 2
			} else {
				panic(fmt.Sprintf("unknown character: %v", line[i]))
			}
		}
		paths = append(paths, path)
	}

	return paths
}

func (path *Path) getTargetPosition() Position {
	var x, y float32
	for _, direction := range *path {
		switch direction {
		case East:
			x++
		case West:
			x--
		case SouthEast:
			y++
			x += 0.5
		case SouthWest:
			y++
			x -= 0.5
		case NorthEast:
			y--
			x += 0.5
		case NorthWest:
			y--
			x -= 0.5
		}
	}
	return Position{x, y}
}

func computeBlackTiles(paths []Path) TileList {
	blackTiles := map[Position]struct{}{}
	for _, path := range paths {
		pos := path.getTargetPosition()
		if _, found := blackTiles[pos]; found {
			delete(blackTiles, pos)
		} else {
			blackTiles[pos] = struct{}{}
		}
	}
	return blackTiles
}

func (position *Position) shift(offset Offset) Position {
	return Position{
		x: position.x + offset.x,
		y: position.y + offset.y,
	}
}

func flipTilesNTimes(blackTiles TileList, n int) TileList {
	for i := 0; i < n; i++ {
		blackTiles = flipTiles(blackTiles)
	}
	return blackTiles
}

func flipTiles(blackTiles TileList) TileList {
	offsets := []Offset{
		{1, 0},
		{-1, 0},
		{0.5, 1},
		{-0.5, 1},
		{0.5, -1},
		{-0.5, -1},
	}
	newBlackTiles := make(TileList, 0)
	whiteTiles := map[Position]int{}
	for position := range blackTiles {
		surroundingBlackTiles := 0
		for _, offset := range offsets {
			surroundingTile := position.shift(offset)
			if _, found := blackTiles[surroundingTile]; found {
				surroundingBlackTiles++
			} else {
				whiteTiles[surroundingTile]++
			}
		}
		if !(surroundingBlackTiles == 0 || surroundingBlackTiles > 2) {
			newBlackTiles[position] = struct{}{}
		}
	}

	for position, surroundingBlackTileCount := range whiteTiles {
		if surroundingBlackTileCount == 2 {
			if _, found := newBlackTiles[position]; found {
				log.Fatalf("tile %v already black", position)
			}
			newBlackTiles[position] = struct{}{}
		}
	}
	return newBlackTiles
}

package day17

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Cube struct {
	x int
	y int
	z int
}

type Hypercube struct {
	x int
	y int
	z int
	w int
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	activeCubes := parseFile(fh)
	activeHypercubes := buildHypercubesFromCubes(activeCubes)

	activeCubes = cycleNTimes(activeCubes, 6)
	fmt.Printf("active cubes after 6 cycles (part 1): %d\n", len(activeCubes))

	activeHypercubes = cycleNTimes(activeHypercubes, 6)
	fmt.Printf("active hypercubes after 6 cycles (part 2): %d\n", len(activeHypercubes))

}

func parseFile(reader io.Reader) map[Cube]struct{} {
	activeCubes := map[Cube]struct{}{}

	scanner := bufio.NewScanner(reader)
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char == '#' {
				cube := Cube{x, y, 0}
				activeCubes[cube] = struct{}{}
			}
		}
		y++
	}
	return activeCubes
}

func (cube Cube) getNeighbors() []Cube {
	cubes := []Cube{}
	for zOffset := -1; zOffset <= 1; zOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			for xOffset := -1; xOffset <= 1; xOffset++ {
				if !(xOffset == 0 && yOffset == 0 && zOffset == 0) {
					cubes = append(cubes, Cube{cube.x + xOffset, cube.y + yOffset, cube.z + zOffset})
				}
			}
		}
	}
	return cubes
}

func (cube Hypercube) getNeighbors() []Hypercube {
	cubes := []Hypercube{}
	for zOffset := -1; zOffset <= 1; zOffset++ {
		for yOffset := -1; yOffset <= 1; yOffset++ {
			for xOffset := -1; xOffset <= 1; xOffset++ {
				for wOffset := -1; wOffset <= 1; wOffset++ {
					if !(xOffset == 0 && yOffset == 0 && zOffset == 0 && wOffset == 0) {
						cubes = append(cubes, Hypercube{cube.x + xOffset, cube.y + yOffset, cube.z + zOffset, cube.w + wOffset})
					}
				}
			}
		}
	}
	return cubes
}

func cycle[T interface {
	Cube | Hypercube
	getNeighbors() []T
}](cubes map[T]struct{}) map[T]struct{} {
	newActiveCubes := map[T]struct{}{}
	inactiveCubes := map[T]int{}

	for cube := range cubes {
		activeNeighbors := 0
		for _, neighbor := range cube.getNeighbors() {
			if _, found := cubes[neighbor]; found {
				activeNeighbors++
			} else {
				inactiveCubes[neighbor]++
			}
		}
		if activeNeighbors == 2 || activeNeighbors == 3 {
			newActiveCubes[cube] = struct{}{}
		}
	}

	for cube, surroundingActiveCubes := range inactiveCubes {
		if surroundingActiveCubes == 3 {
			newActiveCubes[cube] = struct{}{}
		}
	}

	return newActiveCubes
}

func cycleNTimes[T interface {
	Cube | Hypercube
	getNeighbors() []T
}](cubes map[T]struct{}, n int) map[T]struct{} {
	newActiveCubes := cubes
	for i := 0; i < n; i++ {
		newActiveCubes = cycle(newActiveCubes)
	}
	return newActiveCubes
}

func buildHypercubesFromCubes(cubes map[Cube]struct{}) map[Hypercube]struct{} {
	newCubes := map[Hypercube]struct{}{}
	for cube := range cubes {
		newCube := Hypercube{cube.x, cube.y, cube.z, 0}
		newCubes[newCube] = struct{}{}
	}
	return newCubes
}

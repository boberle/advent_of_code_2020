package main

import (
	"aoc2020/day01"
	"aoc2020/day02"
	"aoc2020/day03"
	"aoc2020/day04"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AOC 2020")

	flag.Parse()

	day, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Missing argument: day\n")
		return
	}

	file := flag.Arg(1)
	if len(strings.TrimSpace(file)) == 0 {
		panic("Missing argument: input file")
	}

	fmt.Printf("Running day %d, with file '%s'\n", day, file)

	switch day {
	case 1:
		day01.Run(file)
	case 2:
		day02.Run(file)
	case 3:
		day03.Run(file)
	case 4:
		day04.Run(file)
	default:
		fmt.Printf("Alas, I can't run for day %d.\n", day)
	}
}

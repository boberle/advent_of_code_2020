package main

import (
	"aoc2020/day01"
	"aoc2020/day02"
	"aoc2020/day03"
	"aoc2020/day04"
	"aoc2020/day05"
	"aoc2020/day06"
	"aoc2020/day07"
	"aoc2020/day08"
	"aoc2020/day09"
	"aoc2020/day10"
	"aoc2020/day11"
	"aoc2020/day12"
	"aoc2020/day13"
	"aoc2020/day14"
	"aoc2020/day15"
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
	case 5:
		day05.Run(file)
	case 6:
		day06.Run(file)
	case 7:
		day07.Run(file)
	case 8:
		day08.Run(file)
	case 9:
		day09.Run(file)
	case 10:
		day10.Run(file)
	case 11:
		day11.Run(file)
	case 12:
		day12.Run(file)
	case 13:
		day13.Run(file)
	case 14:
		day14.Run(file)
	case 15:
		day15.Run(file)
	default:
		fmt.Printf("Alas, I can't run for day %d.\n", day)
	}
}

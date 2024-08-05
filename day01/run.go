package day01

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	integers := parseFile(fh)

	twoIntegers := findTwoIntegersThatAddUpTo2020(integers)

	if len(twoIntegers) != 2 {
		log.Fatalln("couldn't find the 2 integers")
	}
	fmt.Printf(
		"PART 1: The integers are: %d and %d (= %d). Their multiplication is %d.\n",
		twoIntegers[0],
		twoIntegers[1],
		twoIntegers[0]+twoIntegers[1],
		twoIntegers[0]*twoIntegers[1],
	)

	threeIntegers := findThreeIntegersThatAddUpTo2020(integers)

	if len(threeIntegers) != 3 {
		log.Fatalln("couldn't find the 3 integers")
	}
	fmt.Printf(
		"PART 2: The integers are: %d, %d and %d (= %d). Their multiplication is %d.\n",
		threeIntegers[0],
		threeIntegers[1],
		threeIntegers[2],
		threeIntegers[0]+threeIntegers[1]+threeIntegers[2],
		threeIntegers[0]*threeIntegers[1]*threeIntegers[2],
	)
}

func parseFile(reader io.Reader) []int {
	integers := []int{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		integers = append(integers, n)
	}

	return integers
}

func findTwoIntegersThatAddUpTo2020(integers []int) []int {
	rv := []int{}
	for i, n1 := range integers {
		for j, n2 := range integers {
			if i == j {
				continue
			}
			if n1+n2 == 2020 {
				rv = append(rv, n1, n2)
				return rv
			}
		}
	}
	return rv
}

func findThreeIntegersThatAddUpTo2020(integers []int) []int {
	rv := []int{}
	for i, n1 := range integers {
		for j, n2 := range integers {
			for k, n3 := range integers {
				if i == j || i == k || j == k {
					continue
				}
				if n1+n2+n3 == 2020 {
					rv = append(rv, n1, n2, n3)
					return rv
				}
			}
		}
	}
	return rv
}

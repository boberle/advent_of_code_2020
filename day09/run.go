package day09

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	numbers := parseFile(fh)

	var preambleSize = 25
	if strings.Contains(infile, "input_test") {
		fmt.Println("TEST MODE")
		preambleSize = 5
	}

	firstNumberWhichIsNotTheSumOfThePreviousOnes, err := findFirstNumberWhichIsNotTheSumOfThePreviousOnes(&numbers, preambleSize)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Part 1: first number which is not the sum of the previous ones: %d\n", firstNumberWhichIsNotTheSumOfThePreviousOnes)

	set, err := findContiguousNumbersThatEqualTo(&numbers, firstNumberWhichIsNotTheSumOfThePreviousOnes)
	if err != nil {
		log.Fatalln(err)
	}
	min, max := findMinMax(&set)
	fmt.Printf("Print 2: sum of min and max of the contiguous set of numbers summing to %d: %d + %d = **%d**\n", firstNumberWhichIsNotTheSumOfThePreviousOnes, min, max, min+max)

}

func parseFile(reader io.Reader) []int {
	numbers := []int{}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		numbers = append(numbers, n)
	}
	return numbers
}

func findFirstNumberWhichIsNotTheSumOfThePreviousOnes(numbers *[]int, preambleSize int) (int, error) {
	length := len(*numbers)
	for i := preambleSize; i < length; i++ {
		found := false
		for j := i - preambleSize; j < i; j++ {
			for k := i - preambleSize; k < i; k++ {
				if j != k && (*numbers)[j]+(*numbers)[k] == (*numbers)[i] {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			return (*numbers)[i], nil
		}
	}
	return 0, errors.New("unable to find the number")
}

func findContiguousNumbersThatEqualTo(numbers *[]int, n int) ([]int, error) {
	length := len(*numbers)
	for i := 0; i < length; i++ {
		var total int
		for j := i; j < length; j++ {
			total += (*numbers)[j]
			if j-i > 0 && total == n {
				return (*numbers)[i : j+1], nil
			}
		}
	}
	return []int{}, errors.New("unable to find a range")
}

func findMinMax(numbers *[]int) (int, int) {
	if len(*numbers) == 0 {
		panic("empty slice")
	}
	var min, max int
	for i, n := range *numbers {
		if i == 0 || n < min {
			min = n
		}
		if i == 0 || n > max {
			max = n
		}
	}
	return min, max
}

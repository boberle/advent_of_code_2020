package day10

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	numbers := parseFile(fh)

	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)

	difference1Count, difference3Count := countDifferences(&numbers)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Part 1: number of difference of 1: %d; of 3: %d; multiplied: %d\n", difference1Count, difference3Count, difference1Count*difference3Count)

	ignoredNumbers := getIgnoredNumbers(&numbers)
	n := numbers[len(numbers)-1]
	cache := map[int]int{}
	pathCount := countPathsAt(ignoredNumbers, n, &cache)
	fmt.Printf("Part 2: number of possible paths: %d\n", pathCount)

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

func countDifferences(numbers *[]int) (int, int) {
	var count1, count3 int
	length := len(*numbers)

	for i := 1; i < length; i++ {
		diff := (*numbers)[i] - (*numbers)[i-1]
		if diff == 1 {
			count1++
		} else if diff == 3 {
			count3++
		} else {
			log.Fatalf("invalid difference: %d - %d = %d", (*numbers)[i], (*numbers)[i-1], diff)
		}
	}
	return count1, count3
}

func getIgnoredNumbers(numbers *[]int) *map[int]struct{} {
	ignored := &map[int]struct{}{}
	last := 0
	for i, n := range *numbers {
		if i > 0 {
			for j := last + 1; j < n; j++ {
				(*ignored)[j] = struct{}{}
			}
		}
		last = n
	}
	return ignored
}

func countPathsAt(ignoredNumbers *map[int]struct{}, n int, cache *map[int]int) int {
	if _, found := (*ignoredNumbers)[n]; found {
		return 0
	}

	if res, found := (*cache)[n]; found {
		return res
	}

	var res int
	if n == 0 {
		res = 1
	} else if n == 1 {
		res = 1
	} else if n == 2 {
		res = 2
	} else {
		res = countPathsAt(ignoredNumbers, n-1, cache) + countPathsAt(ignoredNumbers, n-2, cache) + countPathsAt(ignoredNumbers, n-3, cache)
	}
	(*cache)[n] = res
	return res
}

package day18

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	totalsWithoutPrecedence, totalsWithPrecedence := parseFile(fh)
	fmt.Printf("Sum of results of each line (part 1): %d\n", sum(totalsWithoutPrecedence))
	fmt.Printf("Sum of results of each line (part 2): %d\n", sum(totalsWithPrecedence))

}

func parseFile(reader io.Reader) ([]int, []int) {
	rvWithoutPrecedence := []int{}
	rvWithPrecedence := []int{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.Replace(line, " ", "", -1)
		val, _ := parseLine(line, 0, false)
		rvWithoutPrecedence = append(rvWithoutPrecedence, val)
		val, _ = parseLine(line, 0, true)
		rvWithPrecedence = append(rvWithPrecedence, val)
	}
	return rvWithoutPrecedence, rvWithPrecedence
}

func parseLine(line string, index int, withPrecedence bool) (int, int) {
	var rv, val int
	var found bool
	var op rune

	for index < len(line) {
		if line[index] == ')' {
			return rv, index + 1
		} else if withPrecedence && line[index] == '*' {
			val, index = parseLine(line, index+1, withPrecedence)
			rv *= val
			return rv, index
		} else if line[index] == '(' {
			val, index = parseLine(line, index+1, withPrecedence)
			if op == 0 {
				rv = val
			} else {
				if op == '+' {
					rv += val
				} else {
					rv *= val
				}
			}
		} else if line[index] == '+' || line[index] == '*' {
			op = rune(line[index])
			index++
		} else if val, index, found = parseNum(line, index); found {
			if op == 0 {
				rv = val
			} else {
				if op == '+' {
					rv += val
				} else {
					rv *= val
				}
			}
		}
	}
	return rv, index
}

func parseNum(line string, index int) (int, int, bool) {
	rvString := ""
	for index < len(line) && strings.Contains("0123456789", string(line[index])) {
		rvString += string(line[index])
		index++
	}
	if rvString == "" {
		return 0, index, false
	}
	rv, err := strconv.Atoi(rvString)
	if err != nil {
		panic(err)
	}
	return rv, index, true
}

func sum(numbers []int) int {
	rv := 0
	for _, number := range numbers {
		rv += number
	}
	return rv
}

package day05

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
)

type seatNumber string
type seatID int
type half int

const (
	First  half = 0
	Second      = 1
)

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()
	numbers := parseFile(fh)

	seatIDs := getSeatIDs(numbers)
	fmt.Printf("Part 1: Highest seat id is %d among %d numbers\n", max(seatIDs), len(numbers))

	freeSeatID := findFreeSeatID(seatIDs)
	fmt.Printf("Part 2: The free seat is: %d\n", freeSeatID)
}

func getSeatIDs(numbers []seatNumber) []seatID {
	ids := make([]seatID, len(numbers))
	for _, number := range numbers {
		ids = append(ids, number.getID())
	}
	return ids
}

func max(items []seatID) seatID {
	var max seatID
	for i, item := range items {
		if i == 0 || max < item {
			max = item
		}
	}
	return max
}

func parseFile(reader io.Reader) []seatNumber {
	scanner := bufio.NewScanner(reader)
	numbers := []seatNumber{}
	for scanner.Scan() {
		numbers = append(numbers, seatNumber(scanner.Text()))
	}
	return numbers
}

func getBinaryPosition(seq []half) int {
	position := 0
	for i, v := range seq {
		if v == Second {
			position += int(math.Pow(2, float64(i)))
		}
	}
	return position
}

func convertSeatNumberToHalves(number string, mapping map[rune]half) []half {
	runes := []rune(number)
	rv := []half{}
	for i := len(runes) - 1; i >= 0; i-- {
		half, found := mapping[runes[i]]
		if !found {
			log.Fatalf("invalid code %q in number: %s", runes[i], number)
		}
		rv = append(rv, half)
	}
	return rv
}

func (number seatNumber) getID() seatID {
	row := getBinaryPosition(convertSeatNumberToHalves(string(number)[:7], map[rune]half{'F': First, 'B': Second}))
	col := getBinaryPosition(convertSeatNumberToHalves(string(number)[7:], map[rune]half{'L': First, 'R': Second}))
	return seatID(row*8 + col)
}

func findFreeSeatID(ids []seatID) seatID {
	// the seat is the only one that has an ID surrounded with ID-1 and ID+1
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})
	candidates := []seatID{}
	for i, id := range ids {
		if i < len(ids)-1 {
			if ids[i+1] == id+2 {
				candidates = append(candidates, id+1)
			}
		}
	}
	if len(candidates) != 1 {
		log.Fatalln("unable to find the seat, candidates are", candidates)
	}
	return candidates[0]
}

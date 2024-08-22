package day15

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	numbers []int
}

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	game := parseFile(fh)
	nPart1 := game.play(2020)
	fmt.Printf("The 2020th number spoken is %d (part 1)\n", nPart1)
	nPart2 := game.play(30000000)
	fmt.Printf("The 30000000th number spoken is %d (part 2)\n", nPart2)

}

func parseFile(reader io.Reader) Game {
	numbers := make([]int, 0)

	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		for _, nString := range strings.Split(scanner.Text(), ",") {
			n, err := strconv.Atoi(nString)
			if err != nil {
				log.Fatalf("string %s is not a number", nString)
			}
			numbers = append(numbers, n)
		}
	}
	return Game{numbers: numbers}
}

func (game *Game) play(turns int) int {
	numbers := make(map[int]int, 0)
	for i, n := range game.numbers[:len(game.numbers)-1] {
		numbers[n] = i + 1
	}
	current := game.numbers[len(game.numbers)-1]
	for i := len(game.numbers); i < turns; i++ {
		if index, found := numbers[current]; found {
			numbers[current] = i
			current = i - index
		} else {
			numbers[current] = i
			current = 0
		}
	}
	return current
}

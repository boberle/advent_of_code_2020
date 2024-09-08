package day23

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Cup struct {
	label int
	next  *Cup
}

type Labels []*Cup

type Game struct {
	currentCup *Cup
	labels     Labels
}

func Run(infile string) {

	content, err := os.ReadFile(infile)
	if err != nil {
		panic(err)
	}

	game := parseFile(strings.NewReader(string(content)))
	game.play(100)
	fmt.Printf("The order of the cups after cup 1 and 100 moves is (part 1): %s\n", game.getLabelsAfter1())

	game = parseFile(strings.NewReader(string(content)))
	game.add1MCups()
	game.play(10_000_000)
	fmt.Printf("The multiplied labels of the 2 cups after cup 1 are (part 2): %d\n", game.getMultipliedLabelsAfter1())

}

func parseFile(reader io.Reader) Game {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	labels := make(Labels, len(scanner.Text())+1)
	var first, last *Cup
	for _, char := range scanner.Text() {
		label, err := strconv.Atoi(string(char))
		if err != nil {
			panic(err)
		}
		cup := Cup{label: label}
		if first == nil {
			first = &cup
		}
		if last != nil {
			last.next = &cup
		}
		last = &cup
		labels[label] = &cup
	}
	last.next = first

	return Game{
		currentCup: first,
		labels:     labels,
	}
}

func (g *Game) extractCups() [3]*Cup {
	cups := [3]*Cup{
		g.currentCup.next,
		g.currentCup.next.next,
		g.currentCup.next.next.next,
	}
	g.currentCup.next = g.currentCup.next.next.next.next
	return cups
}

func (g *Game) insertCups(cups [3]*Cup, destinationCup *Cup) {
	cups[2].next = destinationCup.next
	destinationCup.next = cups[0]
}

func (g *Game) getDestinationCup(extractedCups [3]*Cup) *Cup {
	destLabel := g.currentCup.label
	for {
		destLabel--
		if destLabel < 1 {
			destLabel = len(g.labels) - 1
		}
		if destLabel != extractedCups[0].label && destLabel != extractedCups[1].label && destLabel != extractedCups[2].label {
			return g.labels[destLabel]
		}
	}
}

func (g *Game) playOneMove() {
	cups := g.extractCups()
	destCup := g.getDestinationCup(cups)
	g.insertCups(cups, destCup)
	g.currentCup = g.currentCup.next
}

func (g *Game) play(n int) {
	for i := 0; i < n; i++ {
		g.playOneMove()
	}
}

func (g *Game) getLabelsAfter1() string {
	labels := []string{}
	cup := g.labels[1].next
	length := len(g.labels) - 2
	for i := 0; i < length; i++ {
		labels = append(labels, strconv.Itoa(cup.label))
		cup = cup.next
	}
	return strings.Join(labels, "")
}

func (g *Game) findLastCup() *Cup {
	cur := g.currentCup
	for {
		if cur.next == g.currentCup {
			return cur
		}
		cur = cur.next
	}
}

func (g *Game) add1MCups() {
	cur := g.findLastCup()
	for i := len(g.labels); i <= 1_000_000; i++ {
		cup := &Cup{label: i}
		g.labels = append(g.labels, cup)
		cur.next = cup
		cur = cup
	}
	cur.next = g.currentCup
	if len(g.labels) != 1_000_001 {
		panic("not 1M")
	}
}

func (g *Game) getMultipliedLabelsAfter1() int {
	cup1 := g.labels[1]
	return cup1.next.label * cup1.next.next.label
}

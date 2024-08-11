package day07

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Color string

type Bag struct {
	color  Color
	bags   []*Bag
	counts []int
}

type Bags map[Color]*Bag
type Parents map[Color][]*Bag

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	bags, parents := parseFile(fh)
	numberOfShinyGoldAncestors := countShinyGoldAncestors(bags["shiny gold"], &parents)
	fmt.Printf("Part 1: number of shiny gold ancestors: %d\n", numberOfShinyGoldAncestors)

	numberOfBagsInShinyGold := countBags(bags["shiny gold"])
	fmt.Printf("Part 2: number of bags in the shiny gold bag: %d\n", numberOfBagsInShinyGold)

}

func countShinyGoldAncestors(shinyGoldBag *Bag, parents *Parents) int {
	ancestors := map[*Bag]struct{}{}
	getAncestors(shinyGoldBag, parents, &ancestors)
	return len(ancestors)
}

func countBags(bag *Bag) int {
	total := 0
	for i := 0; i < len(bag.bags); i++ {
		total += bag.counts[i]
		//fmt.Println("count", bag.bags[i].color, bag.counts[i], total)
		total += countBags(bag.bags[i]) * bag.counts[i]
	}
	return total
}

func parseFile(reader io.Reader) (Bags, Parents) {
	scanner := bufio.NewScanner(reader)

	rContainer, err := regexp.Compile("([a-z ]+) bags contain ([0-9a-z, ]+)\\.")
	if err != nil {
		log.Fatalln(err)
	}

	rContent, err := regexp.Compile("(\\d) ([a-z ]+) bags?")
	if err != nil {
		log.Fatalln(err)
	}

	bags := Bags{}
	parents := Parents{}

	for scanner.Scan() {
		line := scanner.Text()

		var bag *Bag
		bagMatch := rContainer.FindStringSubmatch(line)
		if len(bagMatch) != 3 {
			log.Fatalf("%s: couldn't match", line)
		}
		bagColor := Color(bagMatch[1])
		bag = bags[bagColor]
		if bag == nil {
			bag = &Bag{color: bagColor}
			bags[bagColor] = bag
		}

		if len(bag.bags) > 0 {
			log.Fatalf("bag %s has already a content", bag.color)
		}

		contentString := bagMatch[2]
		if contentString != "no other bags" {
			contentMatches := rContent.FindAllStringSubmatch(contentString, -1)
			if len(contentMatches) == 0 {
				log.Fatalf("%s: couldn't match", contentString)
			}
			for _, contentMatch := range contentMatches {
				quantity, _ := strconv.Atoi(contentMatch[1])
				bag.counts = append(bag.counts, quantity)
				contentBagColor := Color(contentMatch[2])
				var contentBag *Bag
				contentBag = bags[contentBagColor]
				if contentBag == nil {
					contentBag = &Bag{color: contentBagColor}
					bags[contentBagColor] = contentBag
				}
				bag.bags = append(bag.bags, contentBag)
				parents[contentBag.color] = append(parents[contentBag.color], bag)
			}
		}
	}
	return bags, parents
}

func getAncestors(bag *Bag, parents *Parents, ancestors *map[*Bag]struct{}) {
	parentList := (*parents)[bag.color]
	if parentList != nil {
		for _, p := range parentList {
			(*ancestors)[p] = struct{}{}
			getAncestors(p, parents, ancestors)
		}
	}
}

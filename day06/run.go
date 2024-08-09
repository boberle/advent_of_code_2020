package day06

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type question rune

type person struct {
	yesAnsweredQuestion []question
}

type group struct {
	people []person
}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	groups := parseFile(fh)
	numberOfQuestionsAnyoneYesAnsweredInAllGroups := countQuestionsAnyoneYesAnsweredInAllGroups(groups)
	fmt.Printf("Part 1: number of questions anyone yes-answered in all group: %d\n", numberOfQuestionsAnyoneYesAnsweredInAllGroups)

	numberOfQuestionsEveryoneYesAnsweredInAllGroups := countQuestionsEveryoneYesAnsweredInAllGroups(groups)
	fmt.Printf("Part 2: number of questions everyone yes-answered in all group: %d\n", numberOfQuestionsEveryoneYesAnsweredInAllGroups)
}

func parseFile(reader io.Reader) []group {
	scanner := bufio.NewScanner(reader)
	groups := []group{}
	currentGroup := group{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			groups = append(groups, currentGroup)
			currentGroup = group{}
			continue
		}
		questions := []question{}
		for _, questionCode := range line {
			questions = append(questions, question(questionCode))
		}
		person := person{yesAnsweredQuestion: questions}
		currentGroup.people = append(currentGroup.people, person)
	}
	groups = append(groups, currentGroup)
	return groups
}

func countQuestionsAnyoneYesAnsweredInAllGroups(groups []group) int {
	total := 0
	for _, group := range groups {
		total += group.numberOfQuestionsAnyoneYesAnswered()
	}
	return total
}

func (g *group) numberOfQuestionsAnyoneYesAnswered() int {
	questions := map[question]int{}
	for _, person := range g.people {
		for _, question := range person.yesAnsweredQuestion {
			questions[question] = 1
		}
	}
	return len(questions)
}

func countQuestionsEveryoneYesAnsweredInAllGroups(groups []group) int {
	total := 0
	for _, group := range groups {
		total += group.numberOfQuestionsEveryoneYesAnswered()
	}
	return total
}
func (g *group) numberOfQuestionsEveryoneYesAnswered() int {
	questions := map[question]int{}
	for _, person := range g.people {
		for _, question := range person.yesAnsweredQuestion {
			questions[question]++
		}
	}
	rv := 0
	for _, v := range questions {
		if v == len(g.people) {
			rv++
		}
	}
	return rv
}

package day19

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type RuleChoices struct {
	sequences []RuleSequence
}

type RuleSequence struct {
	rules []*Rule
}

type Terminal struct {
	value rune
}

type Rule interface {
	doesMatch(Message, int) (bool, int)
}

type Rule8 struct {
	rule11 *Rule11
	rule42 *Rule
}

type Rule11 struct {
	rule31 *Rule
	rule42 *Rule
}

type Message string

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	messages, rootRule := parseFile(fh)
	validMessageCount := countValidMessages(&messages, rootRule)
	fmt.Printf("Number of valid messages (part 1): %d\n", validMessageCount)

	r0 := rootRule
	r11 := r0.(RuleSequence).rules[1]
	r42 := (*r11).(RuleSequence).rules[0]
	r31 := (*r11).(RuleSequence).rules[1]

	rule11 := Rule11{
		rule42: r42,
		rule31: r31,
	}
	rule8 := Rule8{
		rule11: &rule11,
		rule42: r42,
	}
	validMessageCount = countValidMessages(&messages, rule8)
	fmt.Printf("Number of valid messages (part 2): %d\n", validMessageCount)

}

func parseFile(reader io.Reader) ([]Message, Rule) {
	scanner := bufio.NewScanner(reader)

	rawRules := map[int]string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		parts := strings.Split(line, ":")
		ruleIndex, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		ruleContent := strings.TrimSpace(parts[1])
		if _, found := rawRules[ruleIndex]; found {
			panic(fmt.Sprintf("rule with index %d already there", ruleIndex))
		}
		rawRules[ruleIndex] = ruleContent
	}
	rootRule := buildRules(&rawRules)

	messages := make([]Message, 0)
	for scanner.Scan() {
		message := Message(strings.TrimSpace(scanner.Text()))
		messages = append(messages, message)
	}

	fmt.Printf("Number of rules: %d\n", len(rawRules))
	fmt.Printf("Number of messages: %d\n", len(messages))

	return messages, *rootRule
}

func buildRules(raw *map[int]string) *Rule {
	rules := map[int]*Rule{}
	return buildRule(raw, &rules, 0)
}

func buildRule(raw *map[int]string, rules *map[int]*Rule, index int) *Rule {
	var rule Rule
	content := (*raw)[index]
	if strings.HasPrefix(content, "\"") && strings.HasSuffix(content, "\"") {
		rule = Terminal{rune(content[1])}
	} else if strings.Contains(content, "|") {
		choices := RuleChoices{sequences: []RuleSequence{}}
		for _, sequenceString := range strings.Split(content, "|") {
			sequenceString = strings.TrimSpace(sequenceString)
			choices.sequences = append(choices.sequences, buildSequence(raw, rules, sequenceString))
		}
		rule = choices
	} else {
		rule = buildSequence(raw, rules, content)
	}
	(*rules)[index] = &rule
	return &rule
}

func buildSequence(raw *map[int]string, rules *map[int]*Rule, content string) RuleSequence {
	indices := strings.Split(content, " ")
	seq := RuleSequence{}
	for _, iString := range indices {
		if i, err := strconv.Atoi(iString); err != nil {
			panic(err)
		} else {
			r, found := (*rules)[i]
			if !found {
				r = buildRule(raw, rules, i)
			}
			seq.rules = append(seq.rules, r)
		}
	}
	return seq
}

func (t Terminal) doesMatch(m Message, i int) (bool, int) {
	if i >= len(m) {
		return false, 0
	}
	var next int
	if i == len(m)-1 {
		next = 0
	} else {
		next = i + 1
	}
	return rune(m[i]) == t.value, next
}

func (seq RuleSequence) doesMatch(m Message, i int) (bool, int) {
	for _, rule := range seq.rules {
		match, next := (*rule).doesMatch(m, i)
		if !match {
			return false, 0
		}
		i = next
	}
	return true, i
}

func (choices RuleChoices) doesMatch(m Message, i int) (bool, int) {
	for _, seq := range choices.sequences {
		if match, next := seq.doesMatch(m, i); match {
			return true, next
		}
	}
	return false, 0
}

func doesMatch(rule Rule, m Message) bool {
	match, next := rule.doesMatch(m, 0)
	if !match {
		return false
	}
	if next == 0 {
		return true
	}
	return false
}

func (rule Rule8) doesMatch(m Message, i int) (bool, int) {
	seq := RuleSequence{rules: []*Rule{rule.rule42}}
	for j := 0; j < len(m)-i; j++ {
		if match, n := seq.doesMatch(m, i); match {
			if match, n = rule.rule11.doesMatch(m, n); match {
				return true, n
			}
		}
		seq.rules = append(seq.rules, rule.rule42)
	}
	return false, i
}

func (rule Rule11) doesMatch(m Message, i int) (bool, int) {
	seq := RuleSequence{rules: []*Rule{rule.rule42, rule.rule31}}
	for j := 0; j < len(m)-i; j++ {
		if match, n := seq.doesMatch(m, i); match {
			return true, n
		}
		seq.rules = append([]*Rule{rule.rule42}, seq.rules...)
		seq.rules = append(seq.rules, rule.rule31)
	}
	return false, i
}

func countValidMessages(messages *[]Message, rule Rule) int {
	total := 0
	for _, message := range *messages {
		if doesMatch(rule, message) {
			total++
		}
	}
	return total
}

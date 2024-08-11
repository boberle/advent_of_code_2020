package day07

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getTestData() (Bags, Parents) {
	gold := Bag{color: Color("gold")}
	blue := Bag{color: Color("blue"), counts: []int{1}, bags: []*Bag{&gold}}
	green := Bag{color: Color("green"), counts: []int{1}, bags: []*Bag{&blue}}
	red := Bag{color: Color("red"), counts: []int{2, 3}, bags: []*Bag{&blue, &green}}
	purple := Bag{color: Color("purple")}
	yellow := Bag{color: Color("yellow"), counts: []int{2, 2}, bags: []*Bag{&green, &purple}}
	bags := Bags{
		Color("gold"):   &gold,
		Color("blue"):   &blue,
		Color("green"):  &green,
		Color("red"):    &red,
		Color("purple"): &purple,
		Color("yellow"): &yellow,
	}
	parents := Parents{
		Color("gold"):   []*Bag{&blue},
		Color("blue"):   []*Bag{&red, &green},
		Color("green"):  []*Bag{&red, &yellow},
		Color("purple"): []*Bag{&yellow},
	}
	return bags, parents
}

func Test_parseFile(t *testing.T) {
	map1, parents1 := getTestData()

	fadedBlue := Bag{color: Color("faded blue")}
	dottedBlack := Bag{color: Color("dotted black")}
	lightRed := Bag{color: Color("light red"), counts: []int{1, 2}, bags: []*Bag{&fadedBlue, &dottedBlack}}
	map2 := Bags{
		Color("faded blue"):   &fadedBlue,
		Color("dotted black"): &dottedBlack,
		Color("light red"):    &lightRed,
	}
	parents2 := Parents{
		Color("faded blue"):   []*Bag{&lightRed},
		Color("dotted black"): []*Bag{&lightRed},
	}

	tests := []struct {
		content         string
		expected        Bags
		expectedParents Parents
	}{
		{
			"red bags contain 2 blue bags, 3 green bags.\nblue bags contain 1 gold bag.\ngold bags contain no other bags.\ngreen bags contain 1 blue bag.\nyellow bags contain 2 green bags, 2 purple bags.\npurple bags contain no other bags.",
			map1,
			parents1,
		},
		{
			"light red bags contain 1 faded blue bag, 2 dotted black bags.\nfaded blue bags contain no other bags.\ndotted black bags contain no other bags.",
			map2,
			parents2,
		},
	}
	for _, test := range tests {
		reader := strings.NewReader(test.content)
		actual, actualParents := parseFile(reader)
		assert.Equal(t, test.expected, actual, test.content)
		assert.Equal(t, test.expectedParents, actualParents, test.content)
	}
}

func Test_getAncestors(t *testing.T) {
	bags, parents := getTestData()

	tests := []struct {
		color    Color
		expected []Color
	}{
		{"green", []Color{"red", "yellow"}},
		{"red", []Color{}},
		{"blue", []Color{"red", "green", "yellow"}},
		{"gold", []Color{"blue", "red", "green", "yellow"}},
	}

	for _, test := range tests {
		ancestors := map[*Bag]struct{}{}
		getAncestors(bags[test.color], &parents, &ancestors)
		actual := map[Color]int{}
		for ancestor := range ancestors {
			actual[ancestor.color] = 1
		}
		expected := map[Color]int{}
		for _, color := range test.expected {
			expected[color] = 1
		}
		assert.Equal(t, expected, actual, test.color)
	}
}

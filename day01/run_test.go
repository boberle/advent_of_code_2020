package day01

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	reader := strings.NewReader("1\n2\n3")

	actual := parseFile(reader)
	assert.Equal(t, []int{1, 2, 3}, actual)
}

func Test_findTwoIntegersThatAddUpTo2020(t *testing.T) {
	cases := []struct {
		name     string
		data     []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"2 integers in input", []int{10, 2010}, []int{10, 2010}},
		{"at the end", []int{5, 10, 2010}, []int{10, 2010}},
		{"at the beginning", []int{10, 2010, 5}, []int{10, 2010}},
		{"around", []int{10, 5, 2010}, []int{10, 2010}},
		{"no number", []int{10, 2020, 5}, []int{}},
	}

	for _, test := range cases {
		actual := findTwoIntegersThatAddUpTo2020(test.data)
		if !assert.Equal(t, test.expected, actual) {
			t.Errorf("test %s failed", test.name)
		}
	}
}

func Test_findThreeIntegersThatAddUpTo2020(t *testing.T) {
	cases := []struct {
		name     string
		data     []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"3 integers in input", []int{10, 2008, 2}, []int{10, 2008, 2}},
		{"at the end", []int{5, 10, 2008, 2}, []int{10, 2008, 2}},
		{"at the beginning", []int{10, 2008, 2, 5}, []int{10, 2008, 2}},
		{"around", []int{10, 5, 2008, 2}, []int{10, 2008, 2}},
		{"no number", []int{10, 2010, 2, 5}, []int{}},
	}

	for _, test := range cases {
		actual := findThreeIntegersThatAddUpTo2020(test.data)
		if !assert.Equal(t, test.expected, actual) {
			t.Errorf("test %s failed", test.name)
		}
	}
}

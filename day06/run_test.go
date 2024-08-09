package day06

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	reader := strings.NewReader("abc\n\na\nb\nc\n\nab\nac\n\na\na\na\na\n\nb")
	expected := []group{
		{
			[]person{
				{[]question{'a', 'b', 'c'}},
			},
		},
		{
			[]person{
				{[]question{'a'}},
				{[]question{'b'}},
				{[]question{'c'}},
			},
		},
		{
			[]person{
				{[]question{'a', 'b'}},
				{[]question{'a', 'c'}},
			},
		},
		{
			[]person{
				{[]question{'a'}},
				{[]question{'a'}},
				{[]question{'a'}},
				{[]question{'a'}},
			},
		},
		{
			[]person{
				{[]question{'b'}},
			},
		},
	}
	actual := parseFile(reader)
	assert.Equal(t, expected, actual)
}

func Test_group_numberOfQuestionsAnyoneYesAnswered(t *testing.T) {
	cases := []struct {
		group    group
		expected int
	}{
		{group: group{
			[]person{
				{[]question{'a', 'b', 'c'}},
			},
		}, expected: 3},
		{group: group{
			[]person{
				{[]question{'a'}},
				{[]question{'b'}},
				{[]question{'c'}},
			},
		}, expected: 3},
		{group: group{
			[]person{
				{[]question{'a', 'b'}},
				{[]question{'a', 'c'}},
			},
		}, expected: 3},
		{group: group{
			[]person{
				{[]question{'a'}},
				{[]question{'a'}},
			},
		}, expected: 1},
	}

	for _, test := range cases {
		actual := test.group.numberOfQuestionsAnyoneYesAnswered()
		assert.Equal(t, test.expected, actual, test.group)
	}
}

func Test_countQuestionsEveryoneYesAnsweredInAllGroups(t *testing.T) {
	groups := []group{
		{
			[]person{
				{[]question{'a', 'b', 'c'}},
			},
		},
		{
			[]person{
				{[]question{'a'}},
				{[]question{'b'}},
				{[]question{'d'}},
			},
		},
	}
	expected := 3
	actual := countQuestionsEveryoneYesAnsweredInAllGroups(groups)
	assert.Equal(t, expected, actual)
}

func Test_group_numberOfQuestionsEveryoneYesAnswered(t *testing.T) {
	cases := []struct {
		group    group
		expected int
	}{
		{group: group{
			[]person{
				{[]question{'a', 'b', 'c'}},
			},
		}, expected: 3},
		{group: group{
			[]person{
				{[]question{'a'}},
				{[]question{'b'}},
				{[]question{'c'}},
			},
		}, expected: 0},
		{group: group{
			[]person{
				{[]question{'a', 'b'}},
				{[]question{'a', 'c'}},
			},
		}, expected: 1},
		{group: group{
			[]person{
				{[]question{'a'}},
				{[]question{'a'}},
			},
		}, expected: 1},
	}

	for _, test := range cases {
		actual := test.group.numberOfQuestionsEveryoneYesAnswered()
		assert.Equal(t, test.expected, actual, test.group)
	}
}

func Test_countQuestionsAnyoneYesAnsweredInAllGroups(t *testing.T) {
	groups := []group{
		{
			[]person{
				{[]question{'a', 'b', 'c'}},
			},
		},
		{
			[]person{
				{[]question{'a'}},
				{[]question{'b'}},
				{[]question{'d'}},
			},
		},
	}
	expected := 6
	actual := countQuestionsAnyoneYesAnsweredInAllGroups(groups)
	assert.Equal(t, expected, actual)
}

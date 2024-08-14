package day08

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	input := "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"
	expected := []Op{
		{"nop", 0, false},
		{"acc", 1, false},
		{"jmp", 4, false},
		{"acc", 3, false},
		{"jmp", -3, false},
		{"acc", -99, false},
		{"acc", 1, false},
		{"jmp", -4, false},
		{"acc", 6, false},
	}

	actual := parseFile(strings.NewReader(input))
	assert.Equal(t, expected, actual)
}

func Test_executeProgram(t *testing.T) {
	tests := []struct {
		program  []Op
		accValue int
		finished bool
	}{
		{[]Op{
			{"nop", 0, false},
			{"acc", 1, false},
			{"jmp", 4, false},
			{"acc", 3, false},
			{"jmp", -3, false},
			{"acc", -99, false},
			{"acc", 1, false},
			{"jmp", -4, false},
			{"acc", 6, false},
		}, 5, false},
		{[]Op{
			{"nop", 0, false},
			{"acc", 1, false},
			{"jmp", 4, false},
			{"acc", 3, false},
			{"jmp", -3, false},
			{"acc", -99, false},
			{"acc", 1, false},
			{"nop", -4, false},
			{"acc", 6, false},
		}, 8, true},
	}

	for i, test := range tests {
		actualValue, actualFinished := executeProgram(&test.program)
		if actualValue != test.accValue {
			t.Errorf("test #%d: expected %d, got %d", i, test.accValue, actualValue)
		}
		if actualFinished != test.finished {
			t.Errorf("test #%d: expected %v, got %v", i, test.finished, actualFinished)
		}

	}

}

func Test_fixProgram(t *testing.T) {
	program := []Op{
		{"nop", 0, false},
		{"acc", 1, false},
		{"jmp", 4, false},
		{"acc", 3, false},
		{"jmp", -3, false},
		{"acc", -99, false},
		{"acc", 1, false},
		{"jmp", -4, false},
		{"acc", 6, false},
	}
	actualValue := fixProgram(&program)
	if actualValue != 8 {
		t.Errorf("expected %d, got %d", 8, actualValue)
	}
}

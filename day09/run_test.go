package day09

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	input := "1\n2\n3\n10\n11"
	expected := []int{1, 2, 3, 10, 11}

	actual := parseFile(strings.NewReader(input))
	assert.Equal(t, expected, actual, input)
}

func Test_findFirstNumberWhichIsNotTheSumOfThePreviousOnes(t *testing.T) {
	tests := []struct {
		name         string
		numbers      *[]int
		preambleSize int
		expected     int
		expectedErr  bool
	}{
		{"only one number to be found", &[]int{1, 2, 3, 50}, 3, 50, false},
		{"choose the first number", &[]int{1, 2, 3, 50, 51}, 3, 50, false},
		{"choose the first number, after other ones", &[]int{1, 2, 3, 3, 6, 50}, 3, 50, false},
		{"number not found", &[]int{1, 2, 3, 3}, 3, 0, true},
	}
	for _, tt := range tests {
		actual, err := findFirstNumberWhichIsNotTheSumOfThePreviousOnes(tt.numbers, tt.preambleSize)
		if err == nil && tt.expectedErr {
			t.Error(tt.name, "expected an error, got no error")
		} else if err != nil && !tt.expectedErr {
			t.Error(tt.name, "expected no error, got an error:", err)
		} else if actual != tt.expected {
			t.Errorf("%s: expected %d, got %d", tt.name, tt.expected, actual)
		}
	}
}

func Test_findContiguousNumbersThatEqualTo(t *testing.T) {
	tests := []struct {
		name        string
		numbers     []int
		n           int
		expected    []int
		expectedErr bool
	}{
		{"range of 2 numbers", []int{1, 2, 3, 4}, 5, []int{2, 3}, false},
		{"range of 3 numbers", []int{1, 2, 3, 4}, 6, []int{1, 2, 3}, false},
		{"no range of 1 number", []int{1, 2, 3, 4}, 2, []int{}, true},
		{"no range", []int{1, 2, 3, 4}, 50, []int{}, true},
	}
	for _, tt := range tests {
		actual, err := findContiguousNumbersThatEqualTo(&tt.numbers, tt.n)
		if err == nil && tt.expectedErr {
			t.Errorf("%s: expected an error, got none", tt.name)
		} else if err != nil && !tt.expectedErr {
			t.Errorf("%s: expected no error, got one", tt.name)
		} else {
			assert.Equal(t, tt.expected, actual)
		}
	}
}

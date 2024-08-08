package day05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getBinaryPosition(t *testing.T) {
	cases := []struct {
		sequence []half
		expected int
	}{
		{[]half{First, First, Second, Second, First, Second, First}, 44},
		{[]half{Second, First, Second}, 5},
		{[]half{First, Second, Second, First, First, First, Second}, 70},
		{[]half{First, Second, Second, Second, First, First, First}, 14},
		{[]half{First, Second, Second, First, First, Second, Second}, 102},
		{[]half{Second, Second, Second}, 7},
		{[]half{First, First, Second}, 4},
	}

	for _, test := range cases {
		actual := getBinaryPosition(test.sequence)
		assert.Equal(t, test.expected, actual, test.sequence)
	}
}

func Test_convertSeatNumberToHalves(t *testing.T) {
	cases := []struct {
		number   string
		mapping  map[rune]half
		expected []half
	}{
		{"BFBF", map[rune]half{'F': First, 'B': Second}, []half{First, Second, First, Second}},
		{"LR", map[rune]half{'L': First, 'R': Second}, []half{Second, First}},
	}

	for _, test := range cases {
		actual := convertSeatNumberToHalves(test.number, test.mapping)
		assert.Equal(t, test.expected, actual, test.number)
	}
}

func Test_seatNumber_getID(t *testing.T) {
	cases := []struct {
		number     seatNumber
		expectedId seatID
	}{
		{"FBFBBFFRLR", 357},
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, test := range cases {
		actual := test.number.getID()
		assert.Equal(t, test.expectedId, actual, test.number)
	}
}

func Test_findFreeSeatID(t *testing.T) {
	cases := []struct {
		ids      []seatID
		expected seatID
	}{
		{[]seatID{1, 2, 3, 5, 6}, 4},
		{[]seatID{3, 5, 1, 4, 6}, 2},
	}

	for _, test := range cases {
		actual := findFreeSeatID(test.ids)
		assert.Equal(t, test.expected, actual, test.ids)
	}
}

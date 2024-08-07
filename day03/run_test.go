package day03

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	reader := strings.NewReader(".#.#\n....\n####")
	actualBoard := parseFile(reader)

	expectedBoard := GameBoard{
		trees: map[Point]struct{}{
			{1, 0}: {},
			{3, 0}: {},
			{0, 2}: {},
			{1, 2}: {},
			{2, 2}: {},
			{3, 2}: {},
		},
		width:  4,
		height: 3,
	}
	assert.Equal(t, expectedBoard, actualBoard)
}

func Test_traverseBoardAndCountTrees(t *testing.T) {
	board := GameBoard{
		trees: map[Point]struct{}{
			{2, 1}: {},
			{0, 2}: {},
		},
		width:  4,
		height: 3,
	}
	actual := board.traverseBoardAndCountTrees(2, 1)
	assert.Equal(t, 2, actual)
}

func Test_checkAllSlopes(t *testing.T) {
	board := GameBoard{
		map[Point]struct{}{
			{2, 1}: {},
			{3, 1}: {},
			{0, 2}: {},
			{2, 2}: {},
		},
		4,
		3,
	}
	actual := board.checkAllSlopes(slopeList{{2, 1}, {3, 1}})
	assert.Equal(t, 4, actual)
}

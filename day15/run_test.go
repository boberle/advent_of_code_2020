package day15

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGame_play(t *testing.T) {
	tests := []struct {
		name  string
		game  Game
		turns int
		want  int
	}{
		{"play", Game{[]int{1, 2}}, 3, 0},
		{"play", Game{[]int{1, 1}}, 3, 1},
		{"play", Game{[]int{1, 2, 1}}, 4, 2},
		{"play", Game{[]int{0, 3, 6}}, 4, 0},
		{"play", Game{[]int{0, 3, 6}}, 5, 3},
		{"play", Game{[]int{0, 3, 6}}, 6, 3},
		{"play", Game{[]int{0, 3, 6}}, 7, 1},
		{"play", Game{[]int{0, 3, 6}}, 8, 0},
		{"play", Game{[]int{0, 3, 6}}, 9, 4},
		{"play", Game{[]int{0, 3, 6}}, 10, 0},
		{"play", Game{[]int{1, 3, 2}}, 2020, 1},
		{"play", Game{[]int{2, 1, 3}}, 2020, 10},
		{"play", Game{[]int{1, 2, 3}}, 2020, 27},
		{"play", Game{[]int{2, 3, 1}}, 2020, 78},
		{"play", Game{[]int{3, 2, 1}}, 2020, 438},
		{"play", Game{[]int{3, 1, 2}}, 2020, 1836},
		{"play", Game{[]int{0, 3, 6}}, 30000000, 175594},
		{"play", Game{[]int{1, 3, 2}}, 30000000, 2578},
		{"play", Game{[]int{2, 1, 3}}, 30000000, 3544142},
		{"play", Game{[]int{1, 2, 3}}, 30000000, 261214},
		{"play", Game{[]int{2, 3, 1}}, 30000000, 6895259},
		{"play", Game{[]int{3, 2, 1}}, 30000000, 18},
		{"play", Game{[]int{3, 1, 2}}, 30000000, 362},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.game.play(tt.turns)
			assert.Equal(t, tt.want, got)
		})
	}
}

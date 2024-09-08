package day23

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		labels []int
	}{
		{"3", "123", []int{1, 2, 3}},
		{"3 reverse", "321", []int{3, 2, 1}},
		{"mixed order", "389125467", []int{3, 8, 9, 1, 2, 5, 4, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			gotGame := parseFile(reader)
			cur := gotGame.currentCup
			first := cur
			for _, lbl := range tt.labels {
				if cur.label != lbl {
					t.Errorf("%s: label mismatch in linked list, expected %d, got %d\n", tt.name, lbl, cur.label)
				}
				cur = cur.next
				if gotGame.labels[lbl].label != lbl {
					t.Errorf("%s: label mismatch in label array, expected %d, got %d\n", tt.name, lbl, gotGame.labels[lbl].label)
				}
			}
			if first != cur {
				t.Errorf("%s: last_cup.next is not the first cup", tt.name)
			}
		})
	}
}

func TestGame_playOneMove(t *testing.T) {
	cup7 := &Cup{7, nil}
	cup6 := &Cup{6, cup7}
	cup4 := &Cup{4, cup6}
	cup5 := &Cup{5, cup4}
	cup2 := &Cup{2, cup5}
	cup1 := &Cup{1, cup2}
	cup9 := &Cup{9, cup1}
	cup8 := &Cup{8, cup9}
	cup3 := &Cup{3, cup8}
	cup7.next = cup3
	game := Game{
		cup3,
		[]*Cup{nil, cup1, cup2, cup3, cup4, cup5, cup6, cup7, cup8, cup9},
	}
	game.playOneMove()
	assert.Equal(t, cup2, cup3.next)
	assert.Equal(t, cup8, cup2.next)
	assert.Equal(t, cup9, cup8.next)
	assert.Equal(t, cup1, cup9.next)
	assert.Equal(t, cup5, cup1.next)
	assert.Equal(t, cup4, cup5.next)
	assert.Equal(t, cup6, cup4.next)
	assert.Equal(t, cup7, cup6.next)
	assert.Equal(t, cup3, cup7.next)
	assert.Equal(t, cup2, game.currentCup)
}

func TestGame_play_100moves(t *testing.T) {
	cup7 := &Cup{7, nil}
	cup6 := &Cup{6, cup7}
	cup4 := &Cup{4, cup6}
	cup5 := &Cup{5, cup4}
	cup2 := &Cup{2, cup5}
	cup1 := &Cup{1, cup2}
	cup9 := &Cup{9, cup1}
	cup8 := &Cup{8, cup9}
	cup3 := &Cup{3, cup8}
	cup7.next = cup3
	game := Game{
		cup3,
		[]*Cup{nil, cup1, cup2, cup3, cup4, cup5, cup6, cup7, cup8, cup9},
	}
	game.play(100)
	assert.Equal(t, cup6, cup1.next)
	assert.Equal(t, cup7, cup6.next)
	assert.Equal(t, cup3, cup7.next)
	assert.Equal(t, cup8, cup3.next)
	assert.Equal(t, cup4, cup8.next)
	assert.Equal(t, cup5, cup4.next)
	assert.Equal(t, cup2, cup5.next)
	assert.Equal(t, cup9, cup2.next)
	assert.Equal(t, cup1, cup9.next)
}

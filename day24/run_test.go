package day24

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []Path
	}{
		{"one letter directions", "ewe", []Path{{East, West, East}}},
		{"two letter directions", "eseswnenw", []Path{{East, SouthEast, SouthWest, NorthEast, NorthWest}}},
		{"multiline", "ewe\nwew", []Path{{East, West, East}, {West, East, West}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := strings.NewReader(tt.input)
			if got := parseFile(reader); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath_getTargetPosition(t *testing.T) {
	tests := []struct {
		name string
		path Path
		want Position
	}{
		{"horizontal", Path{East, West, East, East}, Position{2, 0}},
		{"vertival", Path{East, SouthWest, SouthEast, NorthEast, NorthWest}, Position{1, 0}},
		{"negative", Path{NorthWest, West}, Position{-1.5, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.path.getTargetPosition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTargetPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_computeBlackTiles(t *testing.T) {
	paths := []Path{
		{East, West, East, East},
		{East, SouthWest, SouthEast, NorthEast, NorthWest},
		{NorthWest, West},
		{East, East},
	}
	want := TileList{
		Position{1, 0}:     struct{}{},
		Position{-1.5, -1}: struct{}{},
	}
	got := computeBlackTiles(paths)
	assert.Equal(t, want, got)
}

func TestPosition_shift(t *testing.T) {
	tests := []struct {
		name     string
		position Position
		offset   Offset
		want     Position
	}{
		{"positive", Position{1, 2}, Offset{10, 20}, Position{11, 22}},
		{"negative", Position{1, 2}, Offset{-10, -20}, Position{-9, -18}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.position.shift(tt.offset))
		})
	}
}

func Test_flipTiles(t *testing.T) {
	tests := []struct {
		name       string
		blackTiles []Position
		want       []Position
	}{
		{"one isolated black", []Position{{0, 0}}, []Position{}},
		{"two adjacent blacks", []Position{{0, 0}, {1, 0}}, []Position{{0, 0}, {1, 0}, {0.5, -1}, {0.5, 1}}},
		{"two non-adjacent blacks", []Position{{2, 1}, {3.5, 2}}, []Position{{2.5, 2}, {3, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			blackTiles := TileList{}
			for _, tile := range tt.blackTiles {
				blackTiles[tile] = struct{}{}
			}
			want := TileList{}
			for _, tile := range tt.want {
				want[tile] = struct{}{}
			}
			assert.Equal(t, want, flipTiles(blackTiles))
		})
	}
}

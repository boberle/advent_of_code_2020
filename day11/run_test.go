package day11

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_parseFile(t *testing.T) {
	tests := []struct {
		name    string
		reader  io.Reader
		want    Game
		wantErr bool
	}{
		{"case 1", strings.NewReader("#.#\nL.L\n#..\n.L#"), Game{
			board:         Board{{O, F, O}, {E, F, E}, {O, F, F}, {F, E, O}},
			width:         3,
			height:        4,
			occupiedCount: 4,
		}, false},
		{"unknown tile", strings.NewReader("#.!\nL.L\n#..\n.L#"), Game{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFile(tt.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_countOccupiedAdjacentSeats(t *testing.T) {
	game := &Game{
		board: Board{
			{O, O, O, F, O, F},
			{O, E, O, F, O, F},
			{O, O, O, F, O, F},
		},
		width:  6,
		height: 3,
	}
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"case 0, 0", 0, 0, 2},
		{"case 1, 1", 1, 1, 8},
		{"case 2, 1", 2, 1, 4},
		{"case 3, 1", 3, 1, 6},
		{"case 4, 1", 4, 1, 2},
		{"case 5, 1", 5, 1, 3},
		{"case 0, 2", 0, 2, 2},
		{"case 5, 2", 5, 2, 2},
		{"case 5, 0", 5, 0, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.countOccupiedAdjacentSeats(SeatPosition{tt.x, tt.y}); got != tt.want {
				t.Errorf("countOccupiedAdjacentSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_computeNextRound_rules1(t *testing.T) {
	tests := []struct {
		name string
		game Game
		want Game
	}{
		{"case 1", Game{board: Board{
			{E, F, E, E},
			{E, E, E, E},
			{E, E, F, E},
		}, width: 4, height: 3}, Game{board: Board{
			{O, F, O, O},
			{O, O, O, O},
			{O, O, F, O},
		}, width: 4, height: 3, occupiedCount: 10}},
		{"case 2", Game{board: Board{
			{O, F, O, O},
			{O, O, O, O},
			{O, O, F, O},
		}, width: 4, height: 3, occupiedCount: 10}, Game{board: Board{
			{O, F, E, O},
			{E, E, E, E},
			{O, E, F, O},
		}, width: 4, height: 3, occupiedCount: 4}},
	}
	for _, tt := range tests {
		rules := Rules{
			numberOfOccupiedSeatToBeEmpty: 4,
			countOccupiedSeat:             (*Game).countOccupiedAdjacentSeats,
		}
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.game.computeNextRound(rules); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeNextRound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_hasSameBoardAs(t *testing.T) {
	tests := []struct {
		name  string
		game1 Game
		game2 Game
		want  bool
	}{
		{"case 2", Game{board: Board{
			{O, F, E, O},
			{E, E, E, E},
			{O, E, F, O},
		}, width: 4, height: 3, occupiedCount: 10}, Game{board: Board{
			{O, F, E, O},
			{E, E, E, E},
			{O, E, F, O},
		}, width: 4, height: 3, occupiedCount: 4}, true},
		{"case 2", Game{board: Board{
			{O, F, E, O},
			{E, E, E, E},
			{O, E, O, O},
		}, width: 4, height: 3, occupiedCount: 10}, Game{board: Board{
			{O, F, E, O},
			{E, E, E, E},
			{O, E, F, O},
		}, width: 4, height: 3, occupiedCount: 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.game1.hasSameBoardAs(&tt.game2); got != tt.want {
				t.Errorf("hasSameBoardAs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_findVisibleSeats(t *testing.T) {
	tests := []struct {
		name string
		game Game
		want *VisibleSeatPositions
	}{
		{"case 1", Game{board: Board{
			{O, E, F, F},
			{F, F, F, O},
			{O, F, E, F},
		}, width: 4, height: 3}, &VisibleSeatPositions{
			{0, 0}: []SeatPosition{{1, 0}, {0, 2}, {2, 2}},
			{1, 0}: []SeatPosition{{0, 0}},
			{0, 2}: []SeatPosition{{0, 0}, {2, 2}},
			{2, 2}: []SeatPosition{{0, 0}, {3, 1}, {0, 2}},
			{3, 1}: []SeatPosition{{2, 2}},
		}},
		{"diagonal", Game{board: Board{
			{O, F},
			{F, O},
		}, width: 2, height: 2}, &VisibleSeatPositions{
			{0, 0}: []SeatPosition{{1, 1}},
			{1, 1}: []SeatPosition{{0, 0}},
		}},
		{"diagonal 2", Game{board: Board{
			{F, O},
			{O, F},
		}, width: 2, height: 2}, &VisibleSeatPositions{
			{0, 1}: []SeatPosition{{1, 0}},
			{1, 0}: []SeatPosition{{0, 1}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.game.findVisibleSeats(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findVisibleSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGame_countOccupiedVisibleSeats(t *testing.T) {
	game := Game{
		board: Board{
			{O, E, F, F},
			{F, F, F, O},
			{O, F, E, F},
		},
		width:  4,
		height: 3,
		visibleSeats: &VisibleSeatPositions{
			{0, 0}: []SeatPosition{{1, 0}, {0, 2}, {2, 2}},
			{1, 0}: []SeatPosition{{0, 0}},
			{0, 2}: []SeatPosition{{0, 0}, {2, 2}},
			{2, 2}: []SeatPosition{{0, 0}, {3, 1}, {0, 2}},
			{3, 1}: []SeatPosition{{2, 2}},
		},
	}
	tests := []struct {
		name string
		seat SeatPosition
		want int
	}{
		{"case 0, 0", SeatPosition{0, 0}, 1},
		{"case 1, 0", SeatPosition{1, 0}, 1},
		{"case 0, 2", SeatPosition{0, 2}, 1},
		{"case 2, 2", SeatPosition{2, 2}, 3},
		{"case 3, 1", SeatPosition{3, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := game.countOccupiedVisibleSeats(tt.seat); got != tt.want {
				t.Errorf("countOccupiedVisibleSeats() = %v, want %v", got, tt.want)
			}
		})
	}
}

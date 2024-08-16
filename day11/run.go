package day11

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

type Tile int

const (
	Floor    Tile = iota
	Empty    Tile = iota
	Occupied Tile = iota
)

const (
	F Tile = Floor
	E Tile = Empty
	O Tile = Occupied
)

type Board [][]Tile

type SeatPosition struct {
	x int
	y int
}

type VisibleSeatPositions map[SeatPosition][]SeatPosition

type Game struct {
	board         Board
	width         int
	height        int
	occupiedCount int
	visibleSeats  *VisibleSeatPositions
}

type Rules struct {
	numberOfOccupiedSeatToBeEmpty int
	// https://stackoverflow.com/questions/51780781/function-type-with-a-receiver
	countOccupiedSeat func(*Game, SeatPosition) int
}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		log.Fatalln(err)
	}
	defer fh.Close()

	game, err := parseFile(fh)
	if err != nil {
		log.Fatalln(err)
	}

	var occupiedSeatCount int

	rulesPart1 := Rules{
		numberOfOccupiedSeatToBeEmpty: 4,
		countOccupiedSeat:             (*Game).countOccupiedAdjacentSeats,
	}
	occupiedSeatCount = game.playTillStabilization(rulesPart1)
	fmt.Printf("Part 1: number of occuped seats: %d\n", occupiedSeatCount)

	game.visibleSeats = game.findVisibleSeats()
	rulesPart2 := Rules{
		numberOfOccupiedSeatToBeEmpty: 5,
		countOccupiedSeat:             (*Game).countOccupiedVisibleSeats,
	}
	occupiedSeatCount = game.playTillStabilization(rulesPart2)
	fmt.Printf("Part 2: number of occuped seats: %d\n", occupiedSeatCount)

}

func parseFile(reader io.Reader) (Game, error) {
	game := Game{}
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		if game.width == 0 {
			game.width = len(line)
		}
		game.board = append(game.board, []Tile{})
		for _, c := range line {
			var tile Tile
			switch c {
			case '.':
				tile = Floor
			case 'L':
				tile = Empty
			case '#':
				tile = Occupied
				game.occupiedCount++
			default:
				return Game{}, errors.New(fmt.Sprintf("unknown tile: %v\n", c))
			}
			game.board[game.height] = append(game.board[game.height], tile)
		}
		game.height++
	}

	return game, nil
}

func (game *Game) findVisibleSeats() *VisibleSeatPositions {
	seats := VisibleSeatPositions{}
	for y := 0; y < game.height; y++ {
		for x := 0; x < game.width; x++ {
			if game.board[y][x] == Floor {
				continue
			}

			for i := 1; x+i < game.width; i++ {
				if game.board[y][x+i] != Floor {
					var pos SeatPosition
					pos = SeatPosition{x, y}
					seats[pos] = append(seats[pos], SeatPosition{x + i, y})
					pos = SeatPosition{x + i, y}
					seats[pos] = append(seats[pos], SeatPosition{x, y})
					break
				}
			}

			for i := 1; y+i < game.height; i++ {
				if game.board[y+i][x] != Floor {
					var pos SeatPosition
					pos = SeatPosition{x, y}
					seats[pos] = append(seats[pos], SeatPosition{x, y + i})
					pos = SeatPosition{x, y + i}
					seats[pos] = append(seats[pos], SeatPosition{x, y})
					break
				}
			}

			for i := 1; x+i < game.width && y+i < game.height; i++ {
				if game.board[y+i][x+i] != Floor {
					var pos SeatPosition
					pos = SeatPosition{x, y}
					seats[pos] = append(seats[pos], SeatPosition{x + i, y + i})
					pos = SeatPosition{x + i, y + i}
					seats[pos] = append(seats[pos], SeatPosition{x, y})
					break
				}
			}

			for i := 1; x-i >= 0 && y+i < game.height; i++ {
				if game.board[y+i][x-i] != Floor {
					var pos SeatPosition
					pos = SeatPosition{x, y}
					seats[pos] = append(seats[pos], SeatPosition{x - i, y + i})
					pos = SeatPosition{x - i, y + i}
					seats[pos] = append(seats[pos], SeatPosition{x, y})
					break
				}
			}

		}
	}
	return &seats
}

func (game *Game) playTillStabilization(rules Rules) int {
	lastGame := *game
	for {
		newGame := lastGame.computeNextRound(rules)
		if newGame.hasSameBoardAs(&lastGame) {
			return newGame.occupiedCount
		}
		lastGame = newGame
	}
}

func (game *Game) hasSameBoardAs(other *Game) bool {
	for y := 0; y < game.height; y++ {
		for x := 0; x < game.width; x++ {
			if game.board[y][x] != other.board[y][x] {
				return false
			}
		}
	}
	return true
}

func (game *Game) computeNextRound(rules Rules) Game {
	newGame := game.clone()
	for y := 0; y < game.height; y++ {
		for x := 0; x < game.width; x++ {
			seat := game.board[y][x]
			occupiedAdjacentSeats := rules.countOccupiedSeat(game, SeatPosition{x, y})
			if seat == Empty && occupiedAdjacentSeats == 0 {
				newGame.board[y][x] = Occupied
				newGame.occupiedCount++
			} else if seat == Occupied && occupiedAdjacentSeats >= rules.numberOfOccupiedSeatToBeEmpty {
				newGame.board[y][x] = Empty
				newGame.occupiedCount--
			}
		}
	}
	return newGame
}

func (game *Game) countOccupiedAdjacentSeats(seat SeatPosition) int {
	rv := 0
	x := seat.x
	y := seat.y
	if x > 0 && y > 0 && game.board[y-1][x-1] == Occupied {
		rv++
	}
	if x < game.width-1 && y < game.height-1 && game.board[y+1][x+1] == Occupied {
		rv++
	}
	if x > 0 && game.board[y][x-1] == Occupied {
		rv++
	}
	if x < game.width-1 && game.board[y][x+1] == Occupied {
		rv++
	}
	if y > 0 && game.board[y-1][x] == Occupied {
		rv++
	}
	if y < game.height-1 && game.board[y+1][x] == Occupied {
		rv++
	}
	if x < game.width-1 && y > 0 && game.board[y-1][x+1] == Occupied {
		rv++
	}
	if x > 0 && y < game.height-1 && game.board[y+1][x-1] == Occupied {
		rv++
	}
	return rv
}

func (game *Game) countOccupiedVisibleSeats(seat SeatPosition) int {
	total := 0
	for _, visibleSeat := range (*game.visibleSeats)[seat] {
		if game.board[visibleSeat.y][visibleSeat.x] == Occupied {
			total++
		}
	}
	return total
}

func (board *Board) clone() Board {
	sy := len(*board)
	newBoard := make([][]Tile, sy)
	for y := 0; y < sy; y++ {
		sx := len((*board)[y])
		newBoard[y] = make([]Tile, sx)
		for x := 0; x < sx; x++ {
			newBoard[y][x] = (*board)[y][x]
		}
	}
	return newBoard
}

func (game *Game) clone() Game {
	return Game{
		board:         game.board.clone(),
		width:         game.width,
		height:        game.height,
		occupiedCount: game.occupiedCount,
		visibleSeats:  game.visibleSeats,
	}
}

func (game *Game) print() {
	for y := 0; y < game.height; y++ {
		for x := 0; x < game.width; x++ {
			switch game.board[y][x] {
			case Occupied:
				fmt.Printf("#")
			case Empty:
				fmt.Printf("L")
			case Floor:
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("--------")
}

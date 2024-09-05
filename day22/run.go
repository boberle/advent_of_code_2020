package day22

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Card int

type Deck []Card

type Game struct {
	player1Deck Deck
	player2Deck Deck
}

func Run(infile string) {
	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	game := parseFile(fh)
	winningDeck1 := game.play()
	fmt.Printf("Winning player's score (part 1): %d\n", winningDeck1.computeScore())

	winningDeck2, _ := game.recursivePlay()
	fmt.Printf("Winning player's score (part 2): %d\n", winningDeck2.computeScore())

}

func parseFile(reader io.Reader) Game {
	game := Game{
		player1Deck: make(Deck, 0),
		player2Deck: make(Deck, 0),
	}

	scanner := bufio.NewScanner(reader)
	var currentDeck *Deck
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// nothing
		} else if line == "Player 1:" {
			currentDeck = &game.player1Deck
		} else if line == "Player 2:" {
			currentDeck = &game.player2Deck
		} else {
			card, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			*currentDeck = append(*currentDeck, Card(card))
		}
	}

	if len(game.player1Deck) != len(game.player2Deck) {
		panic("length mismatch")
	}

	return game
}

func (game *Game) play() Deck {
	deck1 := game.player1Deck
	deck2 := game.player2Deck
	for {
		if len(deck1) == 0 {
			return deck2
		} else if len(deck2) == 0 {
			return deck1
		}
		if deck1[0] > deck2[0] {
			deck1 = append(deck1, deck1[0], deck2[0])
		} else {
			deck2 = append(deck2, deck2[0], deck1[0])
		}
		deck1 = append(Deck{}, deck1[1:]...)
		deck2 = append(Deck{}, deck2[1:]...)
	}
}

func (deck *Deck) computeScore() int {
	score := 0
	l := len(*deck)
	for i, card := range *deck {
		score += int(card) * (l - i)
	}
	return score
}

func (game *Game) recursivePlay() (Deck, int) {
	deck1 := game.player1Deck
	deck2 := game.player2Deck
	history := make([]Game, 0)

	for {
		if len(deck1) == 0 {
			return deck2, 2
		} else if len(deck2) == 0 {
			return deck1, 1
		}

		currentGame := Game{deck1, deck2}
		if isGameInHistory(&history, &currentGame) {
			return deck1, 1
		}
		history = append(history, currentGame)

		card1 := deck1[0]
		card2 := deck2[0]
		deck1 = append(Deck{}, deck1[1:]...)
		deck2 = append(Deck{}, deck2[1:]...)

		winner := 0
		if len(deck1) >= int(card1) && len(deck2) >= int(card2) {
			newGame := Game{deck1[:int(card1)], deck2[:int(card2)]}
			_, winner = newGame.recursivePlay()
		} else {
			if card1 > card2 {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			deck1 = append(deck1, card1, card2)
		} else if winner == 2 {
			deck2 = append(deck2, card2, card1)
		} else {
			panic(fmt.Sprintf("unknown winner: %d", winner))
		}
	}
}

func isGameInHistory(history *[]Game, game *Game) bool {
	for _, g := range *history {
		if g.player1Deck.isEqual(&game.player1Deck) && g.player2Deck.isEqual(&game.player2Deck) {
			return true
		}
	}
	return false
}

func (deck *Deck) isEqual(other *Deck) bool {
	if len(*deck) != len(*other) {
		return false
	}
	for i, card := range *deck {
		if card != (*other)[i] {
			return false
		}
	}
	return true
}

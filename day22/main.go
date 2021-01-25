package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := Input("day22/input.txt")
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type Deck struct {
	player int
	cards []int
}

func (d *Deck) Peek() int {
	return d.cards[0]
}

func (d *Deck) Add(c1, c2 int) {
	d.cards = append(d.cards, Max(c1, c2), Min(c1, c2))
}

func (d *Deck) Pop() int {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) Size() int {
	return len(d.cards)
}

func (d *Deck) String() string {
	return fmt.Sprintf("Player %d's deck: %s", d.player, StringFromIntSlice(d.cards))
}

func Draw(p1 *Deck, p2 *Deck) *Deck {
	if p1.Peek() > p2.Peek() {
		return p1
	}
	return p2
}

func (d *Deck) Score() int {
	score := 0
	size := d.Size()
	for i := 0; i < size; i++ {
		score += d.cards[i] * (size-i)
	}
	return score
}

func (d *Deck) CopySubset(n int) *Deck {
	return &Deck{player: d.player, cards: d.cards[:n]}
}

func Solve1(input string) int {
	p1, p2 := ParseInput(input)

	var winner *Deck
	round := 1
	for p1.Size() > 0 && p2.Size() > 0 {
		winner = Draw(p1, p2)
		winner.Add(p1.Pop(), p2.Pop())
		round++
	}

	if p1.Size() == 0 {
		return p2.Score()
	}
	return p1.Score()
}

func Solve2(input string) int {
	p1, p2 := ParseInput(input)

	// Before either player deals a card, if there was a previous round in this game that had exactly the same cards
	// in the same order in the same players' decks, the game instantly ends in a win for player 1.

	// Otherwise, this round's cards must be in a new configuration;
	// the players begin the round by each drawing the top card of their deck as normal.

	// If both players have at least as many cards remaining in their deck as the value of the card they just drew,
	// the winner of the round is determined by playing a new game of Recursive Combat

	// Otherwise, at least one player must not have enough cards left in their deck to recurse;
	// the winner of the round is the player with the higher-value card.

	// To play a sub-game of Recursive Combat, each player creates a new deck by making a copy of the next cards in their deck
	// (the quantity of cards copied is equal to the number on the card they drew to trigger the sub-game).

	return RecursiveCombat(1, p1, p2).Score()
}

func RecursiveCombat(game int, p1, p2 *Deck) *Deck {
	Printfln("=== Game %d ===", game)

	seen := NewStringSet()
	var winner *Deck
	round := 1
	for p1.Size() > 0 && p2.Size() > 0 {
		key := p1.String() + "&" + p2.String()
		if seen.Contains(key) {
			return p1
		}
		seen.Add(key)

		Printfln("\n-- Round %d (Game %d) --", round, game)
		if isBreakpoint(p1,p2,game,round) {
			time.Sleep(100)
		}
		Printfln(p1.String())
		Printfln(p2.String())

		c1 := p1.Pop()
		c2 := p2.Pop()

		Printfln("Player 1 plays: %d", c1)
		Printfln("Player 2 plays: %d", c2)

		if p1.Size() >= c1 && p2.Size() >= c2 {
			Printfln("Playing a sub-game to determine the winner...\n")
			winner = RecursiveCombat(game+1, p1.CopySubset(c1), p2.CopySubset(c2))
		} else {
			if c1 > c2 {
				winner = p1
			} else {
				winner = p2
			}
		}
		if winner.player == 1 {
			p1.Add(c1, c2)
		} else {
			p2.Add(c1, c2)
		}
		//winner.Add(c1, c2)
		Printfln("Player %d wins round %d of game %d!", winner.player, round, game)
		round++
	}

	if p1.Size() == 0 {
		winner = p2
	} else {
		winner = p1
	}
	Printfln("The winner of game %d is player %d!\n", game, winner.player)
	if game > 1 {
		Printfln("...anyway, back to game %d.", game-1)
	}
	return winner
}

func isBreakpoint(p1, p2 *Deck, game int, round int) bool {
	if game == 1 && round == 10 {
		return true
	}
	return false
}

func ParseInput(input string) (*Deck, *Deck) {
	s := SplitByEmptyNewline(input)
	p1 := Ints(strings.Join(Lines(s[0])[1:], ","))
	p2 := Ints(strings.Join(Lines(s[1])[1:], ","))
	return &Deck{player: 1, cards: p1}, &Deck{player: 2, cards: p2}
}

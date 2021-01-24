package main

import (
	. "aoc"
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
	return fmt.Sprintf("Player %d's deck: %s",
		d.player, strings.Join(strings.Fields(fmt.Sprint(d.cards)), ", "))
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

	winner := RecursiveCombat(1, p1, p2, []Deck{})

	if p1.Size() == 0 {
		return p2.Score()
	}
	return p1.Score()
}

func RecursiveCombat(game int, p1, p2 *Deck, history []Deck) *Deck {
	// todo: check if previous round had exact same cards
	if playedBefore(p1, history) || playedBefore(p2, history) {
		return p1
	}

	c1 := p1.Pop()
	c2 := p2.Pop()
	if c1 == p1.Size() && c2 == p2.Size() {
		return RecursiveCombat(game+1, p1.CopySubset(c1), p2.CopySubset(c2), history)
	} else {
		return Draw(p1, p2)
	}
}

func playedBefore(d *Deck, history []Deck) bool {
	for _, h := range history {
		if equal(d.cards, h.cards) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ParseInput(input string) (*Deck, *Deck) {
	s := SplitByEmptyNewline(input)
	p1 := Ints(strings.Join(Lines(s[0])[1:], ","))
	p2 := Ints(strings.Join(Lines(s[1])[1:], ","))
	return &Deck{player: 1, cards: p1}, &Deck{player: 2, cards: p2}
}

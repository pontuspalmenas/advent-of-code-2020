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
	cards  []int
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
		score += d.cards[i] * (size - i)
	}
	return score
}

func (d *Deck) CopySubset(n int) *Deck {
	cards := make([]int, n)
	copy(cards, d.cards[:n])
	return &Deck{player: d.player, cards: cards}
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
	return RecursiveCombat(p1, p2).Score()
}

func RecursiveCombat(p1, p2 *Deck) *Deck {
	seen := NewStringSet()
	var winner *Deck
	for p1.Size() > 0 && p2.Size() > 0 {
		key := p1.String() + "&" + p2.String()
		if seen.Contains(key) {
			return p1
		}
		seen.Add(key)

		c1 := p1.Pop()
		c2 := p2.Pop()

		if p1.Size() >= c1 && p2.Size() >= c2 {
			winner = RecursiveCombat(p1.CopySubset(c1), p2.CopySubset(c2))
		} else {
			if c1 > c2 {
				winner = p1
			} else {
				winner = p2
			}
		}
		if winner.player == 1 {
			p1.cards = append(p1.cards, c1, c2)
		} else {
			p2.cards = append(p2.cards, c2, c1)
		}
	}

	if p1.Size() == 0 {
		return p2
	}
	return p1
}

func ParseInput(input string) (*Deck, *Deck) {
	s := SplitByEmptyNewline(input)
	p1 := Ints(strings.Join(Lines(s[0])[1:], ","))
	p2 := Ints(strings.Join(Lines(s[1])[1:], ","))
	return &Deck{player: 1, cards: p1}, &Deck{player: 2, cards: p2}
}

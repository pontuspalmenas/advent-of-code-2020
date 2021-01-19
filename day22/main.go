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

func (d *Deck) Top() int {
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
	if p1.Top() > p2.Top() {
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

func Solve1(input string) int {
	p1, p2 := ParseInput(input)

	var winner *Deck
	round := 1
	for p1.Size() > 0 && p2.Size() > 0 {
		Printfln("-- Round %d --", round)
		Printfln(p1.String())
		Printfln(p2.String())
		Printfln("Player 1 plays: %d", p1.Top())
		Printfln("Player 2 plays: %d", p2.Top())
		winner = Draw(p1, p2)
		Printfln("Player %d wins the round!", winner.player)
		winner.Add(p1.Pop(), p2.Pop())
		round++
	}

	Printfln("== Post-game results ==")
	Printfln(p1.String())
	Printfln(p2.String())

	if p1.Size() == 0 {
		return p2.Score()
	}
	return p1.Score()
}

func Solve2(input string) int {
	return 0
}

func ParseInput(input string) (*Deck, *Deck) {
	s := SplitByEmptyNewline(input)
	p1 := Ints(strings.Join(Lines(s[0])[1:], ","))
	p2 := Ints(strings.Join(Lines(s[1])[1:], ","))
	return &Deck{player: 1, cards: p1}, &Deck{player: 2, cards: p2}
}

package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	p1 := Solve1("685974213")
	p2 := Solve2("685974213")
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type Circle struct {
	head *Cup
}

func NewCircle() *Circle {
	return &Circle{head: nil}
}

func (c *Circle) add(v int) {
	if c.head == nil {
		c.head = &Cup{val: v}
		c.head.next = c.head
		return
	}

	curr := c.head
	for curr.next != c.head && curr.next != nil {
		curr = curr.next
	}

	curr.val = v
	curr.next = c.head
}

type Cup struct {
	val int
	next *Cup
}

func Solve1(input string) int {
	circle := InitState(input)
	Printfln("cups: %s", circle.String())

	return 0
}

func InitState(input string) *Circle {
	c := NewCircle()

	for _, ch := range input {
		c.add(Int(string(ch)))
	}
	return c
}

func (c *Circle) String() string {
	curr := c.head
	s := fmt.Sprintf("%d ", curr.val)
	for curr.next != c.head && curr.next != nil {
		curr = curr.next
		s += fmt.Sprintf("%d ", curr.val)
	}

	return s
}

func Solve2(input string) int {
	return 0
}

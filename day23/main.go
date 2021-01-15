package main

import (
	. "aoc"
	. "aoc/types"
	"container/ring"
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

func Solve1(input string) string {
	dict, cups := Setup(input)

	curr := cups
	size := curr.Len()
	for i := 1; i <= 100; i++ {
		move(dict, curr, size)
	}

	p := dict[1]
	ans := ""
	for i := 1; i < len(input); i++ {
		p = p.Next()
		ans += fmt.Sprintf("%d", p.Value.(int))
	}

	return ans
}

func move(dict map[int]*ring.Ring, curr *ring.Ring, size int) {
	removedCups := curr.Unlink(3)
	target := 1 + ((size + curr.Value.(int) - 2) % size)
	removed := NewIntSet()
	removed.Add(removedCups.Move(0).Value.(int))
	removed.Add(removedCups.Move(1).Value.(int))
	removed.Add(removedCups.Move(2).Value.(int))

	for removed.Contains(target) {
		target = 1 + ((size + target - 2) % size)
	}

	dict[target].Link(removedCups)
	curr = curr.Next()
}

func Setup(input string) (map[int]*ring.Ring, *ring.Ring) {
	dict := make(map[int]*ring.Ring)
	cups := ring.New(len(input))
	for _, c := range input {
		n := Int(string(c))
		cups.Value = n
		dict[n] = cups
		cups = cups.Next()
	}
	return dict, cups
}

const Mil = 1000000

func Solve2(input string) int {
	dict, cups := Setup(input)

	last := cups.Prev()
	added := ring.New(Mil - len(input))
	for i:=len(input)+1; i<=Mil; i++ {
		added.Value = i
		dict[i] = added
		added = added.Next()
	}
	last.Link(added)

	curr := cups
	size := curr.Len()
	for i := 1; i <= 10*Mil; i++ {
		removedCups := curr.Unlink(3)
		target := 1 + ((size + curr.Value.(int) - 2) % size)
		removed := NewIntSet()
		removed.Add(removedCups.Move(0).Value.(int))
		removed.Add(removedCups.Move(1).Value.(int))
		removed.Add(removedCups.Move(2).Value.(int))

		for removed.Contains(target) {
			target = 1 + ((size + target - 2) % size)
		}

		dict[target].Link(removedCups)
		curr = curr.Next()
	}

	p := dict[1]
	return p.Move(1).Value.(int) * p.Move(2).Value.(int)
}

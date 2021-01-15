package main

import (
	. "aoc"
	"container/ring"
	"fmt"
	"time"
)


// hints by lizthegrey
func main() {
	start := time.Now()
	p1 := Solve1("685974213")
	p2 := Solve2("685974213")
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

// todo: don't want these global, but if I send them as args to move(),
//		 Solve2 breaks for some mysterious reason. map is actually a pointer
//		 so it _should_ work, but Go really do be like that sometimes.
var dict map[int]*ring.Ring
var curr *ring.Ring

func Solve1(input string) string {
	var cups *ring.Ring
	dict, cups = setup(input)

	curr = cups

	for i := 1; i <= 100; i++ {
		move()
	}

	p := dict[1]
	ans := ""
	for i := 1; i < len(input); i++ {
		p = p.Next()
		ans += fmt.Sprintf("%d", p.Value.(int))
	}

	return ans
}

const Mil = 1000000

func Solve2(input string) int {
	var cups *ring.Ring
	dict, cups = setup(input)

	last := cups.Prev()
	added := ring.New(Mil - len(input))
	for i:=len(input)+1; i<=Mil; i++ {
		added.Value = i
		dict[i] = added
		added = added.Next()
	}
	last.Link(added)

	curr = cups
	for i := 1; i <= 10*Mil; i++ {
		move()
	}

	p := dict[1]
	return p.Move(1).Value.(int) * p.Move(2).Value.(int)
}

func setup(input string) (map[int]*ring.Ring, *ring.Ring) {
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

func move() {
	removed := curr.Unlink(3)
	target := curr.Value.(int) - 1
	target = wrap(target)
	for contains(removed, target) {
		target = wrap(target-1)
	}

	dict[target].Link(removed)
	curr = curr.Next()
}

func contains(ring *ring.Ring, v int) bool {
	r := ring
	for i:=0; i < ring.Len(); i++ {
		if r.Value.(int) == v {
			return true
		}
		r = r.Next()
	}
	return false
}

func wrap(i int) int {
	if i == 0 {
		return Mil
	}
	return i
}

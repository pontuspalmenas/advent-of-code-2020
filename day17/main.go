package main

import (
	. "aoc"
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day17/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type P3 struct {
	z, y, x int
}

type Universe3 struct {
	state map[P3]bool
}

func NewUniverse3() *Universe3 {
	state := make(map[P3]bool)
	return &Universe3{state}
}

func (u *Universe3) set(p P3) {
	u.state[p] = true
}

func (u *Universe3) unset(p P3) {
	u.state[p] = false
}

func (u *Universe3) count() int {
	count := 0
	for _, active := range u.state {
		if active {
			count++
		}
	}
	return count
}

func (u *Universe3) activeNeighbors(p P3) int {
	count := 0
	for z := p.z-1; z <= p.z+1; z++ {
		for y := p.y-1; y <= p.y+1; y++ {
			for x := p.x-1; x <= p.x+1; x++ {
				np := P3{z: z, y: y, x: x}
				if samePoint(p, np) {
					continue
				}
				if u.active(np) {
					count++
				}
			}
		}
	}
	return count
}

func (u *Universe3) active(p P3) bool {
	return u.state[p]
}

func Solve1(input []string) int {
	u := NewUniverse3()
	// init
	for i, row := range input {
		for j, col := range row {
			if col == '#' {
				u.set(P3{y: i, x: j, z: 0})
			} else {
				u.unset(P3{y: i, x: j, z: 0})
			}
		}
	}

	// starting size
	size := len(input[0])

	for gen := 1; gen <= 6; gen++ {
		newState := u.NewState()
		for z := -gen; z <= gen; z++ {
			for y := -gen; y < size+gen; y++ {
				for x := -gen; x < size+gen; x++ {
					p := P3{z: z, y: y, x: x}
					activeNs := u.activeNeighbors(p)
					if u.active(p) {
						if !(activeNs == 2 || activeNs == 3) {
							newState[p] = false
						}
					} else {
						if activeNs == 3 {
							newState[p] = true
						}
					}
				}
			}
		}
		u.state = newState
	}

	return u.count()
}

func (u *Universe3) NewState() map[P3]bool {
	newState := make(map[P3]bool)
	for k, v := range u.state {
		newState[k] = v
	}
	return newState
}

func Solve2(input []string) int {
	return 0
}

func samePoint(p1 P3, p2 P3) bool {
	return p1.z == p2.z &&
		p1.y == p2.y &&
		p1.x == p2.x
}

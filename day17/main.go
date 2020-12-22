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

// todo: want a generic PointN with dim []int to solve both problems, but can't use slice as map key...

type P3 struct {
	z, y, x int
}

type P4 struct {
	z, y, x, w int
}

type Universe3 struct {
	state map[P3]bool
}

type Universe4 struct {
	state map[P4]bool
}

func NewUniverse3() *Universe3 {
	state := make(map[P3]bool)
	return &Universe3{state}
}

func NewUniverse4() *Universe4 {
	state := make(map[P4]bool)
	return &Universe4{state}
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

func (u *Universe4) count() int {
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
				if p == np {
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

func (u *Universe4) activeNeighbors(p P4) int {
	count := 0
	for z := p.z-1; z <= p.z+1; z++ {
		for y := p.y-1; y <= p.y+1; y++ {
			for x := p.x-1; x <= p.x+1; x++ {
				for w := p.w-1; w <= p.w+1; w++ {
					np := P4{z: z, y: y, x: x, w: w}
					if p == np {
						continue
					}
					if u.active(np) {
						count++
					}
				}
			}
		}
	}
	return count
}

func (u *Universe3) active(p P3) bool {
	return u.state[p]
}

func (u *Universe4) active(p P4) bool {
	return u.state[p]
}

func Solve1(input []string) int {
	u := NewUniverse3()
	// init
	for i, row := range input {
		for j, col := range row {
			u.state[P3{y: i, x: j}] = col == '#'
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

func (u *Universe4) NewState() map[P4]bool {
	newState := make(map[P4]bool)
	for k, v := range u.state {
		newState[k] = v
	}
	return newState
}

func Solve2(input []string) int {
	u := NewUniverse4()
	// init
	for i, row := range input {
		for j, col := range row {
			u.state[P4{y: i, x: j}] = col == '#'
		}
	}

	// starting size
	size := len(input[0])

	for gen := 1; gen <= 6; gen++ {
		newState := u.NewState()
		for z := -gen; z <= gen; z++ {
			for y := -gen; y < size+gen; y++ {
				for x := -gen; x < size+gen; x++ {
					for w := -gen; w < size+gen; w++ {
						p := P4{z: z, y: y, x: x, w: w}
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
		}
		u.state = newState
	}

	return u.count()
}

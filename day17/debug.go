package main

import (
	. "aoc"
	"fmt"
)

func (u *Universe3) print(size int) {
	for z := -1; z < 2; z++ {
		Printfln("z=%d",z)
		for y:=-size; y<size+size; y++ {
			for x:=-size; x<size+size; x++ {
				if u.active(P3{z:z, y:y, x:x}) {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func (u *Universe3) debug(p P3) {
	Printfln("{%d,%d,%d} %v", p.z, p.y, p.x, u.active(p))
}

func (u *Universe3) debugActiveNeighs(p P3) {
	Printfln("{%d,%d,%d} %d", p.z, p.y, p.x, u.activeNeighbors(p))
}

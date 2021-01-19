package main

import (
	. "aoc"
	. "aoc/day08/vm"
	"fmt"
)

func main() {
	input := Lines(Input("day08/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)
}

func Solve1(input []string) int {
	vm := VM{}
	vm.Load(parse(input))
	vm.Run()

	return vm.Accumulator
}

func parse(in []string) map[int]Instruction {
	code := make(map[int]Instruction)
	var op string
	var arg int
	for i, s := range in {
		Sscanf(s, "%s %d", &op, &arg)
		code[i] = Instruction{Op: op, Arg: arg}
	}

	return code
}

func Solve2(input []string) int {
	code := parse(input)
	vm := VM{}

	for i := 0; i < len(code); i++ {
		modCode := copyMap(code)

		if modCode[i].Op == "jmp" {
			modCode[i] = Instruction{Op: "nop", Arg: modCode[i].Arg}
			vm.Load(modCode)
			if vm.Run() == EXIT_SUCCESS {
				return vm.Accumulator
			}
		}
		if modCode[i].Op == "nop" {
			modCode[i] = Instruction{Op: "jmp", Arg: modCode[i].Arg}
			vm.Load(modCode)
			if vm.Run() == EXIT_SUCCESS {
				return vm.Accumulator
			}
		}
	}

	panic("no solution")
}

func copyMap(m map[int]Instruction) map[int]Instruction {
	c := make(map[int]Instruction, len(m))
	for k,v := range m {
		c[k] = v
	}
	return c
}

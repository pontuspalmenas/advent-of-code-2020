package main

import (
	. "aoc"
	"fmt"
)

func main() {
	input := Lines(Input("day08/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("P1: ", p1)
	fmt.Println("P2: ", p2)
}

func Solve1(input []string) int {
	_, acc := run(parse(input))

	return acc
}

type instr struct {
	op string
	arg int
}

func parse(in []string) map[int]instr {
	code := make(map[int]instr)
	var op string
	var arg int
	for i, s := range in {
		Sscanf(s, "%s %d", &op, &arg)
		code[i] = instr{op, arg}
	}

	return code
}

func Solve2(input []string) int {
	code := parse(input)

	for i := 0; i < len(code); i++ {
		modCode := copyMap(code)

		if modCode[i].op == "jmp" {
			modCode[i] = instr{"nop", modCode[i].arg}
			status, acc := run(modCode)
			if status == EXIT_SUCCESS {
				return acc
			}
		}
		if modCode[i].op == "nop" {
			modCode[i] = instr{"jmp", modCode[i].arg}
			status, acc := run(modCode)
			if status == EXIT_SUCCESS {
				return acc
			}
		}
	}

	return -1
}

func copyMap(m map[int]instr) map[int]instr {
	c := make(map[int]instr, len(m))
	for k,v := range m {
		c[k] = v
	}
	return c
}

var EXIT_INFINTE_LOOP = -1
var EXIT_SUCCESS = 0
func run(code map[int]instr) (int, int) {
	acc := 0
	pc := 0

	run := make(map[int]bool)

	for {
		Printfln("%s %d", code[pc].op, code[pc].arg)
		if pc >= len(code) {
			return EXIT_SUCCESS, acc
		}
		if run[pc] {
			return EXIT_INFINTE_LOOP, acc
		}
		run[pc] = true
		switch code[pc].op {
		case "nop":
			pc++
		case "acc":
			acc += code[pc].arg
			pc++
		case "jmp":
			pc += code[pc].arg
		default:
			panic(fmt.Sprintf("undefined opcode: %s", code[pc].op))
		}
	}
}

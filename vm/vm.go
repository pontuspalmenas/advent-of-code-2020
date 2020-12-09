package vm

import (
	"fmt"
	. "aoc"
)

type Instruction struct {
	Op string
	Arg int
}

type VM struct {
	program map[int]Instruction
	Accumulator int
	programCounter int
}

const (
	EXIT_SUCCESS = iota
	EXIT_INFINTE_LOOP
)

func (vm *VM) Load(program map[int]Instruction) {
	vm.program = program
}

// Run resets the accumulator register and runs the loaded program
func (vm *VM) Run() int {
	pc := 0
	vm.Accumulator = 0

	run := make(map[int]bool)

	for {
		Printfln("%s %d", vm.program[pc].Op, vm.program[pc].Arg)
		if pc >= len(vm.program) {
			return EXIT_SUCCESS
		}
		if run[pc] {
			return EXIT_INFINTE_LOOP
		}
		run[pc] = true
		switch vm.program[pc].Op {
		case "nop":
			pc++
		case "acc":
			vm.Accumulator += vm.program[pc].Arg
			pc++
		case "jmp":
			pc += vm.program[pc].Arg
		default:
			panic(fmt.Sprintf("undefined opcode: %s", vm.program[pc].Op))
		}
	}
}

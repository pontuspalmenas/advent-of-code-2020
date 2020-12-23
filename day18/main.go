package main

import (
	. "aoc"
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day18/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type Stack struct {
	d []int
}

func NewStack() *Stack {
	d := make([]int, 0)
	return &Stack{d}
}

func (s *Stack) push(v int) {
	s.d = append(s.d, v)
}

func (s *Stack) pop() int {
	n := s.d[len(s.d)-1]
	s.d = s.d[:len(s.d)-1]
	return n
}

func (s *Stack) peek() int {
	return s.d[len(s.d)-1]
}

const Add = int('+')
const Mul = int('*')
const Lpn = int('(')
const Rpn = int(')')

func Solve1(input []string) int {
	sum := 0
	for _, s := range input {
		s = strings.ReplaceAll(s, " ", "")
		sum += eval(s, part1)
	}
	return sum
}

func Solve2(input []string) int {
	sum := 0
	for _, s := range input {
		s = strings.ReplaceAll(s, " ", "")
		sum += eval(s, part2)
	}
	return sum
}

func eval(s string, condition func(*Stack) bool) int {
	values := NewStack()
	ops := NewStack()
	ops.push(Lpn)

	for _, c := range s {
		if c == '+' {
			calc(ops, values, condition)
			ops.push(Add)
		} else if c == '*' {
			calc(ops, values, condition)
			ops.push(Mul)
		} else if c == '(' {
			ops.push(Lpn)
		} else if c == ')' {
			calc(ops, values, condition)
			ops.pop()
		} else {
			values.push(Int(string(c)))
		}
	}

	calc(ops, values, condition)
	return values.pop()
}

func part1(ops *Stack) bool {
	return ops.peek() != Lpn
}

func part2(ops *Stack) bool {
	return !(ops.peek() == Lpn || ops.peek() == Mul)
}

func calc(ops *Stack, values *Stack, condition func(*Stack) bool) {
	for condition(ops) {
		op := ops.pop()
		v1 := values.pop()
		v2 := values.pop()
		if op == Add {
			values.push(v1 + v2)
		} else if op == Mul {
			values.push(v1 * v2)
		}
	}
}



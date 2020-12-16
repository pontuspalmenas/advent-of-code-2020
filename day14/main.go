package main

import (
	. "aoc"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day14/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

const max = 65485

func Solve1(input []string) int64 {
	var mask string
	//var mem [max+1]int64
	mem := make([]int64, max+1)

	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			mask = strings.Split(s, "mask = ")[1]
		} else {
			op := Regex(`mem\[(\d+)\] = (\d+)`, s)
			mem[Int(op[0])] = binToInt(flip(pad(strToBin(op[1])), mask))
		}
	}

	return sumMem(mem)
}

func flip(s string, mask string) string {
	result := ""
	//Printfln("value:\t%s\t(decimal %d)", s, binToInt(s))
	//Printfln("mask:\t%s", mask)
	for i, m := range mask {
		if m == 'X' {
			result += string(s[i])
		} else {
			result += string(m)
		}
	}
	//Printfln("result:\t%s\t(decimal %d)", result, binToInt(result))
	return result
}

func strToBin(s string) string {
	return strconv.FormatInt(Int64(s), 2)
}

func binToInt(s string) int64 {
	n, err := strconv.ParseInt(s, 2, 64)
	Check(err)
	return n
}

func pad(s string) string {
	s2 := ""
	l := len(s)
	if l < 36 {
		s2 = strings.Repeat("0", 36-l)
	}
	return s2 + s
}

func sumMem(mem []int64) int64 {
	var sum int64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func Solve2(input []string) int {
	return 0
}

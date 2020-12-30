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

const size = 65485+1

func Solve1(input []string) int64 {
	var mask string
	mem := make([]int64, size)

	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			mask = strings.Split(s, "mask = ")[1]
		} else {
			op := Regex(`mem\[(\d+)\] = (\d+)`, s)
			mem[Int(op[0])] = binToInt(flip(strToBin(op[1]), mask))
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
	return fmt.Sprintf("%036b", Int64(s))
}

func binToInt(s string) int64 {
	n, err := strconv.ParseInt(s, 2, 64)
	Check(err)
	return n
}

func sumMem(mem []int64) int64 {
	var sum int64
	for _, v := range mem {
		sum += v
	}
	return sum
}

func Solve2(input []string) int64 {
	var mask string
	mem := make([]int64, size)
	//floating := make([]int64, 36)


	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			mask = strings.Split(s, "mask = ")[1]
		} else {
			op := Regex(`mem\[(\d+)\] = (\d+)`, s)
			for _, addr := range addresses(strToBin(op[0]), mask) {
				mem[Int(addr)] = Int64(op[1])
			}
		}
	}

	return sumMem(mem)
}

/*
	takes an address in binary, returns all permutations given its mask, eg:
	address: 000000000000000000000000000000101010  (decimal 42)
	mask:    000000000000000000000000000000X1001X
	gives:
	000000000000000000000000000000011010  (decimal 26)
	000000000000000000000000000000011011  (decimal 27)
	000000000000000000000000000000111010  (decimal 58)
	000000000000000000000000000000111011  (decimal 59)
 */
func addresses(addr string, mask string) []string {
	out := make([]string, 0)

	if !strings.Contains(addr, "X") {
		return addr
	}

	return out
}

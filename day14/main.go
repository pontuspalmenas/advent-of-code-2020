package main

import (
	. "aoc"
	. "aoc/types"
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
	for i, m := range mask {
		if m == 'X' {
			result += string(s[i])
		} else {
			result += string(m)
		}
	}
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
	mem := make(map[int64]int64)

	for _, s := range input {
		if strings.HasPrefix(s, "mask") {
			mask = strings.Split(s, "mask = ")[1]
		} else {
			op := Regex(`mem\[(\d+)\] = (\d+)`, s)
			addrs := addresses(strToBin(op[0]), mask)
			for _, addr := range addrs {
				mem[binToInt(addr)] = Int64(op[1])
			}
		}
	}

	var sum int64
	for _, v := range mem {
		sum += v
	}
	return sum
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
	// todo: clean up this... mess.
	var result [36]rune
	floating := make([]int, 0)

	// apply mask
	for i := range addr {
		if mask[i] == 'X' {
			result[i] = 'X'
			floating = append(floating, i)
		} else if mask[i] == '1' || addr[i] == '1' {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}

	tempAddr := ""
	for _, r := range result {
		tempAddr += string(r)
	}
	addrs := NewStringSet()
	addrs.Add(tempAddr)

	for containsWildcard(addrs.ToSlice()) {
		for _, addr := range addrs.ToSlice() {
			if !strings.Contains(addr, "X") {
				continue
			}
			addrs.Remove(addr)
			addrs.Add(strings.Replace(addr, "X", "0", 1))
			addrs.Add(strings.Replace(addr, "X", "1", 1))
		}
	}

	return addrs.ToSlice()
}

func containsWildcard(addrs []string) bool {
	for _, a := range addrs {
		if strings.Contains(a, "X") {
			return true
		}
	}
	return false
}

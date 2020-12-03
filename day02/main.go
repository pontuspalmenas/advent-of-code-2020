package main

import (
	. "aoc"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	input := Input("day02/input.txt")
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

func Solve1(s string) int {
	valid := 0

	regex := *regexp.MustCompile(`(\d+)-(\d+) (.): (\w+)`)
	res := regex.FindAllStringSubmatch(s, -1)
	for i := range res {
		count := strings.Count(res[i][4], res[i][3])
		if count >= ToInt(res[i][1]) && count <= ToInt(res[i][2]) {
			valid++
		}
	}

	return valid
}

func Solve2(s string) int {
	valid := 0

	regex := *regexp.MustCompile(`(\d+)-(\d+) (.): (\w+)`)
	res := regex.FindAllStringSubmatch(s, -1)
	for i := range res {

		str := res[i][4]
		chr := res[i][3]
		pos1 := ToInt(res[i][1])
		pos2 := ToInt(res[i][2])

		pos1ok := str[pos1-1:pos1] == chr
		pos2ok := str[pos2-1:pos2] == chr

		if pos1ok != pos2ok {
			valid++
		}
	}

	return valid
}

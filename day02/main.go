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
		min := ToInt(res[i][1])
		max := ToInt(res[i][2])
		chr := res[i][3]
		str := res[i][4]

		count := strings.Count(str, chr)
		if count >= min && count <= max {
			valid++
		}
	}

	return valid
}

func Solve2(s string) int {
	return 0
}

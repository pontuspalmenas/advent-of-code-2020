package main

import (
	. "aoc"
	"fmt"
	"regexp"
	"strings"
	"time"
	
)

func main() {
	start := time.Now()
	input := SplitByEmptyNewline(Input("day19/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	rules := parseRules(Lines(input[0]))
	messages := input[1]

	regex := fmt.Sprintf("^%s$", expand(rules, "0"))
	r := regexp.MustCompile(regex)

	valid := 0
	for _, s := range Lines(messages) {
		if r.MatchString(s) {
			valid++
		}
	}

	return valid
}

func Solve2(input []string) int {
	rules := parseRules(Lines(input[0]))
	rules[8] = "42 +"
	rules[11] = "(?P<R> 42 (?&R)? 31 )" //
	messages := input[1]

	regex := fmt.Sprintf("^%s$", expand(rules, "0"))
	r := regexp.MustCompile(regex)

	valid := 0
	for _, s := range Lines(messages) {
		if r.MatchString(s) {
			valid++
		}
	}

	return valid
}

func expand(rules map[int]string, val string) string {
	if !IsNumber(val) {
		return val
	}
	set := strings.Fields(rules[Int(val)])
	var list []string
	for _, s := range set {
		list = append(list, expand(rules, s))
	}
	return "(?:" + strings.Join(list, "") + ")"
}

func parseRules(in []string) map[int]string {
	r := make(map[int]string)
	for _, s := range in {
		split := Split(s, ":")
		rule := strings.TrimSpace(split[1])
		rule = strings.ReplaceAll(rule, "\"", "")
		r[Int(split[0])] = rule
	}
	return r
}

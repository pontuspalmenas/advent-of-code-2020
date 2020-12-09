package main

import (
	. "aoc"
	"fmt"
)

type Rules map[string]map[string]int

func main() {
	input := Lines(Input("day07/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

func Solve1(input []string) int {
	rules := parse(input)

	isFound := true
	visited := make(map[string]bool)
	visited["shiny gold"] = true

	// While we still find shiny gold, keep looking for more
	for isFound {
		isFound = false
		for k, v := range rules { // check top-level bags
			if visited[k] { // already seen?
				continue
			}
			for key, _ := range v { // check sub-level bags
				if visited[key] {
					visited[k] = true
					isFound = true
					break
				}
			}
		}
	}

	return len(visited) - 1
}

func Solve2(input []string) int {
	rules := parse(input)
	return count(rules, make(map[string]int), "shiny gold")
}

func count(rules Rules, list map[string]int, name string) int {
	if list[name] > 0 {
		return list[name]
	}
	sum := 0
	for childName, capacity := range rules[name] {
		sum += capacity * (1 + count(rules, list, childName)) // keep looking recursively
	}
	list[name] = sum
	return sum
}

func parse(in []string) Rules {
	bags := Rules{}
	for _, s := range in {
		match := Regex(`^(.*) bags contain (.*)`, s)
		name := match[0]
		bags[name] = map[string]int{}
		for _, content := range RegexAll(`(\d+) (.+?) bag`, match[1]) {
			subName := content[1]
			amount := Int(content[0])
			bags[name][subName] = amount
		}
	}

	return bags
}

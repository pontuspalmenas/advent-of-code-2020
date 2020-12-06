package main

import (
	. "aoc"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	input := input("day04/input.txt")
	fmt.Println(Solve1(input))
	fmt.Println(Solve2(input))
}

func Solve1(input []string) int {
	count := 0
	for _, s := range input {
		if hasAllFields(s) {
			count++
		}
	}

	return count
}

func Solve2(input []string) int {
	count := 0
	for i, passport := range input {
		fmt.Printf("\nPassport #%d\n", i+1)
		validPassport := false

		if hasAllFields(passport) {
			validPassport = true
			for k, v := range parse(passport) {
				if validate(k, v) {
					fmt.Printf("%s valid:\t\t%s\n", k, v)
				} else {
					validPassport = false
					fmt.Printf("%s invalid:\t%s\n", k, v)
				}
			}
		}
		if validPassport {
			fmt.Println("VALID")
			count++
		} else {
			fmt.Println("INVALID")
		}
	}

	return count
}

func hasAllFields(s string) bool {
	for _, k := range []string{"byr","iyr","eyr","hgt","hcl","ecl","pid"} {
		if !strings.Contains(s, k) {
			return false
		}
	}
	return true
}

func inList(s string, list []string) bool {
	for _, k := range list {
		if s == k {
			return true
		}
	}
	return false
}

func parse(s string) map[string]string {
	kvs := make(map[string]string)

	// todo: remove this earlier
	s = strings.ReplaceAll(s, "\n", " ")

	for _, thing := range strings.Split(s, " ") {
		ss := strings.Split(thing, ":")
		if len(ss) != 2 {
			continue
		}

		kvs[ss[0]] = ss[1]
	}

	return kvs
}

func validate(k string, v string) bool {
	f := map[string] func(s string) bool{"byr":byr, "iyr":iyr, "eyr":eyr, "hgt":hgt, "hcl":hcl, "ecl":ecl, "pid":pid, "cid":cid}
	return f[k](v)
}

func byr(s string) bool {
	return year(s, 1920, 2002)
}

func iyr(s string) bool {
	return year(s, 2010, 2020)
}

func eyr(s string) bool {
	return year(s, 2020, 2030)
}

func hgt(s string) bool {
	var size int
	var unit string
	_, err := fmt.Sscanf(s, "%d%s", &size, &unit)
	if err != nil {
		return false
	}
	switch unit {
	case "cm":
		return size >= 150 && size <= 193
	case "in":
		return size >= 59 && size <= 76
	}

	return false
}

func hcl(s string) bool {
	return regexp.MustCompile("^#(?:[0-9a-f]{6})$").MatchString(s)
}

func ecl(s string) bool {
	return inList(s, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
}

func pid(s string) bool {
	return regexp.MustCompile("^(?:[0-9]{9})$").MatchString(s)
}

func cid(s string) bool {
	return true
}

func year(s string, min int, max int) bool {
	n := ToInt(s)
	return n >= min && n <= max
}

func input(f string) []string {
	b, err := ioutil.ReadFile(f)
	Check(err)
	return splitByEmptyNewline(string(b))
}

func splitByEmptyNewline(str string) []string {
	s := strings.ReplaceAll(str, "\r\n", "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(s, -1)
}

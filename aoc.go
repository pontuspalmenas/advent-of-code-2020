package aoc

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

type TestTable struct {
	In  interface{}
	Out interface{}
}

func AssertEq(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual: %v; Expected: %v", actual, expected)
	}
}

func Input(f string) string {
	b, err := ioutil.ReadFile(f)
	Check(err)
	return string(b)
}

func Int(s string) int {
	n, err := strconv.Atoi(s)
	Check(err)
	return n
}

func Int64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	Check(err)
	return n
}

func Lines(s string) []string {
	ls := strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")

	// If last line is empty, skip it
	i := len(ls)
	if ls[i-1] == "" {
		ls = ls[:i-1]
	}
	return ls
}

func Ints(s string) []int {
	var ints []int
	re := regexp.MustCompile(`-?\d+`)
	for _, e := range re.FindAllString(s, -1) {
		ints = append(ints, Int(e))
	}

	return ints
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Sscanf(s string, format string, a ...interface{}) {
	_, err := fmt.Sscanf(s, format, a...)
	Check(err)
}

func SplitByEmptyNewline(str string) []string {
	s := strings.ReplaceAll(str, "\r\n", "\n")
	return regexp.MustCompile(`\n\s*\n`).Split(s, -1)
}

func Printfln(format string, a ...interface{}) {
	fmt.Printf("%s\n", fmt.Sprintf(format, a...))
}

// Regex returns all matches, skipping the full match
func Regex(r string, s string) []string {
	var out []string
	submatch := regexp.MustCompile(r).FindStringSubmatch(s)
	for i := 1; i < len(submatch); i++ {
		out = append(out, submatch[i])
	}
	return out
}

// RegexAll returns all submatches, skipping the full match
func RegexAll(r string, s string) [][]string {
	var out [][]string
	submatches := regexp.MustCompile(r).FindAllStringSubmatch(s, -1)
	for _, submatch := range submatches {
		var parts []string
		for i := 1; i < len(submatch); i++ {
			parts = append(parts, submatch[i])
		}
		out = append(out, parts)
	}

	return out
}

type Point struct {
	X int
	Y int
}

func Manhattan(p1 Point, p2 Point) int {
	return int(math.Abs(float64(p1.X - p2.X))) + int(math.Abs(float64(p1.Y - p2.Y)))
}

func MaxInts(ints []int) int {
	max := ints[0]
	for _, n := range ints {
		if n > max {
			max = n
		}
	}
	return max
}

func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

func SplitByComma(s string) (out []string) {
	for _, n := range strings.Split(s, ",") {
		out = append(out, strings.TrimSpace(n))
	}

	return out
}

func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func Make2DRuneSlice(n, m int) [][]rune {
	a := make([][]rune, n)
	for i := range a {
		a[i] = make([]rune, m)
	}
	return a
}

func Copy2DRuneSlice(matrix [][]rune) [][]rune {
	duplicate := make([][]rune, len(matrix))
	for i := range matrix {
		duplicate[i] = make([]rune, len(matrix[i]))
		copy(duplicate[i], matrix[i])
	}
	return duplicate
}

func StringFromIntSlice(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ", "), "[]")
}

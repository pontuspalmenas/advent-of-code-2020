package aoc

import (
	"fmt"
	"io/ioutil"
	"os"
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

func ChrAt(s string, i int) string {
	return s[i:i+1]
}

func Sscanf(s string, format string, a ...interface{}) {
	_, err := fmt.Sscanf(s, format, a...)
	Check(err)
}

func Fscanf(path string, format string, a ...interface{}) {
	file, err := os.Open(path)
	Check(err)
	defer file.Close()
	_, err = fmt.Fscanf(file, format, a...)
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

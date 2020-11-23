package aoc

import (
	"bufio"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
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
	if err != nil {
		panic(err)
	}
	return string(b)
}

func ToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func InputLines(path string) []string {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var out []string
	for sc.Scan() {
		out = append(out, sc.Text())
	}

	return out
}

func InputIntLines(path string) []int {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanLines)

	var out []int
	for sc.Scan() {
		i, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		out = append(out, i)
	}

	return out
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

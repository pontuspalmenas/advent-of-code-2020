package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	input := Lines(Input("day21/input.txt"))
	p1 := Solve1(input)
	p2 := Solve2(input)
	fmt.Println(time.Since(start))
	fmt.Println("p1:", p1)
	fmt.Println("p2:", p2)
}

type dish struct {
	ingredients *StringSet
	allergens *StringSet
}

func Solve1(input []string) int {
	dishes := parse(input)

	// Remove unsafe ingredients from food
	for _, ingr := range candidates(dishes) {
		for _, dish := range dishes {
			dish.ingredients.Remove(ingr)
		}
	}

	// Count how many are left
	count := 0
	for _, dish := range dishes {
		count += dish.ingredients.Size()
	}

	return count
}

func Solve2(input []string) string {
	dishes := parse(input)
	final := candidates(dishes)

	var allergs []string
	for k, _ := range final {
		allergs = append(allergs, k)
	}
	sort.Strings(allergs)
	var result string
	for _, s := range allergs {
		result += final[s] + ","
	}
	return strings.TrimRight(result, ",")
}

func candidates(dishes []dish) map[string]string {
	candidates := make(map[string]*StringSet)
	for _, dish := range dishes {
		for _, allergen := range dish.allergens.ToSlice() {
			if candidates[allergen] == nil {
				candidates[allergen] = NewStringSet()
			}
			candidates[allergen] = candidates[allergen].Union(dish.ingredients)
		}
	}

	// Now reduce to only possible candidates
	for _, dish := range dishes {
		for _, allergen := range dish.allergens.ToSlice() {
			candidates[allergen] = candidates[allergen].Intersection(dish.ingredients)
		}
	}

	mapping := make(map[string]string)
	size := len(candidates)
	for len(mapping) < size {
		for allerg, ingrs := range candidates {
			if ingrs.Size() == 1 {
				found := ingrs.ToSlice()[0]
				mapping[allerg] = found
				delete(candidates, allerg)
				for _, v := range candidates {
					v.Remove(found)
				}
			}
		}
	}
	return mapping
}

func parse(input []string) []dish {
	var dishes []dish
	for _, s := range input {
		// todo: replace this mess with regex
		split := Split(s, " (")
		d := dish{ingredients: NewStringSet(), allergens: NewStringSet()}
		for _, ingr := range strings.Fields(split[0]) {
			d.ingredients.Add(ingr)
		}
		allergens := split[1][len("(contains"):len(split[1])-1]
		for _, allergen := range SplitByComma(allergens) {
			d.allergens.Add(allergen)
		}

		dishes = append(dishes, d)
	}
	return dishes
}

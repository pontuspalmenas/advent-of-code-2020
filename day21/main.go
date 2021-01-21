package main

import (
	. "aoc"
	. "aoc/types"
	"fmt"
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
	for _, ingrs := range candidates(dishes) {
		for _, ingr := range ingrs.ToSlice() {
			for _, dish := range dishes {
				dish.ingredients.Remove(ingr)
			}
		}
	}

	// Count how many are left
	count := 0
	for _, dish := range dishes {
		count += dish.ingredients.Size()
	}

	return count
}

func Solve2(input []string) int {
	dishes := parse(input)
	candidates := candidates(dishes)

	final := make(map[string]string)
	changed := true
	for changed {
		changed = false

	}

	return 0
}

func candidates(dishes []dish) map[string]*StringSet {
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
	return candidates
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

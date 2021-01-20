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

	possible := make(map[string]*StringSet)
	for _, dish := range dishes {
		for _, ingr := range dish.ingredients.ToSlice() {
			if possible[ingr] == nil {
				possible[ingr] = NewStringSet()
			}
			possible[ingr] = possible[ingr].Union(dish.allergens)
		}
	}

	// ok, this time, reduce.
	for k, v := range possible {
		for _, dish := range dishes {
			if dish.ingredients.Contains(k) {
				for _, a := range possible[k].ToSlice() {
					if !dish.allergens.Contains(a) {
						v.Remove(a)
					}
				}
			}
		}
	}

	// för varje ingrediens
		// gå igenom matlistan
			// om ingrediensen finns där
				// om inte allergenen finns där
					// plocka bort från possible

	for k, v := range possible {
		Printfln("%s %v", k, v)
	}

	safe := NewStringSet()
	return countSafeOccurrences(safe, dishes)
}

func Solve2(input []string) int {
	return 0
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

func countSafeOccurrences(safe *StringSet, dishes []dish) int {
	count := 0
	for _, s := range safe.ToSlice() {
		for _, dish := range dishes {
			for _, ingredient := range dish.ingredients.ToSlice() {
				if s == ingredient {
					count++
				}
			}
		}
	}
	return count
}

func uniqueIngredients(dishes []dish) *StringSet {
	set := NewStringSet()
	for _, d := range dishes {
		for _, s := range d.ingredients.ToSlice() {
			set.Add(s)
		}
	}
	return set
}
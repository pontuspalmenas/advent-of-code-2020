package main

import (
	. "aoc"
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
	ingredients []string
	allergens []string
}

func Solve1(input []string) int {
	dishes := parse(input)

	// what ingredients may contain allergen?
	allergenToIngredients := map[string][][]string{}
	for _, dish := range dishes {
		for _, allergen := range dish.allergens {
			if _, ok := allergenToIngredients[allergen]; !ok {
				allergenToIngredients[allergen] = [][]string{dish.ingredients}
			} else {
				allergenToIngredients[allergen] = append(allergenToIngredients[allergen], dish.ingredients)
			}
		}
	}

	// find candidates for safe ingredients
	candidates := NewStringSet()
	for _, list := range allergenToIngredients {
		for _, ingredients := range list {
			for _, ingredient := range ingredients {
				if !ingredientIsInAll(list, ingredient) {
					candidates.Add(ingredient)
					println("candidate:", ingredient)
				}
			}
		}
	}

	// reduce candidates to one ingredient per allergen
	safe := NewStringSet()



	// count safe in ingredients
	count := 0
	//safe = ToStringSet([]string{"kfcds","nhms","sbzzf","trh"})
	for _, s := range safe.ToSlice() {
		for _, dish := range dishes {
			for _, ingredient := range dish.ingredients {
				if s == ingredient {
					count++
				}
			}
		}
	}

	return count
}

func Solve2(input []string) int {
	return 0
}

func parse(input []string) []dish {
	var dishes []dish
	for _, s := range input {
		// todo: replace this mess with regex
		split := Split(s, " (")
		d := dish{}
		for _, ingr := range strings.Fields(split[0]) {
			d.ingredients = append(d.ingredients, ingr)
		}
		allergens := split[1][len("(contains"):len(split[1])-1]
		for _, allergen := range SplitByComma(allergens) {
			d.allergens = append(d.allergens, allergen)
		}

		dishes = append(dishes, d)
	}
	return dishes
}

func ingredientIsInAll(list [][]string, ingredient string) bool {
	for _, ingredients := range list {
		found := false
		for _, ingr := range ingredients {
			if ingr == ingredient {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func allIngredients(dishes []dish) (all [][]string) {
	for _, d := range dishes {
		all = append(all, d.ingredients)
	}
	return all
}

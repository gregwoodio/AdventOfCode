package day21

import (
	"sort"
	"strings"
)

type food struct {
	ingredients []string
	allergens   []string
}

func solvePartOne(input []string) int {

	foods := parseInput(input)

	allergenicIngredients := make(map[string]map[string]bool)
	allIngredientUsages := make(map[string]int)

	for _, food := range foods {
		for _, i := range food.ingredients {
			allIngredientUsages[i]++
		}

		for _, a := range food.allergens {
			if _, ok := allergenicIngredients[a]; !ok {
				// new allergen, add all ingredients
				allergenicIngredients[a] = make(map[string]bool)
				for _, i := range food.ingredients {
					allergenicIngredients[a][i] = true
				}

				continue
			}

			// existing allergen, trim list
			for ing := range allergenicIngredients[a] {
				found := false
				for _, i := range food.ingredients {
					if ing == i {
						found = true
						break
					}
				}

				if !found {
					delete(allergenicIngredients[a], ing)
				}
			}
		}
	}

	for _, i := range allergenicIngredients {
		for ing := range i {
			delete(allIngredientUsages, ing)
		}
	}

	sum := 0
	for _, v := range allIngredientUsages {
		sum += v
	}

	return sum
}

func solvePartTwo(input []string) string {
	foods := parseInput(input)

	allergenicIngredients := make(map[string]map[string]bool)

	for _, food := range foods {
		for _, a := range food.allergens {
			if _, ok := allergenicIngredients[a]; !ok {
				// new allergen, add all ingredients
				allergenicIngredients[a] = make(map[string]bool)
				for _, i := range food.ingredients {
					allergenicIngredients[a][i] = true
				}

				continue
			}

			// existing allergen, trim list
			for ing := range allergenicIngredients[a] {
				found := false
				for _, i := range food.ingredients {
					if ing == i {
						found = true
						break
					}
				}

				if !found {
					delete(allergenicIngredients[a], ing)
				}
			}
		}
	}

	// find allergen with only one possible ingredient, remove from all other possible allergens
	final := make(map[string]string)
	for {

		for k, v := range allergenicIngredients {
			if len(v) == 1 {
				var ingredient string
				for i := range v {
					ingredient = i
				}
				final[k] = ingredient

				// remove ingredient from other allergens
				for allergen := range allergenicIngredients {
					if allergen == k {
						continue
					}

					delete(allergenicIngredients[allergen], ingredient)
				}

				delete(allergenicIngredients, k)
			}
		}

		if len(allergenicIngredients) == 0 {
			break
		}
	}

	allergens := []string{}
	for k := range final {
		allergens = append(allergens, k)
	}

	sort.Strings(allergens)
	ingredients := []string{}
	for _, a := range allergens {
		ingredients = append(ingredients, final[a])
	}

	return strings.Join(ingredients, ",")
}

func parseInput(input []string) []food {
	foods := []food{}

	for _, line := range input {
		f := food{}
		ingredients := strings.Split(line, " ")

		for _, ingredient := range ingredients {
			if ingredient[0] == '(' {
				break
			}

			f.ingredients = append(f.ingredients, ingredient)
		}

		idx := strings.Index(line, "(contains ")
		if idx != -1 {
			idx += len("(contains ")

			allergens := strings.Split(line[idx:], " ")
			for _, allergen := range allergens {
				if allergen[len(allergen)-1] == ')' {
					f.allergens = append(f.allergens, allergen[:len(allergen)-1])
					break
				}

				// still need to trim off the ',' character
				f.allergens = append(f.allergens, allergen[:len(allergen)-1])
			}
		}

		foods = append(foods, f)
	}

	return foods
}

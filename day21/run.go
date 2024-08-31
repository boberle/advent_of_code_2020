package day21

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Allergen string
type Ingredient string

type Group []Ingredient
type AllergenGroups map[Allergen][]*Group

func Run(infile string) {

	fh, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	groups, allergenGroups := parseFile(fh)
	ingredient2allergen := findAllergens(allergenGroups)
	nonAllergenIngredient := findHowManyTimesNonAllergenIngredientAppear(groups, ingredient2allergen)
	fmt.Printf("How many times non allergenic ingredient appears (part 1): %d\n", nonAllergenIngredient)

	canonicalDangerousIngredientList := getCanonicalDangerousIngredientList(ingredient2allergen)
	fmt.Printf("Canonical list (part 2): %s\n", canonicalDangerousIngredientList)

}

func parseFile(reader io.Reader) ([]Group, AllergenGroups) {
	groups := []Group{}
	allergenGroups := AllergenGroups{}

	scanner := bufio.NewScanner(reader)

	pattern, err := regexp.Compile("^((?:\\w+ )*\\w+) \\(contains ((?:\\w+, )*\\w+)\\)$")
	if err != nil {
		panic(err)
	}
	for scanner.Scan() {
		line := scanner.Text()
		data := pattern.FindStringSubmatch(line)
		if data == nil {
			log.Fatalf("don't understand line: %s", line)
		}

		ingredients := Group{}
		for _, ingredientString := range strings.Split(data[1], " ") {
			ingredients = append(ingredients, Ingredient(ingredientString))
		}
		groups = append(groups, ingredients)

		allergensString := strings.Split(data[2], ", ")
		for _, allergenString := range allergensString {
			allergen := Allergen(allergenString)
			if _, found := allergenGroups[allergen]; !found {
				allergenGroups[allergen] = make([]*Group, 0)
			}
			allergenGroups[allergen] = append(allergenGroups[allergen], &ingredients)
		}
	}

	return groups, allergenGroups
}

func findCommonIngredient(groups []*Group, ingredient2allergen *map[Ingredient]Allergen) (Ingredient, bool) {
	totals := map[Ingredient]int{}
	for _, group := range groups {
		for _, ingredient := range *group {
			if _, found := (*ingredient2allergen)[ingredient]; !found {
				totals[ingredient]++
			}
		}
	}

	var commonIngredient Ingredient
	found := false
	for ingredient, total := range totals {
		if total == len(groups) {
			if !found {
				commonIngredient = ingredient
				found = true
			} else {
				return "", false
			}
		}
	}
	return commonIngredient, true
}

func findAllergens(allergenGroups AllergenGroups) map[Ingredient]Allergen {
	rv := map[Ingredient]Allergen{}
	foundAllergens := map[Allergen]struct{}{}
	for {
		foundIngredient := false
		for allergen, groups := range allergenGroups {
			if _, found := foundAllergens[allergen]; found {
				continue
			}
			if commonIngredient, found := findCommonIngredient(groups, &rv); found {
				rv[commonIngredient] = allergen
				foundAllergens[allergen] = struct{}{}
				foundIngredient = true
			}
		}
		if !foundIngredient {
			break
		}
	}
	return rv
}

func findHowManyTimesNonAllergenIngredientAppear(groups []Group, ingredient2allergen map[Ingredient]Allergen) int {
	total := 0
	for _, group := range groups {
		for _, ingredient := range group {
			if _, found := ingredient2allergen[ingredient]; !found {
				total++
			}
		}
	}
	return total
}

func getCanonicalDangerousIngredientList(ingredient2allergen map[Ingredient]Allergen) string {
	allergen2ingredient := map[Allergen]Ingredient{}
	for ingredient, allergen := range ingredient2allergen {
		allergen2ingredient[allergen] = ingredient
	}

	allergens := maps.Values(ingredient2allergen)
	sort.Slice(allergens, func(i, j int) bool {
		return allergens[i] < allergens[j]
	})

	rv := ""
	for _, allergen := range allergens {
		if len(rv) > 0 {
			rv += ","
		}
		rv += string(allergen2ingredient[allergen])
	}
	return rv
}

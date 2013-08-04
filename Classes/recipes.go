package main

import "github.com/pmylund/sortutil"
import "github.com/hahnicity/filterutils"
import "fmt"

// Recipe struct //
type Recipe struct {
    type_ string
    ingredients []string    
}

// Recipe Book //
type RecipeBook struct {
    recipes []Recipe    
}

func (r *RecipeBook) SortByType(type_ string) {
    sortutil.AscByField(r.recipes, "type_")
}

func (r *RecipeBook) GetByIngredient(ingredient string) []Recipe {
    desired_recipes := make([]Recipe, 0)
    for _, j := range r.recipes {
        sortutil.Asc(j.ingredients)
        if filterutils.StringInSortedSlice(j.ingredients, ingredient) {
            desired_recipes = append(desired_recipes, j)
        }
    }
    return desired_recipes
}

func (r *RecipeBook) AddRecipe(recipe Recipe) {
    // More a convenience method than anything else
    r.recipes = append(r.recipes, recipe)    
}

// main
func main() {
    rb := new(RecipeBook)    
    rb.AddRecipe(Recipe{"eggs n cheese", []string{"eggs", "cheese"}})
    rb.AddRecipe(Recipe{"eggs n ham", []string{"eggs", "ham"}})
    rb.AddRecipe(Recipe{"steak n potatoes", []string{"steak", "potatoes"}})
    fmt.Println(rb.GetByIngredient("potatoes"))
}

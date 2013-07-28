package main

import "fmt"

// Product Class //
type Product struct {
    id, price, quantity int    
}

// Inventory Class //
type Inventory struct {
    Products[]Product // Capitalized variables are visible outside of pkg
    // products[]Product //Lowercase variables are not visible outside of pkg
}

func (inventory Inventory) GetSumOfValues() int {
    productSum := 0
    for _, value := range inventory.Products {
        productSum = productSum + value.price * value.quantity
    }
    return productSum
}

/*
Thanks to http://www.golang-book.com/9 I was able to figure out 
(inventory Inventory) performs much like a static function and 
(inventory *Inventory) performs like an instance function, and is able
to modify variables within the inventory instance. Cool
*/
func (inventory *Inventory) AddProduct(product Product) []Product {
    size := len(inventory.Products)
    if size + 1 > cap(inventory.Products) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]Product, (size+1*2))
        // The copy function is predeclared and works for any products type.
        copy(newSlice, inventory.Products)
        inventory.Products = newSlice
    }
    inventory.Products = inventory.Products[0:size+1]
    inventory.Products[size] = product
    return nil
}

// Main
func main() {
    inventory := new(Inventory)
    alienArtifacts := Product{1, 20, 2}
    inventory.AddProduct(alienArtifacts)
    fmt.Println("SUM OF PRODUCTS: $", inventory.GetSumOfValues())
}

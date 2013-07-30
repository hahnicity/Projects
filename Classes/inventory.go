package main

import "fmt"

// Product Class //
type Product struct {
    id, price, quantity int    
}

// Inventory Class //
type Inventory struct {
    Products []Product // Capitalized variables are visible outside of pkg
    // products[]Product //Lowercase variables are not visible outside of pkg
}

func (inventory Inventory) GetSumOfValues() int {
    productSum := 0
    for _, value := range inventory.Products {
        productSum = productSum + value.price * value.quantity
    }
    return productSum
}

// Inventory Constructor //
func makeInventory() *Inventory {
    inventory := new(Inventory)
    inventory.Products = make([]Product, 0)
    return inventory   
}

// Main
func main() {
    inventory := makeInventory()
    alienArtifacts := Product{1, 20, 2}
    inventory.Products = append(inventory.Products, alienArtifacts)
    indianaJones := Product{2, 2000, 1}
    inventory.Products = append(inventory.Products, indianaJones)
    rubixCubes := Product{3, 1, 10000}
    inventory.Products = append(inventory.Products, rubixCubes)
    fmt.Println("SUM OF PRODUCTS: $", inventory.GetSumOfValues())
}

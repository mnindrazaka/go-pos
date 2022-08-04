package main

import (
	"fmt"
	"strings"
)

func main() {
	shop := Shop{products: []Product{
		{name: "Pencil", price: 3000},
		{name: "Book", price: 3500},
		{name: "Eraser", price: 2000},
	}}

	transaction := Transaction{carts: []Cart{}}

	repeat := true

	for {
		fmt.Println("========================================")
		fmt.Println("Welcome to Our Shop")
		fmt.Println("Product List :")

		shop.PrintProducts()

		fmt.Println("========================================")

		fmt.Print("Enter the product option : ")
		var productOption int
		fmt.Scanf("%d", &productOption)

		product := shop.GetProduct(productOption - 1)

		fmt.Print("How much do you want to buy ? ")
		var amount int
		fmt.Scanf("%d", &amount)

		transaction.AddCart(Cart{product, amount})

		fmt.Print("Do you want to buy again (y/n) ? ")
		var answer string
		fmt.Scanf("%s", &answer)

		repeat = strings.ToLower(answer) == "y"

		if !repeat {
			break
		}
	}

	fmt.Println("Total amount is Rp.", transaction.GetTotal())

	fmt.Print("How much the money ? Rp.")
	var money int
	fmt.Scanf("%d", &money)

	change := money - transaction.GetTotal()

	if money < transaction.GetTotal() {
		fmt.Println("You need Rp.", change*-1, "more money")
	} else {
		fmt.Println("Your change is Rp.", change)
	}
}

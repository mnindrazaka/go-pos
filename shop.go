package main

import "fmt"

type Shop struct {
	products []Product
}

type Shoper interface {
	PrintProducts()
	GetProduct()
}

func (shop Shop) PrintProducts() {
	for i := 0; i < len(shop.products); i++ {
		fmt.Println(i+1, ". ", shop.products[i].name, " \t: Rp.", shop.products[i].price)
	}
}

func (shop Shop) GetProduct(index int) Product {
	return shop.products[index]
}

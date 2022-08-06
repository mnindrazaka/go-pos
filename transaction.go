package main

type Transaction struct {
	carts []Cart
}

func (transaction *Transaction) AddCart(cart Cart) {
	transaction.carts = append(transaction.carts, cart)
}

func (transaction Transaction) GetTotal() int {
	total := 0
	for _, cart := range transaction.carts {
		total += cart.product.price * cart.amount
	}
	return total
}

package main

import "fmt"

type Product struct {
	name     string
	price    float64
	quantity int
}

func ProductsSum(products []Product, total *float64) float64 {
	sum := 0.0

	for _, product := range products {
		sum += product.price * float64(product.quantity)
	}

	*total += sum
	fmt.Println("product sum", *total)

	return sum
}

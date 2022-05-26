package main

type Product struct {
	name     string
	price    float64
	quantity int
}

func newProduct(name string, price float64) Product {
	return Product{
		name:  name,
		price: price,
	}
}

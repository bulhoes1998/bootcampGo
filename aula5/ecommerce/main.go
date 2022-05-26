package main

import "fmt"

func main() {
	u1 := User{"Lucas", "Bulhoes", "lucas.bulhoes@mercadolivre.com", nil}

	p1 := newProduct("iphone X", 10_000.00)

	u1.addProduct(p1, 2)

	fmt.Println(u1)
	fmt.Println(p1)
	u1.delProducts()
	fmt.Println(u1)
	fmt.Println(p1)
}

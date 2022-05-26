package main

type User struct {
	firstName string
	surName   string
	email     string
	products  []Product
}

func (u *User) addProduct(p Product, qnt int) {
	p.quantity = qnt
	u.products = append(u.products, p)
}

func (u *User) delProducts() {
	u.products = nil
}

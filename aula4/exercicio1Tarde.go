package main

import (
	"fmt"
	"os"
)

type Produto struct {
	id         int
	preco      float64
	quantidade int
}

func (p Produto) Descricao() string {
	return fmt.Sprintf("%d;%.2f;%d\ne", p.id, p.preco, p.quantidade)
}

func (p Produto) insertCompra(path string) error {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.Write([]byte(p.Descricao()))

	return nil

}

func main() {
	path := "./compras.csv"

	compra := Produto{2, 2.69, 100}

	compra.insertCompra(path)
}

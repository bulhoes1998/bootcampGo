package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ID         = "id"
	preco      = "preco"
	quantidade = "quantidade"
)

func imprimirDetalhes(compras [][]string) {
	fmt.Printf("%-10v %10v %5v\n", ID, preco, quantidade)
	var soma float32
	for _, compra := range compras {
		fmt.Printf("%-10v", compra[0])
		fmt.Printf("%10v", compra[1])
		fmt.Printf("%9v\n", compra[2])
		p, _ := strconv.ParseFloat(compra[1], 8)
		qtd, _ := strconv.ParseFloat(compra[2], 8)
		soma += float32(p) * float32(qtd)
	}

	fmt.Println(soma)
}

func main() {
	path := "./compras.csv"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	var split [][]string
	for scanner.Scan() {
		split = append(split, strings.Split(scanner.Text(), ";"))
	}

	imprimirDetalhes(split)
}

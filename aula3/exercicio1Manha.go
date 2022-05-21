package main

import (
	"fmt"
	"os"
	"strconv"
)

func imposto(salario float64) float64 {
	if salario < 150_000.0 {
		return salario * 0.17
	} else {
		return salario * 0.27
	}
}

func main() {
	salario, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Println(imposto(salario))
}

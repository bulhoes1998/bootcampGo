package main

import (
	"errors"
	"fmt"
)

func impostNeed(salary int) (string, error) {
	if salary < 15_000 {
		return "", errors.New("erro: o salario digitado nao esta dentro do valor minimo")
	}

	return "Necessario pagar imposto", nil
}

func main() {
	var salary = 16_000
	resp, err := impostNeed(salary)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}

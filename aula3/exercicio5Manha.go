package main

import (
	"errors"
	"fmt"
)

const (
	cat       = "cat"
	dog       = "dog"
	hamster   = "hamster"
	tarantula = "tarantula"
)

func catFunc(qnt int) int {
	return qnt * 5_000
}

func dogFunc(qnt int) int {
	return qnt * 10_000
}

func hasmsterFunc(qnt int) int {
	return qnt * 250
}

func tarantulaFunc(qnt int) int {
	return qnt * 150
}

func Animal(animalCategory string) (func(qnt int) int, error) {
	switch animalCategory {
	case cat:
		return catFunc, nil
	case dog:
		return dogFunc, nil
	case hamster:
		return hasmsterFunc, nil
	case tarantula:
		return tarantulaFunc, nil
	default:
		return nil, errors.New("Categoria nao cadastrada.")
	}
}

func main() {
	testAnimalFunc, _ := Animal(cat)
	foodAmount := testAnimalFunc(10)
	fmt.Println(foodAmount)

	_, err := Animal("turtle")
	if err != nil {
		fmt.Println(err)
	}
}

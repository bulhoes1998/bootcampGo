package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	maximum = "maximum"
	average = "average"
)

func minFunc(notas ...float64) float64 {
	minValue := notas[0]
	for _, nota := range notas {
		if nota < minValue {
			minValue = nota
		}
	}

	return minValue
}

func maFunc(notas ...float64) float64 {
	maxValue := notas[0]
	for _, nota := range notas {
		if nota > maxValue {
			maxValue = nota
		}
	}

	return maxValue
}

func averFunc(notas ...float64) float64 {
	soma := 0.0
	for _, nota := range notas {
		soma += nota
	}

	return soma / float64(len(notas))
}

func operation(functionChoose string) (func(valores ...float64) float64, error) {
	switch functionChoose {
	case minimum:
		return minFunc, nil
	case maximum:
		return maFunc, nil
	case average:
		return averFunc, nil
	default:
		return averFunc, errors.New("Operacao invalida.")
	}
}

func main() {
	minhaFunc, _ := operation(minimum)
	averageFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minValue := minhaFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println(minValue, averageValue, maxValue)
}

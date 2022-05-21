package main

import "fmt"

const (
	A = 3_000
	B = 1_500
	C = 1_000
)

func calcularSalario(categoria int, horas int) float64 {

	switch categoria {
	case A:
		if horas > 160 {
			return float64(categoria) * float64(horas) * 1.5
		} else {
			return float64(categoria) * float64(horas)
		}
	case B:
		if horas > 160 {
			return float64(categoria) * float64(horas) * 1.2
		} else {
			return float64(categoria) * float64(horas)
		}
	case C:
		return float64(categoria) * float64(horas)
	}

	return 0.0

}

func main() {
	fmt.Println(calcularSalario(10, 120))
}

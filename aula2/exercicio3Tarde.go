package main

import (
	"fmt"
	"os"
	"strconv"
)

// Função numeroParaMes converte um numero para a string do respectivo mes
func numeroParaMes(numero int) string {
	meses := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}

	return meses[numero]
}

func main() {
	// recebe um numero como argumento no terminal
	numero, _ := strconv.Atoi(os.Args[1])
	fmt.Println(numeroParaMes(numero))
}

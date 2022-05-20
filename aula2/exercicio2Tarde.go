package main

import "fmt"

// Define a struct cliente, para armazenar informações de um cliente
type cliente struct {
	idade          int
	trabalha       bool
	tempoNoEmprego int
	salario        float64
}

// Função podeEmprestar retorna um bool que define se o
// cliente possui direito a emprestimo
func podeEmprestar(pessoa cliente) bool {
	switch {
	case pessoa.idade <= 22:
		fmt.Println("cliente não possui idade suficiente.")
		return false
	case !pessoa.trabalha:
		fmt.Println("cliente está desempregado.")
		return false
	case pessoa.tempoNoEmprego <= 12:
		fmt.Println("cliente tem pouco tempo empregado.")
		return false
	default:
		fmt.Println("cliente apto para emprestimo.")
		if pessoa.salario > 100_000.0 {
			fmt.Println("cliente tem direito a isenção de juros.")
		} else {
			fmt.Println("cliente não tem direito a isenção de juros.")
		}
		return true
	}
}

func main() {
	// Instrancia um cliente e realiza a chamada na função podeEmprestar
	pessoa := cliente{23, true, 13, 112000.0}
	podeEmprestar(pessoa)
}

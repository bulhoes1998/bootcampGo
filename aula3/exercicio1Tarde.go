package main

import "fmt"

type Aluno struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao string
}

func (a Aluno) detalhe() {
	fmt.Printf("Aluno: %v\n", a)
}

func main() {
	aluno1 := Aluno{"Lucas", "Bulhoes", "1365790002", "13/05/2017"}

	aluno1.detalhe()
}

package main

import "fmt"

type SalaryError struct {
	Message string
}

func (se SalaryError) Error() string {
	return se.Message
}

func impostNeed(salary int) (string, error) {
	if salary < 15_000 {
		return "", SalaryError{
			Message: "erro: o salario digitado nao esta dentro do valor minimo",
		}
	}

	return "Necessario pagamento de imposto", nil
}

func main() {
	var salary = 2000
	resp, err := impostNeed(salary)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

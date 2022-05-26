package main

import "fmt"

func impostNeed(salary int) (string, error) {
	if salary < 15_000 {
		return "", fmt.Errorf("erro: o minimo tributavel e 15_000 e o salario informado e: %v", salary)
	}

	return "Necessario pagamento de imposto", nil
}

func main() {
	var salary = 5_000
	resp, err := impostNeed(salary)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}

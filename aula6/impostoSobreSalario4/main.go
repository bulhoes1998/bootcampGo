package main

import (
	"errors"
	"fmt"
)

func monthSalary(hours int, price float64) (float64, error) {
	if hours < 80 {
		return 0, errors.New("error: o trabalhador nao pode ter trabalhado menos de 80h por mes")
	}

	salary := float64(hours) * price

	return salary, nil
}

func impostNeed(salary float64) (string, error) {
	if salary < 20_000 {
		return "", fmt.Errorf("error: salario nao atingiu valor minimo de 20_000, salary: %v", salary)
	}

	impost := salary * 0.1
	return fmt.Sprintf("Necessario pagamento de 10%% de imposto, equivalente a: %.2f", impost), nil
}

func main() {
	hours := 81
	price := 100.76

	resp, err := monthSalary(hours, price)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	resp2, err := impostNeed(resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp2)

}

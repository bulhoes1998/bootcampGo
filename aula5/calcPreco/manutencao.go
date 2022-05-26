package main

import "fmt"

type Maintenance struct {
	name  string
	price float64
}

func MaintenanceSum(m []Maintenance, total *float64) float64 {
	sum := 0.0

	for _, maintenance := range m {
		sum += maintenance.price
	}

	*total += sum
	fmt.Println("maintenance sum", *total)

	return sum
}

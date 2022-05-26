package main

import (
	"fmt"
	"math"
)

type Service struct {
	name        string
	price       float64
	workMinutes int
}

func ServicesSum(services []Service, total *float64) float64 {
	sum := 0.0

	for _, service := range services {
		sum += service.price * math.Max(30, float64(service.workMinutes))
	}

	*total += sum
	fmt.Println("service sum", *total)

	return sum
}

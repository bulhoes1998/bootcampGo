package main

import "time"

func main() {
	p1 := Product{}
	s1 := Service{}
	m1 := Maintenance{}

	products := []Product{p1}
	services := []Service{s1}
	maintenances := []Maintenance{m1}

	total := 0.0

	go ProductsSum(products, &total)
	go ServicesSum(services, &total)
	go MaintenanceSum(maintenances, &total)

	time.Sleep(5000 * time.Microsecond)
}

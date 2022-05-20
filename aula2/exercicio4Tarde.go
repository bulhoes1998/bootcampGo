package main

import "fmt"

func main() {
	employees := map[string]int{
		"Benjamin": 20,
		"Manuel":   26,
		"Brenda":   19,
		"Dario":    44,
		"Pedro":    30,
	}

	contador := 0
	for _, funcionario := range employees {
		if funcionario > 21 {
			contador++
		}
	}

	fmt.Println(contador)

	employees["Federico"] = 25

	delete(employees, "Pedro")
}

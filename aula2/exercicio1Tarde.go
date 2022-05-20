package main

import "fmt"

func letrasPalavra(palavra string) {

	// Conta o número de letras na palavra e imprime seu valor
	tamanho := len(palavra)
	fmt.Println(tamanho)

	// Itera sobre a palavra e imprime as letras separadamente
	for _, letra := range palavra {
		fmt.Printf("%q ", letra)
	}
	fmt.Println()

}

func main() {
	palavra := "palavra"
	letrasPalavra(palavra)
}

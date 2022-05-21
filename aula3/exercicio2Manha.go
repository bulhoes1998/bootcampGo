package main

import (
	"errors"
	"fmt"
)

func mediaDoAluno(notas ...float64) (float64, error) {
	soma := 0.0
	for _, nota := range notas {
		if nota < 0 {
			return -1.0, errors.New("Não pode ter nota negativa.")
		} else {
			soma += nota
		}
	}

	media := soma / float64(len(notas))

	return media, nil
}

func main() {
	media, err := mediaDoAluno(10.0, 8.0, 3.0, 4.0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(media)
	}
}

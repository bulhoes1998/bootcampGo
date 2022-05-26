package main

import (
	"math/rand"
)

func main() {
	variavel1 := rand.Perm(100)
	variavel2 := rand.Perm(1000)
	variavel3 := rand.Perm(10000)

	InsertionSort(variavel1)
	InsertionSort(variavel2)
	InsertionSort(variavel3)
}

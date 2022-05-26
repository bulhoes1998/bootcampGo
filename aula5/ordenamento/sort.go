package main

func InsertionSort(numbers []int) {
	var n = len(numbers)

	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if numbers[j-1] > numbers[j] {
				numbers[j-1], numbers[j] = numbers[j], numbers[j-1]
			}
			j = j - 1
		}
	}
}

func 

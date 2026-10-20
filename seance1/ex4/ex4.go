package main

import "fmt"

func triSelection(nombres []int) {
	for i := 0; i < len(nombres); i++ {
		indexMin := i

		for j := i + 1; j < len(nombres); j++ {
			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
		}

		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]

		// Показываем список после каждого прохода.
		fmt.Println(nombres)
	}
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}
	triSelection(nombres)
}

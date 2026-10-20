package main

import "fmt"

func triSelection(nombres []int) {
	for i := 0; i < len(nombres); i++ {
		// Запоминаем позицию минимального числа.
		indexMin := i

		// Ищем самое маленькое среди оставшихся чисел.
		for j := i + 1; j < len(nombres); j++ {
			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
		}

		// Ставим найденное число на позицию i.
		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]
	}
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}

	triSelection(nombres)

	fmt.Println(nombres)
}

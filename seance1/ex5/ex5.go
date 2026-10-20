package main

import "fmt"

// (int, int) означает: функция возвращает два целых числа.
func triSelection(nombres []int) (int, int) {
	comparaisons := 0 // Счётчик сравнений.
	echanges := 0     // Счётчик обменов.

	for i := 0; i < len(nombres); i++ {
		indexMin := i

		for j := i + 1; j < len(nombres); j++ {
			// Каждый раз, когда сравниваем числа, прибавляем 1.
			comparaisons++

			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
		}

		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]

		// Строка обмена выполнилась — прибавляем 1.
		echanges++
	}

	// Возвращаем два результата в таком порядке.
	return comparaisons, echanges
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}

	// Получаем первый результат в comparaisons,
	// второй — в echanges.
	comparaisons, echanges := triSelection(nombres)

	fmt.Println(nombres)
	fmt.Println("comparaisons :", comparaisons)
	fmt.Println("échanges :", echanges)
}

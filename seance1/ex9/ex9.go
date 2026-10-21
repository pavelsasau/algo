package main

import (
	"fmt"
	"math/rand"
)

// Создаём список из n случайных чисел.
func tranche(n int) []int {
	resultat := make([]int, n)

	for i := 0; i < n; i++ {
		resultat[i] = rand.Intn(100)
	}

	return resultat
}

// Сортировка выбором.
func triSelection(nombres []int) (int, int) {
	comparaisons := 0
	echanges := 0

	for i := 0; i < len(nombres); i++ {
		indexMin := i

		for j := i + 1; j < len(nombres); j++ {
			comparaisons++

			if nombres[j] < nombres[indexMin] {
				indexMin = j
			}
		}

		nombres[i], nombres[indexMin] = nombres[indexMin], nombres[i]
		echanges++
	}

	return comparaisons, echanges
}

// Пузырьковая сортировка.
func triBulles(nombres []int) (int, int) {
	comparaisons := 0
	echanges := 0

	for {
		echangeFait := false

		for i := 0; i < len(nombres)-1; i++ {
			comparaisons++

			if nombres[i] > nombres[i+1] {
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]
				echanges++
				echangeFait = true
			}
		}

		if !echangeFait {
			break
		}
	}

	return comparaisons, echanges
}

func main() {
	fmt.Println("n | Selection : comparaisons | Bulles : comparaisons | Bulles : echanges")

	for _, n := range []int{5, 10, 20, 50} {
		original := tranche(n)

		// Отдельная копия для сортировки выбором.
		copie1 := make([]int, len(original))
		copy(copie1, original)

		// Отдельная копия для пузырьковой сортировки.
		copie2 := make([]int, len(original))
		copy(copie2, original)

		// "_" означает: второй результат нам здесь не нужен.
		compSelection, _ := triSelection(copie1)
		compBulles, echBulles := triBulles(copie2)

		fmt.Println(n, "|", compSelection, "|", compBulles, "|", echBulles)
	}
}

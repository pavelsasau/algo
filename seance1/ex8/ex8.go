package main

import "fmt"

func triBulles(nombres []int) (int, int) {
	// Счётчики создаём до циклов.
	comparaisons := 0
	echanges := 0

	for {
		echangeFait := false

		for i := 0; i < len(nombres)-1; i++ {
			// Сейчас сравним двух соседей — считаем сравнение.
			comparaisons++

			if nombres[i] > nombres[i+1] {
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]

				// Обмен произошёл — считаем его.
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
	nombres := []int{5, 2, 9, 1, 7}

	comparaisons, echanges := triBulles(nombres)

	fmt.Println(nombres)
	fmt.Println("comparaisons :", comparaisons)
	fmt.Println("échanges :", echanges)
}

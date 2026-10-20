package main

import "fmt"

func triBulles(nombres []int) {
	for {
		// В начале прохода обменов ещё не было.
		echangeFait := false

		for i := 0; i < len(nombres)-1; i++ {
			// Сравниваем соседние числа.
			if nombres[i] > nombres[i+1] {
				// Меняем их местами.
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]

				// Запоминаем: обмен был.
				echangeFait = true
			}
		}

		// Если за весь проход обменов не было — заканчиваем.
		if !echangeFait {
			break
		}
	}
}

func main() {
	nombres := []int{5, 2, 9, 1, 7}

	triBulles(nombres)

	fmt.Println(nombres)
}

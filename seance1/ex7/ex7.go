package main

import "fmt"

func triBulles(nombres []int) {
	for {
		// Пока обменов в этом проходе не было.
		echangeFait := false

		for i := 0; i < len(nombres)-1; i++ {
			// Если левое число больше правого:
			if nombres[i] > nombres[i+1] {
				// Меняем соседние числа местами.
				nombres[i], nombres[i+1] = nombres[i+1], nombres[i]

				// Запоминаем, что обмен был.
				echangeFait = true
			}
		}

		// Если за проход не было обменов — останавливаемся.
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

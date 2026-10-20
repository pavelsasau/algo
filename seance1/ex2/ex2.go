package main

import "fmt"

func main() {
	nombres := []int{5, 2, 9, 1, 7}

	// Пока считаем, что минимум находится на позиции 0.
	indexMin := 0

	// Проверяем остальные числа, начиная с позиции 1.
	for i := 1; i < len(nombres); i++ {
		// Если текущее число меньше найденного минимума:
		if nombres[i] < nombres[indexMin] {
			// Запоминаем позицию нового минимума.
			indexMin = i
		}
	}

	// Выводим минимальное число и его позицию.
	fmt.Println("minimum", nombres[indexMin], "en position", indexMin-1)
}

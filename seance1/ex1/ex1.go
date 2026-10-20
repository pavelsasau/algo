package main

import "fmt"

func main() {
	// Список из пяти целых чисел.
	nombres := []int{5, 2, 9, 1, 7}

	// Перебираем позиции от 0 до 4.
	for i := 0; i < len(nombres); i++ {
		// Выводим позицию и число на этой позиции.
		fmt.Println(i, nombres[i])
	}
}

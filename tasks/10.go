// Дана последовательность температурных колебаний:
// -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножествах не важна.

package main

import "fmt"

func main() {
	nums := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	subset := make(map[int][]float32) // Тут будут храниться последовательности
	for _, n := range nums {
		index := int(n) / 10 * 10 // Вычисление подмножества
		subset[index] = append(subset[index], n)
	}
	fmt.Println(subset)
}

// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.

package main

import "fmt"

func main() {
	array := []int{5, 2, 1, 3}
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

// QuickSort - быстрая сортировка массива
//
// first - last рассматривается как часть массива
func QuickSort(arr []int, first, last int) {
	if first < last { // сортируется если first меньше last
		// находим опорным элемент
		p := partition(arr, first, last)
		fmt.Println(arr)
		QuickSort(arr, first, p-1) // сортировка левой части
		QuickSort(arr, p+1, last)  // сортировка правой части
	}
	return
}

// partition возвращает опорную точку и перемещает элементы в правильно сторону от него.
func partition(arr []int, first, last int) int {
	pivot := arr[last] // конечный элемент
	i := first         //
	for j := first; j < last; j++ {
		// если j элемент меньше конечного элемент, то
		// меняем местами его с i элементом и инкриминирует i
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[last] = arr[last], arr[i] // в конце меняем i и last местами

	return i
}

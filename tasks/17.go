// Реализовать бинарный поиск встроенными методами языка.

package main

import "fmt"

// BinarySearch итеративная реализация
func BinarySearch(slice []int, elem int) int {
	// мы рассматриваем first - last как часть массива
	first, last := 0, len(slice)-1
	middle := 0
	for first <= last {
		middle = (first + last) / 2 // вычисления середины
		// сравниваем elem с серединой этой части middle
		if elem > slice[middle] { // если elem > middle, то отсеиваем левую часть
			first = middle + 1
		} else if elem < slice[middle] { // если elem < middle, то правую
			last = middle - 1
		} else {
			return middle // если elem = middle то выводим число.
		}
	}
	return -1 // выводим -1, если не нашли индекс
}

func main() {
	slice := []int{1, 2, 3, 4, 6, 8, 9, 11, 15, 18, 20}
	fmt.Println(BinarySearch(slice, 6))
}

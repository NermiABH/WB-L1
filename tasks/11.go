// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func NewSetInt(nums ...int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}
func main() {
	// Создаю множества через map со значением структуры (т.к. пустая структура почти ничего не весит)
	set1 := NewSetInt(6, 2, 9, 1, 0, 3)
	set2 := NewSetInt(2, 5, 9, 1, 4, 6)

	fmt.Println(IntersectionSetInt(set1, set2))
}

func IntersectionSetInt(set1, set2 map[int]struct{}) map[int]struct{} {
	// В set1 лучше хранить наименьший элемент,
	// чтобы было меньше итераций в цикле
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	// Пересечение двух множеств
	intersect := make(map[int]struct{})
	for n := range set1 {
		// Если элемент есть в другом множестве, то записываем в пересечение
		if _, ok := set2[n]; ok {
			intersect[n] = struct{}{}
		}
	}
	return intersect
}

// Удалить i-ый элемент из слайса
package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	arr = RemoveElemSliceWithAppend(arr, 2)
	fmt.Println(arr)
	arr = RemoveElemSliceChangingOrder(arr, 1)
	fmt.Println(arr)
}

// RemoveElemSliceWithAppend удаляет элементы с помощью append
// удаление довольно медленное
func RemoveElemSliceWithAppend(arr []int, i int) []int {
	return append(arr[:i], arr[i+1:]...)
}

// RemoveElemSliceChangingOrder меняет порядок,
// но удаление выполняется очень быстро
func RemoveElemSliceChangingOrder(arr []int, i int) []int {
	// подставляем к i-ому элементу последнее значение,
	// и попутно это последнее значение обнуляем
	arr[i], arr[len(arr)-1] = arr[len(arr)-1], 0
	return arr[:len(arr)-1] // Срезаем срез без последнего элемента
}

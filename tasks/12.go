// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package main

import "fmt"

func main() {
	arr := [...]string{"cat", "cat", "dog", "cat", "tree"}
	// Создаю множество через map со значением структуры (т.к. пустая структура почти ничего не весит)
	set := make(map[string]struct{})
	for _, s := range arr {
		set[s] = struct{}{}
	}
	fmt.Println(arr)
}

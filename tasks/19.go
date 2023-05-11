// Разработать программу, которая переворачивает подаваемую на ход
// строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import "fmt"

func main() {
	fmt.Println(Reverse("главрыба — абырвалг"))
}

func Reverse(s string) string {
	runes := []rune(s) // Преобразовываем строку в слайс рун
	// i идет слева на право, а j наоборот, пока они не сравняются
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами символы
	}
	return string(runes) // преобразовываем руны в строку
}

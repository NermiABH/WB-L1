// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(AllUniqSymbols("abcd"))
	fmt.Println(AllUniqSymbols("abCdefAf"))
	fmt.Println(AllUniqSymbols("aabcd"))
}

func AllUniqSymbols(s string) bool {
	s = strings.ToLower(s)            // Делаю все символы строчными
	chars := make(map[int32]struct{}) // Сет, в котором будут храниться все символы
	for _, c := range s {
		if _, ok := chars[c]; ok { // если такой символ в сете есть возвращаем false
			return false
		}
		chars[c] = struct{}{} // иначе записываем значение в сет
	}
	return true
}

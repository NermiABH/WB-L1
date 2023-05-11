// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("snow dog sun — sun dog snow"))
}

func ReverseWords(str string) string {
	words := strings.Split(str, " ")        // разделяем слова в массив по пробелам
	newStr := strings.Builder{}             // Создаем string builder
	newStr.Grow(len(str))                   // Задаем capacity
	newStr.WriteString(words[len(words)-1]) // Сразу записываю последнее слово
	for i := len(words) - 2; i >= 0; i-- {
		newStr.WriteString(" ")      // Перед каждым словом записываю пробел
		newStr.WriteString(words[i]) // Само слово
	}
	return newStr.String() // Вывод строки
}

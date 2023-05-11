// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}
//
//func main() {
//  someFunc()
//}

package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	//v[:100] даст 100 байт, но многие символы занимают больше 1 байта,
	//поэтому нам стоит делать срез в рунах
	justString = string([]rune(v)[:100])
}

func main() {
	someFunc()
}

func createHugeString(size int) string {
	var s strings.Builder
	for i := 0; i < size; i++ {
		s.WriteString("a")
	}
	return s.String()
}

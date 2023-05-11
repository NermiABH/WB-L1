// Поменять местами два числа без
// создания временной переменной.
package main

import "fmt"

func main() {
	a, b := 5, 10
	b, a = a, b // Спасибо за это go
	fmt.Println(a, b)

	// Для чисел можно еще так
	a += b
	b = a - b
	a -= b
	fmt.Println(a, b)
}

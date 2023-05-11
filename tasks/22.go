// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a, b, значение которых > 2^20.

package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	// Для использования больших целых чисел лучше использовать bit.Int
	a, ok := big.NewInt(0).SetString("234245234512102938741023847019283401", 10)
	if !ok { // если false
		log.Fatal("Error of creating big.Int")
	}
	b, ok := big.NewInt(0).SetString("231412312351923517325912395197237911", 10)
	if !ok {
		log.Fatal("Error of creating big.Int")
	}
	fmt.Println(a.Add(a, b)) // Сложение двух элементов
	fmt.Println(a.Mul(a, b)) // Умножение
	fmt.Println(a.Sub(a, b)) // Разность
	fmt.Print(a.Div(a, b))   // Деление
}

// Дана переменная int64. Разработать программу,
// которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

// SetBitOne принимает указатель на число nums и индекс бита i,
// чтобы в нем установить 1
func SetBitOne(num *int64, i uint) {
	*num |= 1<<i | *num
}

// SetBitZero принимает указатель на число nums и индекс бита i,
// чтобы в нем установить 0
func SetBitZero(num *int64, i uint) {
	*num &= 1<<i - 1
}
func main() {
	var num int64 = 5
	SetBitOne(&num, 1)
	fmt.Println(num)
	SetBitZero(&num, 2)
	fmt.Println(num)
}

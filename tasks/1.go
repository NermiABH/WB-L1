// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name    string
	Surname string
}

func (h *Human) FullName() string {
	return h.Name + " " + h.Surname
}

// Action через встраивание human, может неявно использовать все его поля и методы
type Action struct {
	Human
}

func (a *Action) Hello() string {
	// Не нужно обращаться к структуре Human, чтобы использовать метод
	return "Hello, I am " + a.FullName()
}

func main() {
	david := Human{"David", "Smyr"}
	action := Action{david}

	fmt.Println(action.Hello())
}

// Разработать программу, которая в рантайме способна
// определить тип переменной: int, string, bool, channel из
// переменной типа interface{}.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		Int     int
		String  string
		Bool    bool
		Channel chan int
	)

	// C помощью библиотеки reflect
	fmt.Println(reflect.TypeOf(any(Int)))
	fmt.Println(reflect.TypeOf(any(String)))
	fmt.Println(reflect.TypeOf(any(Bool)))
	fmt.Println(reflect.TypeOf(any(Channel)))
}

// Можно и так ...
func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "integer"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan int"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}

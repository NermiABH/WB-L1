// Реализовать собственную функцию sleep.
package main

import "time"

func main() {
	n := 1
	Sleep(time.Duration(n))
}

func Sleep(d time.Duration) {
	// Поток останавливается пока в канал не придет значение
	<-time.After(time.Second * d)
}

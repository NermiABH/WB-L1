// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// Counter удобно делать структурой, чтобы использовать mutex
type Counter struct {
	Val int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock() // Блокировка мьютекса, чтобы избежать состояния гонки
	c.Val++
	c.Unlock()
}

func main() {
	done := make(chan os.Signal)
	signal.Notify(done, os.Interrupt) // Сигнал на прерывание программы

	counter := &Counter{} // Инициализирую счетчик
	for i := 0; i < 3; i++ {
		go func() { // создание горутин которые постоянно инкриминируют счетчик
			for {
				counter.Increment()
			}
		}()
	}

	<-done                       // Когда получит сигнал программа остановится
	fmt.Print("\n", counter.Val) // Вывод значения счетчика
}

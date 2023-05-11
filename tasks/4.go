// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Количество воркеров
	var workers int
	// Канал в которые будут записываться данные
	ch := make(chan int)
	_, err := fmt.Scan(&workers)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < workers; i++ {
		//Создание воркера
		go func() {
			for d := range ch {
				fmt.Println(d)
			}
		}()
	}

	// Создание сигнала для прерывания
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt)
	rand.Seed(time.Now().UnixNano())

	// Если получим сигнал, программа завершится
	// Или запись случайного числа в канал
	for {
		select {
		case <-done:
			fmt.Println("Программа завершилась")
			return
		default:
			ch <- rand.Intn(100)
		}

	}
}

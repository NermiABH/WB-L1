// Реализовать все возможные способы остановки
// выполнения горутины.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//1
	ch := make(chan struct{})
	go stopRecCh(ch)
	ch <- struct{}{}
	time.Sleep(time.Second / time.Duration(2))

	//2
	go stopClosedCh(ch)
	close(ch)
	time.Sleep(time.Second / time.Duration(2))

	//3
	ctx, cancel := context.WithCancel(context.Background())
	go stopCancelContext(ctx)
	cancel()
	time.Sleep(time.Second / time.Duration(2))

	//4
	ctx, _ = context.WithTimeout(context.Background(), time.Second/time.Duration(4))
	go stopTimeOutContext(ctx)
	time.Sleep(time.Second / time.Duration(2))

	//5
	for i := 0; i < 5; i++ {
		go stopWithMain()
	}
	fmt.Println("main завершается, все работающие горутины тоже завершатся")
}

// 1. Остановка когда получаем данные с канала
func stopRecCh(ch chan struct{}) {
	for {
		select {
		case <-ch:
			fmt.Println("Получил данные с канала, останавливаем горутину")
			return
		}
	}
}

// 2. Остановка если канал закрыт
func stopClosedCh(ch chan struct{}) {
	for {
		if _, opened := <-ch; !opened {
			fmt.Println("Канал закрыт, останавливаем горутину")
			return
		}
	}
}

// 3. Остановка через cancel в context
func stopCancelContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст canceled, останавливаем горутину")
			return
		}
	}
}

// 4. Остановка горутины через таймаут контекста
func stopTimeOutContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Время контекста вышло, останавливаем горутину")
			return
		}
	}
}

// 5. Остановка горутины когда остановится главный поток
func stopWithMain() {
	select {}
}

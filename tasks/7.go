// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		// Mьютекс который и будет синхронизировать операции
		mu   sync.Mutex
		wg   sync.WaitGroup
		nums = make(map[int]int)

		// Функция добавления значений в мапу
		addToMap = func(i int) {
			mu.Lock() // Блокировка доступ, чтобы не было состояния гонки в мапе
			nums[i] = i
			mu.Unlock() // Разблокировка доступа
			wg.Done()
		}
	)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go addToMap(i)
	}
	wg.Wait()
	fmt.Println(nums)
}

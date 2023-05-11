// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов (2^2+3^2+4^2 ...) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var (
		nums = [...]int{2, 4, 6, 8, 10}
		wg   sync.WaitGroup
		sum  uint32
	)
	wg.Add(len(nums))
	for _, num := range nums {
		go func(n int) {
			defer wg.Done()
			// Обязательно делаем атомарное добавление,
			//чтобы не произошло состояние гонки между горутинами.
			atomic.AddUint32(&sum, uint32(n*n))
		}(num)
	}

	wg.Wait()
	fmt.Println(sum)
}

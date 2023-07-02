package main

import (
	"fmt"
)

func main() {
	ch, cancel := enumerate(10)
	// cчитывать из канала пока он не будет закрыт
	for i := range ch {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}

	// преклащаем выполнение горутины
	cancel()
}

func enumerate(n int) (<-chan int, func()) {
	// создаём канал
	ch := make(chan int)

	// реализуем паттерн done
	done := make(chan struct{})

	// создаём закрывающую функцию
	cancel := func() {
		close(done)
	}

	go func() {
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			default:
				ch <- i
			}
		}

		// range из цикла ожидает закрытие канала,
		close(ch)
	}()

	return ch, cancel
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	// реализация ожидания окончания работы горутин
	var wg2 sync.WaitGroup

	// создание канала, каналы являются указателями, как и карты
	ch := make(chan int)

	a := []int{2, 4, 6, 8, 10}
	wg2.Add(1)
	go func() {
		defer wg2.Done()

		for _, value := range a {
			// записываем в канал значения
			ch <- value * 2
		}
		// закрытие канала, т.к. окончание горутины не гарантирует его закрытие
		close(ch)
	}()

	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for i := 0; i < len(a); i++ {
			// считываем значения из канала
			value := <-ch
			fmt.Printf("%d ", value)
		}
	}()

	wg2.Wait()
}

package main

import (
	"fmt"
)

func main() {
	// считывать из канала пока он не будет закрыт
	for i := range calculate(10) {
		fmt.Printf("%d ", i)
	}
}

func calculate(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
		}

		// range из цикла ожидает закрытие канала,
		// если этого не выполнить, то произойдёт deadlock
		// окончание горутины не гарантирует его закрытие
		close(ch)
	}()

	return ch
}

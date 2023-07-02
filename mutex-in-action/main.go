package main

import (
	"fmt"
	"sync"
	"time"
)

type Basket struct {
	// мьютекс, позволяющий совместную блокировку на читение,
	// но запрещающий блокировку, если была установлена блокировка на запись
	mutex sync.RWMutex

	store map[int]string
}

func NewBasket() *Basket {
	return &Basket{
		store: map[int]string{},
	}
}

func (b *Basket) Get(key int) (string, bool) {
	b.mutex.RLock()         // блокирует только если мьютекс был заблокирован на запись
	defer b.mutex.RUnlock() // обязательно позаботиться о высвобождении блокировки

	value, ok := b.store[key]
	return value, ok
}

func (b *Basket) Put(key int, value string) {
	b.mutex.Lock()         // блокирует всех
	defer b.mutex.Unlock() // обязательно позаботиться о высвобождении блокировки

	b.store[key] = value
}

func main() {
	basket := NewBasket()
	goods := []string{"chair", "table", "wardrobe", "bed", "knife", "spoons", "forks", "cups"}

	// закинуть в корзину товары в несколько горутин (блокровка на запись)
	for i, value := range goods {
		go func(key int, value string) {
			basket.Put(key, value)
			time.Sleep(time.Millisecond * 50)
		}(i, value)
	}

	// реализация ожидания окончания работы горутин
	var wg sync.WaitGroup

	// бесконечно в несколько горутин пытаться получить товары по ключу (блокровка на чтение)
	for i := 0; i < len(goods); i++ {
		wg.Add(1)

		go func(key int) {
			defer wg.Done()
			for {
				value, ok := basket.Get(key)
				if ok {
					fmt.Println(value)
					break
				}
			}
		}(i)
	}

	wg.Wait()
}

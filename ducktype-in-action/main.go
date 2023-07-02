package main

import (
	"ducktype-in-action/engineer"
	"fmt"
)

type Employee interface {
	Do(string)
}

func main() {
	// можно создать экземпляр как Engineer так и Manager,
	// согласно утиной типизации они оба могут быть Employee и выполнить работу
	eng := new(engineer.Engineer)
	eng.Name = "Peter"

	// можно заметить, что передаётся указатель на структуру,
	// причина в том, что Go определяет соответствие по совпадению функций,
	// которые реализованы у структуры, а реализация Do в пакете engineer требует указатель
	heyDoIt(eng, "drinks coffee")

	// интерфейс сам по себе может хранить как структуру так и указатель на неё
}

func heyDoIt(employee Employee, work string) {
	employee.Do(work)

	// выведем какой тип хранит интерфейс и какого содержимое структуры
	fmt.Printf("[%T] %+v\n", employee, employee)

	// результат: [*engineer.Engineer] &{Name:Peter}
}

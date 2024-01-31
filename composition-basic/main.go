package main

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e *Employee) WhoAmi() {
	fmt.Println("Сотрудник")
}

func (e *Employee) Description() {
	fmt.Printf("%s, id: %s\n", e.Name, e.Id)
}

type Manager struct {
	Employee *Employee
	Position string
}

func (e *Manager) WhoAmi() {
	fmt.Println("Управляющий")
}

func main() {
	// создадим сотрудника
	e := &Employee{
		Name: "Андрей",
		Id:   "452",
	}

	// выведем кем является структура
	e.WhoAmi()
	e.Description()

	// результат: Сотрудник
	// результат: Андрей, id: 452

	// на основе сотрудника создадим управляющего
	m := Manager{
		Employee: e,
		Position: "55",
	}

	// выведем кем является структура
	// доступ к полям Id и Name осуществляется только через Employee
	// прямой доступ можно иметь только через встраивание структуры (forwarding, отдельная тема)
	m.WhoAmi()
	m.Employee.Description()

	// результат: Управляющий
	// результат: Андрей, id: 452
}

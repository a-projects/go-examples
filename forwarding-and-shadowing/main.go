package main

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e *Employee) WhoAmi() {
	fmt.Println("Сотрудник")
	e.Description()
}

func (e *Employee) Description() {
	fmt.Printf("%s, id: %s\n", e.Name, e.Id)
}

type Manager struct {
	*Employee
	Position string
	Employes []Employee
}

// затенение метода WhoAmi от Employee
func (e *Manager) WhoAmi() {
	fmt.Println("Управляющий")
	e.Description()
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

	// на основе сотрудника создадим управляющего (встроенная структура сотрудника)
	// управлющий будет имеет теже поля и методы, что и сотрудник
	// но можно заметить, что метод WhoAmi переопределён (затенён)
	m := Manager{
		Employee: e,
		Position: "55",
	}

	// выведем кем является структура (затенённый метод WhoAmi)
	m.WhoAmi()
	m.Description()

	// результат: Управляющий
	// результат: Андрей, id: 452

	// Присваивание работает
	//var e1 Employee = Employee{}

	// Ошибка компиляци т.к. встраивание это не наследование
	// Если необходим функционал полиморфизма, то нужно использовать интерфейсы
	//var e2 Employee = Manager{}
}

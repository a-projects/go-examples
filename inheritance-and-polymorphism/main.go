package main

import "fmt"

// определим интерфейс всех автомобилей
type Driver interface {
	Go()
	Drive()
}

type Car struct {
	// номер машины
	Number string
}

func (c *Car) GetNumber() string {
	return c.Number
}

// определим метод согласно интерфейсу
// не будем переопределять (затенять) метод и посмотрим как головные структуры им воспользуются
func (c *Car) Go() {
	fmt.Print("Едет ")
}

// определим метод согласно интерфейсу
// позже переопределим (затеним) метод в головных структурах, чтобы привнести уникальную реализацию
func (c *Car) Drive() {
	fmt.Printf("машина с номером: %s\n", c.Number)
}

type MedicCar struct {
	Car
}

// переопределим Drive и воспользуемся полем Number из встроенной структуры
func (m MedicCar) Drive() {
	fmt.Printf("скорая с номером: %s\n", m.Number)
}

type PoliceCar struct {
	Car
}

// переопределим Drive и воспользуемся методом GetNumber из встроенной структуры
func (p PoliceCar) Drive() {
	fmt.Printf("полиция с номером: %s\n", p.GetNumber())
}

func main() {
	cars := []Driver{
		&Car{Number: "AC654RU"},
		&MedicCar{Car: Car{Number: "MC228RU"}},
		&PoliceCar{Car: Car{Number: "PC660RU"}},
	}

	for _, car := range cars {
		car.Go()
		car.Drive()
	}

	// результат:
	// Едет машина с номером: AC654RU
	// Едет скорая с номером: MC228RU
	// Едет полиция с номером: PC660RU
}

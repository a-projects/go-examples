package main

import (
	"fmt"
)

type Pet interface {
	Say(string)
}

type Dog struct {
	Name string
}

// функция принимающая тип в качестве значения
func (d Dog) Say(msg string) {
	fmt.Printf("Dog '%s' say %s\n", d.Name, msg)
}

type Cat struct {
	Name string
}

// функция принимающая тип в качестве указателя
func (c *Cat) Say(msg string) {
	fmt.Printf("Cat '%s' say %s\n", c.Name, msg)
}

func main() {

	dog := Dog{
		Name: "Born",
	}

	cat := Cat{
		Name: "Barsik",
	}

	// передаём структуру как значение
	Voice(dog, "woof")

	// передаём указатель на структуру
	Voice(&cat, "meow")
}

func Voice(foo Pet, msg string) {
	foo.Say(msg)
}

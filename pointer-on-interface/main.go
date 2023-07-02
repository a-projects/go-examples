package main

import (
	"fmt"
)

type Pet interface {
	Say(string)
}

type Cat struct {
	Name string
}

func (c *Cat) Say(msg string) {
	fmt.Printf("Cat '%s' say %s\n", c.Name, msg)
}

func main() {
	cat := &Cat{
		Name: "Vasya",
	}

	// передадим указатель на интерфейс и
	// попробуем им восопльзоваться
	var pet Pet = cat
	training(&pet)

	// передадим как интерфейс, и как указатель на интерефейс,
	// посмотрим какой тип и содержимое будет в выводе
	reflect(pet)
	reflect(&pet)

	// вывод:
	// [*main.Cat] &{Name:Vasya}
	// [*main.Pet] 0xc000052280
}

func training(pet *Pet) {
	(*pet).Say("meow")
}

func reflect(t interface{}) {
	switch t := t.(type) {
	case *Cat:
		fmt.Printf("[%T] %+v\n", t, t)
	case *Pet:
		fmt.Printf("[%T] %+v\n", t, t)
	}
}

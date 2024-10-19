package main

import (
	"basics/packages"
	"fmt"
)

type Person struct {
	Name string
	age  int
}

type address struct {
	house string
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		age:  age,
	}
}

func (p *Person) changeName(NewName string) {
	p.Name = NewName
}

func main() {
	newCircle := packages.NewCicle(4.0)

	circlecum := newCircle.Cal_Circle_Circum()
	area := newCircle.Cal_Circle_Area()
	fmt.Println(circlecum, area)
}

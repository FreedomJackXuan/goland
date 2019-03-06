package interfacetest

import "fmt"

type Animal interface {
	GetName()
	EatFood()
}

type AnimalInfo struct {
	name string
	food string
}

func (a *AnimalInfo) GetName() {
	fmt.Println(a.name)
}

func (a *AnimalInfo) EatFood() {
	fmt.Println(a.food)
}

func Inter2() {
	var dog Animal = &AnimalInfo{"dog", "gutou"}
	var cat Animal = &AnimalInfo{"cat","fish"}

	dog.EatFood()
	cat.EatFood()
}
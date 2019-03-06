package interfacetest

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func Inter1() {
	//dog := Dog{"wang"}
	//var pet Pet = &dog
	//dog.SetName("xxxxxx")
	//
	//fmt.Println(pet.Category())

}
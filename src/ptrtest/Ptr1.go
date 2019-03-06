package ptrtest

type Named interface{
	P1Name() string
}

type Pdog struct {
	name string
}

func (dog *Pdog) SetName(name string) {
	dog.name = name
}

func (dog Pdog) P1Name() string {
	return dog.name
}


func Ptr1() {
	const num  =123
	//_=&num
	var str = "abc"
	_=&str
	//_=&(str[0])
}

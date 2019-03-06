package channeltest

import (
	"fmt"
	"math/rand"
)

func Chantest4() {
	var uselessChan = make(chan<- int,1)
	var anotherChan = make(<-chan int, 1)
	fmt.Printf("The useless channels: %v, %v\n",
		uselessChan, anotherChan)

	intChan1 := make(chan int, 3)
	SendInt(intChan1);

	intChan2 := getIntChan()
	for elem := range intChan2{
		fmt.Println(elem)
	}
	intChan3 := GetIntChan(getIntChan)()
	for elem := range intChan3 {
		fmt.Println(elem)
	}
}

func SendInt(ch chan<- int)  {
	ch <- rand.Intn(1000)
}

func getIntChan() <-chan int{
	num:=5
	ch := make(chan int, num)
	for i := 0; i < num; i++ {
		ch <- i
	}
	close(ch)
	return ch
}

type GetIntChan func() <-chan int
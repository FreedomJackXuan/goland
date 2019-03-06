package panictest

import (
	"errors"
	"fmt"
)

func Panic2(){

	//处理panic只能通过defer处理，而且recover必须放在panic之前，否则捕获不到
	defer func() {
		fmt.Println("defer function")
		//if p := recover(); p != nil {
		//	fmt.Println("panic:",p)
		//}

		//在defer中可以引发panic,但是处理panic的时候需要在defer  defer中
		defer func() {
			if p := recover(); p != nil {
				fmt.Println(p)
			}
		}()
		panic(errors.New("xxxxxxxxxxxxxx"))
		fmt.Println("exit")
	}()
	//error
	//fmt.Printf("no panic: %v\n", recover())

	//panic(errors.New("wrong"))

	//error
	//p:=recover()
	//fmt.Println(p)

}
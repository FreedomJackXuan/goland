package panictest

import (
	"fmt"
	"errors"
)


func PanicTest1() {
	fmt.Println("xxx")
	call()
}

func call() {
	fmt.Println("call")
	panic(errors.New("wrong"))
	panic(fmt.Println)
	fmt.Println("exit")
}

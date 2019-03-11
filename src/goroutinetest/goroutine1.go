package goroutinetest

import (
	"fmt"
	"runtime"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		fmt.Print("hello ")
		runtime.Gosched() //释放CPU权限
	}
}

func sayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println("world")
		runtime.Gosched()
	}
}

func Goroutine1() {
	go sayHello()
	go sayWorld()
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	var st string
	fmt.Scan(&st)
}

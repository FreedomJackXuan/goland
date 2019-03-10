package synctest

import (
	"sync"
	"sync/atomic"
	"fmt"
	"unsafe"
)

/*
单例
*/
type singleton struct {}

var (
	instance *singleton
	initialized uint32
	mu sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		fmt.Println("11111")
		return instance
	}
	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized,1)
		instance = &singleton{}
	}
	//fmt.Println(initialized)
	return instance
}

func Sync3() {
	s := Instance()
	fmt.Println(uintptr(unsafe.Pointer(s)))
	fmt.Println(initialized)
	s1 := Instance()
	fmt.Println(uintptr(unsafe.Pointer(s1)))
	s2 := Instance()
	fmt.Println(uintptr(unsafe.Pointer(s2)))
}

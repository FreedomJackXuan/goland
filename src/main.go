package main

import (
	"unsafe"
	"synctest"
)

type MyTemplate struct {
	name       string
	parseTree  *unsafe.Pointer
	common     *unsafe.Pointer
	leftDelim  string
	rightDelim string
}

func main() {

	synctest.Sync1()

	//fmt.Println("aaaaaaa")
	//defer func() {
	//	fmt.Println("bbbbbbbbbbb")
	//	defer func() {
	//		fmt.Println("eeeeeeeeeee")
	//	}()
	//}()
	//fmt.Println("ccccccccccccccccccc")


	//value1 := [...]int{0, 1, 2, 3, 4, 5, 6}
	//switch 1 + 3 {
	//case value1[1], value1[0]:
	//	fmt.Println("0 or 1")
	//case value1[2], value1[3]:
	//	fmt.Println("2 or 3")
	//case value1[4], value1[5], value1[6]:
	//	fmt.Println("4 or 5 or 6")
	//}


	//numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	////maxIndex2 := len(numbers2) - 1
	//for i, e := range numbers2 {
	//	//if i == maxIndex2 {
	//	//	numbers2[0] += e
	//	//} else {
	//	//	numbers2[i+1] += e
	//	//}
	//	fmt.Println(i,e)
	//}
	//fmt.Println(numbers2)


	//Go 语言并不会去保证这些 goroutine 会以怎样的顺运行
	//for i := 0; i < 10; i++ {
	//	go func() {
	//		fmt.Println(i)
	//	}()
	//}

	//num := 10
	//sign := make(chan struct{}, num)
	//
	//for i := 0; i < num; i++ {
	//	go func() {
	//		fmt.Println(i)
	//		sign <- struct{}{}
	//	}()
	//}
	//
	//// 办法1。
	////time.Sleep(time.Millisecond * 500)
	//
	//// 办法2。
	//for j := 0; j < num; j++ {
	//	<- sign
	//
	//}

	//var count uint32
	//trigger := func(i uint32, fn func()) {
	//	for {
	//		if n := atomic.LoadUint32(&count); n == i {
	//			fn()
	//			atomic.AddUint32(&count, 1)
	//			break
	//		}
	//		time.Sleep(time.Nanosecond)
	//	}
	//}
	//for i := uint32(0); i < 10; i++ {
	//	go func(i uint32) {
	//		fn := func() {
	//			fmt.Println(i)
	//		}
	//		trigger(i, fn)
	//	}(i)
	//}
	//trigger(10, func() {})

	//u := uint32(32)
	//i := int32(1)
	//fmt.Println(&u, &i)
	//p := &i
	//p = (*int32)(unsafe.Pointer(&u))
	//fmt.Println(p)
	//a := A{"aa",12}
	//xa := &a
	//fmt.Println(unsafe.Pointer(xa))
	//fmt.Println(uintptr(unsafe.Pointer(xa)))
	//fmt.Println(uintptr(unsafe.Pointer(xa))+unsafe.Offsetof(xa.name))
	//fmt.Println()
	//ptrtest.Prt2()
}


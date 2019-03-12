package cgotest

/*
#include <stdio.h>

void printint(int v){
	printf("printint: %d\n", v);
}
*/
import "C"

func Test1() {
	v := 42
	C.printint(C.int(v))
}


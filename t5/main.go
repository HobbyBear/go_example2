package main

import (
	"fmt"
	"unsafe"
)

func Add(a int32, b int32) int32 {
	fmt.Println(unsafe.Sizeof(a))
	return a + b
}

func main() {
	Add(3, 4)
}

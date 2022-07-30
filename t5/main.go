package main

func Add(a int32, b int32) (int32, bool) {
	//fmt.Println(unsafe.Sizeof(a))
	return a + b, true
}

func main() {
	Add(3, 4)
}

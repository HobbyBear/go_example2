package main

import "fmt"

func haha() int {
	a := 1
	defer func() {
		a++
		fmt.Println(a, "========")
	}()
	fmt.Println(a)
	return a
}

func main() {
	haha()
}

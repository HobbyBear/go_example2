package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {

	str := strings.NewReader("w")

	buf := bufio.NewReader(str)

	data, err := buf.Peek(2)
	fmt.Println(string(data), err)
}

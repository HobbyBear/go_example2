package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal(err)
	}

	i := 0
	for {
		buf := make([]byte, 20)
		n, err := conn.Read(buf)
		fmt.Println(n, err)
		if err != nil {
			if i >= 1 {
				return
			}
			i++
			time.Sleep(2 * time.Second)
			err = conn.Close()
			n, err = conn.Write([]byte("haha"))
			if err != nil {
				fmt.Println("写入", n)
			}
			fmt.Println("成功写入", n)
			n, err = conn.Write([]byte("haha"))
			if err != nil {
				fmt.Println("写入2", n)
			}
			fmt.Println("成功写入2", n)
		}
	}

}

package main

import (
	"fmt"
	"log"
	"net"
	"syscall"
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
			//err = conn.Close()
			//n, err = conn.Write([]byte("haha"))
			//if err != nil {
			//	fmt.Println("写入", n)
			//}
			//time.Sleep(3 * time.Second)
			tcp := conn.(*net.TCPConn)
			f, _ := tcp.File()

			n, err = syscall.Write(int(f.Fd()), []byte("haha"))
			if err != nil {
				fmt.Println("写入", n, err)
				return
			}
			fmt.Println("成功写入", n)
			time.Sleep(2 * time.Second)
			//n, err = syscall.Write(int(f.Fd()), []byte("haha"))
			//if err != nil {
			//	fmt.Println("写入2", n, err)
			//	return
			//}
			n, err = conn.Write([]byte("haha"))
			if err != nil {
				fmt.Println("写入", n, err)
				return
			}
			fmt.Println("成功写入2", n)
		}
	}

}

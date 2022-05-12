package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func PrintLocalDial(network, addr string) (net.Conn, error) {
	dial := net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	conn, err := dial.Dial(network, addr)
	if err != nil {
		return conn, err
	}
	fmt.Println("conn 握手完成")
	time.Sleep(3 * time.Second)
	fmt.Println("connect done, use", conn.LocalAddr().String())
	return conn, err
}

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			IdleConnTimeout: time.Second * 1000,
			Dial:            PrintLocalDial,
		},
	}
	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()
}

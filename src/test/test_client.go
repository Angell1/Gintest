package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1.创建链接远程链接服务器，得到一个conn链接
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}
	i := 1
	for {
		//2.调用链接Write写数据
		fmt.Println("Hello Server", i)
		_, err := conn.Write([]byte(fmt.Sprintf("%s:%d", "Hello Server", i)))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}
		i++
		time.Sleep(time.Duration(1)*time.Second)
	}
}
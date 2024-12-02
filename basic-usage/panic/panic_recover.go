// 用例1：避免恐慌导致程序崩溃
// 运行此服务器程序，并在另一个终端窗口运行telnet localhost 12345，我们可以观察到服务器程序不会因为客户连接处理协程中的产生的恐慌而导致崩溃。
package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		// 在一个新协程中处理客户端连接。
		go ClientHandler(conn)
	}
}

func ClientHandler(c net.Conn) {
	defer func() {
		if v := recover(); v != nil {
			log.Println("捕获了一个恐慌：", v)
			log.Println("防止了程序崩溃")
		}
		c.Close()
	}()
	panic("未知错误") // 演示目的产生的一个恐慌
}

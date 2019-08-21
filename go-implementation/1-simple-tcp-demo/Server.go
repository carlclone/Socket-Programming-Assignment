package main

import (
	"fmt"
	"net"
	"time"
)

/*
开启监听
accept阻塞
每个客户端都开新协程进行 阻塞读, 被动回复
*/

func main() {
	var (
		listener net.Listener
		err      error
		conn     net.Conn
	)

	listener, err = net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("监听出错", err)
		return
	}

	defer listener.Close()

	conn, err = listener.Accept()
	if err != nil {
		fmt.Println("accept出错", err)
		return
	}

	go func(conn net.Conn) {
		var (
			buf []byte
			n   int
		)
		defer conn.Close()
		buf = make([]byte, 5000)
		for {
			n, err = conn.Read(buf)
			if err != nil {
				fmt.Println("读取出错", err.Error())
				return
			}
			fmt.Println("接收到数据:", string(buf[:n]))

			_, err = conn.Write([]byte("copy that , i am server!"))
			if err != nil {
				fmt.Println("写入错误", err)
				return
			}
		}
	}(conn)

	for {
		time.Sleep(1 * time.Second)
	}
}

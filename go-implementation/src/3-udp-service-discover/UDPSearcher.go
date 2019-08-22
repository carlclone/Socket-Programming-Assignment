package main

/*
	open a local udp socket
	listen for msg at port 2001 in goroutine
	broadcastMsg = "SEARCH_HEADER"
	send a broadcast to port 2000
	print the remote addr ip and port
	done
*/
import (
	"fmt"
	"net"
	"os"
	"time"

	//  "io"
)

func main() {
	laddr := net.UDPAddr{
		//IP:   net.IPv4(192,168,1,186),
		Port: 3000,
	}
	// 这里设置接收者的IP地址为广播地址
	raddr := net.UDPAddr{
		IP:   net.IPv4(255, 255, 255, 255),
		Port: 11110,
	}
	conn, err := net.DialUDP("udp",&laddr, &raddr)
	defer conn.Close()
	if err != nil {
		os.Exit(1)
	}

	go func() {
		var (
			n int
		)
		var msg [20]byte
		for {
			n,err = conn.Read(msg[0:])
			if n==0 {
				continue
			}

			fmt.Println("msg is", string(msg[0:10]))
		}

	}()
	conn.Write([]byte("Hello world!"))

	fmt.Println("send msg")

	for {
		time.Sleep(1*time.Second)
	}
}

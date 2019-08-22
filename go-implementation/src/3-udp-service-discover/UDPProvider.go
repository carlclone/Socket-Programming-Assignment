package main

import (
	"fmt"
	"net"
	"os"
)
/*
	listen for a port 2000

	received a packet

	if packet has a header "SEARCH_HEAD"
		ip = getIP()
		port = getPort()
		msg = "SEARCH_RESULT".ip.port
		response back a packet to the searcher ip and port
*/

func checkError(err error){
	if  err != nil {
		fmt.Println("Error: %s", err.Error())
		os.Exit(1)
	}
}

func recvUDPMsg(conn *net.UDPConn){
	var buf [20]byte

	n, raddr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("msg is ", string(buf[0:n]))

	//WriteToUDP
	//func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
	fmt.Println(raddr)
	_, err = conn.WriteToUDP([]byte("nice to see u"), raddr)
	checkError(err)
}

func main() {
	udp_addr, err := net.ResolveUDPAddr("udp", ":11110")
	fmt.Println(udp_addr)
	checkError(err)

	conn, err := net.ListenUDP("udp", udp_addr)
	defer conn.Close()
	checkError(err)

	//go recvUDPMsg(conn)
	for {
		recvUDPMsg(conn)
	}

}

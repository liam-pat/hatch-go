package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

// nc -u 127.0.0.1 1200
func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":1200")
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "resolve udp error %s", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "listen udp error %s", err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}
	daytime := time.Now().String()

	conn.WriteToUDP([]byte("server datetime: "+daytime+"\n"), addr)
	conn.WriteToUDP([]byte("content: "+string(buf[0:])+"\n"), addr)

}

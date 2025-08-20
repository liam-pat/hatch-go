package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

// nc -u 127.0.0.1 1200
func main() {
	udpAddr, err := net.ResolveUDPAddr("udp4", ":1200")

	if err != nil {
		log.Fatal(os.Stderr, "resolve udp error %s", err.Error())
	}
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(os.Stderr, "listen udp error %s", err.Error())
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

	daytime := time.Now().Format(time.RFC3339)

	conn.WriteToUDP([]byte(fmt.Sprintf("[%s %s] receive: %s\n", addr, daytime, buf[:])), addr)
}

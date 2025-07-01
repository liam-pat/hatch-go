package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	C    chan string //client user sent data channel
	Name string      //client user name
	Addr string      //client address
}

var onlineMap = make(map[string]client)  //save the online client
var broadcastMessage = make(chan string) // broadcast to online client

func main() {
	listener, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal("init server listener err", err)
	}
	defer func(listener net.Listener) { _ = listener.Close() }(listener)
	// waiting message put to broadcastMessage, and send to all client's channel
	go func() {
		for {
			msg := <-broadcastMessage
			for _, onlineClient := range onlineMap {
				onlineClient.C <- msg
			}
		}
	}()
	// main to wait clients to connect 8001 , you can using `nc localhost 8001` to connect
	for {
		con, err1 := listener.Accept()
		if err1 != nil {
			log.Fatal("accept client err", err)
		}
		go handleCon(con)
	}
}

func handleCon(conn net.Conn) {
	defer func(conn net.Conn) { _ = conn.Close() }(conn)

	clientAddr := conn.RemoteAddr().String()
	clientExample := client{make(chan string), clientAddr, clientAddr}
	onlineMap[clientAddr] = clientExample

	// one client run a goroutine to waiting message write to channel
	go func(conn net.Conn, clientExample client) {
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		for msg := range clientExample.C {
			_, err := conn.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("write msg to client error", err)
				return
			}
		}
	}(conn, clientExample)

	// broadcast all client one client is online
	fmt.Printf("client %s online\n", clientExample.Addr)
	broadcastMessage <- fmt.Sprintf("[%s]%s: First time to login", clientExample.Addr, clientExample.Name)

	isQuit := make(chan bool)
	hasData := make(chan bool)

	// wait the one client to operate
	go func() {
		buf := make([]byte, 2048)
		for {
			// read the message from one client ctrl + c to exit , n == 0
			n, _ := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				return
			}

			msgFromClient := string(buf[:n-1])
			if len(msgFromClient) == 3 && msgFromClient == "who" {
				_, err := conn.Write([]byte("########### user list #########:\n"))
				if err != nil {
					fmt.Println("write the message to client error", err)
					return
				}
				for _, tmp := range onlineMap {
					msg := fmt.Sprintf("[%s] %s\n", tmp.Addr, tmp.Name)
					_, err1 := conn.Write([]byte(msg))
					if err1 != nil {
						fmt.Println("write the message to client error", err)
						return
					}
				}
			} else if len(msgFromClient) >= 8 && msgFromClient[:6] == "rename" {
				name := strings.Split(msgFromClient, "|")[1]
				clientExample.Name = name
				onlineMap[clientAddr] = clientExample
				_, err := conn.Write([]byte("########### user rename successfully ###########"))
				if err != nil {
					fmt.Println("write the message to client error", err)
					return
				}
			} else {
				broadcastMessage <- fmt.Sprintf("[%s]%s: %s", clientExample.Addr, clientExample.Name, msgFromClient)
			}
			hasData <- true
		}
	}()

	for {
		select {
		case <-isQuit:
			delete(onlineMap, clientAddr)
			broadcastMessage <- fmt.Sprintf("#############")
			broadcastMessage <- fmt.Sprintf("[%s] log out", clientAddr)
			broadcastMessage <- fmt.Sprintf("#############")
			return
		case <-hasData:
			// has data will reset the timeout time

		case <-time.After(60 * time.Second):
			broadcastMessage <- fmt.Sprintf("[%s] timeout to log out", clientAddr)
			return
		}
	}
}

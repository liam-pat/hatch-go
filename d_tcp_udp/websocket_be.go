package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

func Echo(ws *websocket.Conn) {
	var err error
	for {
		var feMsg string

		if err = websocket.Message.Receive(ws, &feMsg); err != nil {
			fmt.Println("Can't receive frontend msg, err :", err)
			break
		}

		fmt.Println("msg from frontend: " + feMsg)

		data := make(map[string]interface{}, 1)
		data["server_time"] = time.Now().Format("2006-01-02 15:04:05")
		data["frontend_msg"] = feMsg

		reply, _ := json.Marshal(data)
		fmt.Println("return to frontend msg", string(reply))

		if err = websocket.Message.Send(ws, string(reply)); err != nil {
			fmt.Println("Can't send the msg to frontend")
			break
		}
	}
}

func main() {

	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("listen err.:", err)
	}
}

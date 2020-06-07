package main

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

const (
	host     = "10.50.1.182"
	user     = "arturo"
	password = "Hydr0gen7892"

	msg = `{ "id": "6e515df6-a861-11ea-bb37-0242ac130002", "jsonrpc": "2.0", "method": "session.signIn", "params": { "email": "arturo", "password": "Hyd0gen7892" } }`
)

func main() {
	c, _, err := websocket.DefaultDialer.Dial("ws://"+host+"/api/", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer c.Close()

	if err = c.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Sent message")

	c.SetReadDeadline(time.Now().Add(5 * time.Second))
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(message))
	}
}

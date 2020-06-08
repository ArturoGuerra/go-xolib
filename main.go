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
	id       = "a6add22b-d594-443f-a431-e461c6f628c6"

	msg  = `{ "id": "6e515df6-a861-11ea-bb37-0242ac130002", "jsonrpc": "2.0", "method": "session.signIn", "params": { "email": "arturo", "password": "Hydr0gen7892" } }`
	disk = `{"id":"a6add22b-d594-443f-a431-e461c6f628c6","jsonrpc":"2.0","method":"disk.create","params":{"name":"test","size":"10G","sr":"5fa59666-11cc-258e-aa26-99bd09c9dbef"}}`
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

	c.SetReadDeadline(time.Now().Add(time.Second))
	_, message, err := c.ReadMessage()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(message))
	}

	if err = c.WriteMessage(websocket.TextMessage, []byte(disk)); err != nil {
		fmt.Println(err)
		return
	}

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(message))
		}
	}
}

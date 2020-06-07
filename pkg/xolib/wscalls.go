package xolib

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pborman/uuid"
)

func rawCall(ws *websocket.Conn, req *MessageRequest) (*MessageResult, error) {
	raw, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	if err := ws.WriteMessage(websocket.TextMessage, raw); err != nil {
		return nil, err
	}

	ws.SetReadDeadline(time.Now().Add(1 * time.Second))

	_, message, err := ws.ReadMessage()
	if err != nil {
		return nil, err
	}

	data := new(MessageResponse)

	if err = json.Unmarshal(message, data); err != nil {
		return nil, err
	}

	if data.ID != req.ID {
		return nil, errors.New("Mismatched ID's")
	}

	if data.Error != nil {
		return nil, errors.New(data.Error.Message)
	}

	return data.Result, nil
}

func (xo *xolib) getLogin() *MessageRequest {
	id := uuid.NewUUID()

	params := Params{}

	if xo.Config.Token != "" {
		params["token"] = xo.Config.Token
	} else {
		params["username"] = xo.Config.Username
		params["password"] = xo.Config.Password
	}

	return &MessageRequest{
		ID:      id.String(),
		Jsonrpc: "2.0",
		Params:  params,
	}
}

// Call is used to execute calls to the xolib api
func (xo *xolib) Call(req *MessageRequest) (*MessageResult, error) {
	ws, _, err := websocket.DefaultDialer.Dial("ws://"+xo.Config.Host+"/api/", nil)
	if err != nil {
		return nil, err
	}

	login := xo.getLogin()

	if _, err := rawCall(ws, login); err != nil {
		return nil, err
	}

	result, err := rawCall(ws, req)

	return result, err
}

// Init check if login works
func (xo *xolib) Init() error {
	ws, _, err := websocket.DefaultDialer.Dial("ws://"+xo.Config.Host+"/api/", nil)
	if err != nil {
		return err
	}

	login := xo.getLogin()

	if _, err := rawCall(ws, login); err != nil {
		return err
	}

	return nil
}

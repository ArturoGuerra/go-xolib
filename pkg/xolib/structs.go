package xolib

type (
	// Params are the websocket paramteres
	Params map[string]interface{}

	// MessageRequest is the message request format
	MessageRequest struct {
		ID      string `json:"id"`
		Method  string `json:"method"`
		Jsonrpc string `json:"jsonrpc"`
		Params  Params `json:"params"`
	}

	// MessageError represents an error message from xoa
	MessageError struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}

	// MessageResult is the result of a websocket response if successful
	MessageResult map[string]interface{}

	// MessageResponse is the websocket message response
	MessageResponse struct {
		ID      string         `json:"id"`
		Jsonrpc string         `json:"jsonrpc"`
		Error   *MessageError  `json:"error"`
		Result  *MessageResult `json:"result"`
	}
)

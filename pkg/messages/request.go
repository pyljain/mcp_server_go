package messages

type ClientRequest struct {
	JsonRPC string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
}

package messages

type Response struct {
	JsonRPC string                 `json:"jsonrpc"`
	Id      int                    `json:"id"`
	Result  map[string]interface{} `json:"result"`
}

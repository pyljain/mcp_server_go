package methods

import (
	"encoding/json"
	"fmt"
	"mcp_server/pkg/messages"
	"net/http"
)

func Initialize(w http.ResponseWriter, cr *messages.ClientRequest, eventChannel chan string) {

	w.Header().Add("Content-Type", "application/json")

	response := messages.Response{
		JsonRPC: "2.0",
		Id:      cr.Id,
		Result: map[string]interface{}{
			"protocolVersion": "2024-11-05",
			"capabilities": map[string]interface{}{
				"tools": map[string]interface{}{},
			},
			"serverInfo": map[string]interface{}{
				"name":    "hx",
				"version": "1.0.0",
			},
		},
	}

	eventData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error responding from server", http.StatusInternalServerError)
		return
	}

	eventChannel <- fmt.Sprintf("event: message\ndata: %s\n\n", eventData)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)
}

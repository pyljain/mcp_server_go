package methods

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mcp_server/pkg/messages"
	"mcp_server/pkg/tools"
	"net/http"
)

func CallTool(w http.ResponseWriter, cr *messages.ClientRequest, conn *sql.DB, eventChannel chan string) {

	allTools := map[string]tools.Tool{
		"query":      tools.NewQuery(conn),
		"listTables": tools.NewListTables(conn),
	}

	toolCalled, exists := allTools[cr.Params["name"].(string)]
	if !exists {
		http.Error(w, "Tool not found", http.StatusNotFound)
		return
	}

	res, err := toolCalled.Execute(cr.Params["arguments"].(map[string]interface{}))
	if err != nil {
		log.Printf("Error calling tool %s: %s", cr.Params["name"].(string), err)
		http.Error(w, "Error calling tool", http.StatusInternalServerError)
		return
	}

	response := messages.Response{
		JsonRPC: "2.0",
		Id:      cr.Id,
		Result: map[string]interface{}{
			"content": []map[string]string{
				{
					"type": "text",
					"text": res,
				},
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

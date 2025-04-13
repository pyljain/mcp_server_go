package methods

import (
	"encoding/json"
	"fmt"
	"mcp_server/pkg/messages"
	"net/http"
)

func ListTools(w http.ResponseWriter, cr *messages.ClientRequest, eventChannel chan string) {

	w.Header().Add("Content-Type", "application/json")

	response := messages.Response{
		JsonRPC: "2.0",
		Id:      cr.Id,
		Result: map[string]interface{}{
			"tools": []map[string]interface{}{
				{
					"name":        "query",
					"description": "Run queries against a Postgres database tables",
					"inputSchema": map[string]interface{}{
						"type": "object",
						"properties": map[string]interface{}{
							"query": map[string]string{
								"type":        "string",
								"description": "The fully formed select query to run",
							},
						},
					},
					"required": []string{"query"},
				},
				{
					"name":        "listTables",
					"description": "Use this tool to get a list of tables in the Postgres database",
					"inputSchema": map[string]interface{}{
						"type":       "object",
						"properties": map[string]interface{}{},
					},
					"required": []string{},
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

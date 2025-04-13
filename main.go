package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"mcp_server/pkg/messages"
	"mcp_server/pkg/methods"
	"net/http"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const expectedToken = "abcd"

func main() {

	mux := http.NewServeMux()

	eventChannels := make(map[string]chan string)

	db, err := sql.Open("sqlite3", "./mcp.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 7 {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		if authHeader[7:] != expectedToken {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
			return
		}

		sessionID := uuid.NewString()

		eventChannels[sessionID] = make(chan string)

		w.Header().Add("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "event: endpoint\ndata: %s/%s\n\n", "/messages", sessionID)
		flusher.Flush()

		for event := range eventChannels[sessionID] {
			fmt.Fprintf(w, "%s", event)
			flusher.Flush()
		}
	})

	mux.HandleFunc("/messages/{sessionID}", func(w http.ResponseWriter, r *http.Request) {

		sessionID := r.PathValue("sessionID")

		eventChannel := eventChannels[sessionID]

		cr := messages.ClientRequest{}
		err := json.NewDecoder(r.Body).Decode(&cr)
		if err != nil {
			http.Error(w, "Unable to read request", http.StatusBadRequest)
			return
		}

		log.Printf("Client request is %+v", cr)

		switch cr.Method {
		case "initialize":
			methods.Initialize(w, &cr, eventChannel)
		case "tools/list":
			methods.ListTools(w, &cr, eventChannel)
		case "tools/call":
			methods.CallTool(w, &cr, db, eventChannel)
		case "default":
			http.Error(w, "method not supported", http.StatusNotImplemented)
			return
		}
	})

	err = http.ListenAndServe(":8777", Logger(mux))
	if err != nil {
		log.Printf("Error starting server %s", err)
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request called %s at path %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

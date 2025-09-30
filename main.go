package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"webami/actions"
)

type Request struct {
	Action   string `json:"action"`
	Command  string `json:"command"`
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

func handleAMI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	resp, err := actions.SendJSONAction(req.Action, req.Command, req.Username, req.Secret)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	var output string
	if req.Command != "" {
		output = fmt.Sprintf("🚀 Command: %s\n%s\n%s",
			req.Command,
			strings.Repeat("=", 50),
			resp)
	} else {
		output = fmt.Sprintf("🔧 Action: %s\n%s\n%s",
			req.Action,
			strings.Repeat("=", 50),
			resp)
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(output))
}

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// API endpoint for AMI actions
	http.HandleFunc("/ami", handleAMI)

	fmt.Println("🚀 Server running at http://localhost:3000")
	fmt.Println("📁 Serving static files from ./public")
	fmt.Println("🔌 AMI endpoint: http://localhost:3000/ami")

	log.Fatal(http.ListenAndServe(":3000", nil))
}

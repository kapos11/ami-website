package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webami/actions"
)

type Request struct {
	Action   string `json:"action"`
	Command  string `json:"command"`
	Username string `json:"username"`
	Secret   string `json:"secret"`
}

func handleAMI(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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

	if req.Username == "" || req.Secret == "" {
		http.Error(w, "Username and secret are required", http.StatusBadRequest)
		return
	}

	resp, err := actions.SendJSONAction(req.Action, req.Command, req.Username, req.Secret)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(resp))
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

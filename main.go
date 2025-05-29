package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/api/user", apiUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Це сторінка про нас.")
}

func apiUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Name  string
		Email string
	}{Name: "Jane Doe", Email: "jane@example.com"}
	json.NewEncoder(w).Encode(data)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page!")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "guest"
		}
		fmt.Fprintf(w, "Hello, %s!\n", name)
		return
	}

	if r.Method == http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, `"error": "Request body read error"`, http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		response := struct {
			Status string `json:"status"`
			Data   string `json:"data"`
		}{Status: "received", Data: string(bodyBytes)}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, `"error": "Request serialization error"`, http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, `"json": "Method not allowed"`, http.StatusInternalServerError)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go Web!")
	})

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "guest"
		}
		fmt.Fprintf(w, "Hello, %s!\n", name)
	})

	http.HandleFunc("/api/user", apiUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := struct {
		Name  string
		Email string
	}{Name: "Jane Doe", Email: "jane@example.com"}
	json.NewEncoder(w).Encode(data)
}

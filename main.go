package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var users []User

func main() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/greet", greetHandler)
	http.HandleFunc("/api/user", apiUser)
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/500", internalServerErrorHandler)
	http.HandleFunc("/template", templateHandler)
	http.HandleFunc("/submit-form", submitFormHandler)
	http.HandleFunc("/api/users", usersAPIHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	users = []User{
		{ID: "001", Name: "Alice", Email: "alice@example.com"},
		{ID: "002", Name: "Bob", Email: "bob@example.com"},
		{ID: "003", Name: "John", Email: "john@example.com"},
	}

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

	http.Error(w, `"json": "Method not allowed"`, http.StatusMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusNotFound)
	//fmt.Fprintf(w, "Page not found")
	http.Error(w, "Page not found", http.StatusNotFound)
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusInternalServerError)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

type PageData struct {
	Title    string
	Greeting string
	Name     string
	Message  string
	Items    []string
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		//log.Fatal("Template parsing error: %v", err)
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:    "Динамічна сторінка",
		Greeting: "Привіт, світе шаблонів!",
		Name:     "Go Розробник",
		Items:    []string{"Елемент 1", "Елемент 2", "Елемент 3"},
	}
	tmpl.Execute(w, data)
}

func submitFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Form parse error", http.StatusInternalServerError)
		return
	}
	username := r.Form.Get("username")
	comment := r.Form.Get("comment")

	fmt.Fprintf(w, "Received data: Name: %s, Comment: %s", username, comment)
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var usersMutex sync.Mutex

func usersAPIHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(users)
		if err != nil {
			http.Error(w, "Failed to encode users to JSON", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		w.Header().Set("Content-Type", "application/json")
		var newUser User
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		newUID, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, "Failed to generate user ID", http.StatusInternalServerError)
			return
		}
		newUser.ID = newUID.String()

		usersMutex.Lock()
		defer usersMutex.Unlock()

		users = append(users, newUser)

		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(newUser)
		if err != nil {
			http.Error(w, "Failed to encode new user", http.StatusInternalServerError)
		}
		return

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}

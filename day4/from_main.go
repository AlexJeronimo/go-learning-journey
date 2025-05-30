package day4

/*

const filename = "my_test_file.txt"
	contents := []string{"Hello Go!", " this is my first file."}

	for _, content := range contents {
		day4.AppendToFile(filename, content)
	}

	data, err := day4.ReadFromFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}

*/

/*
	day4.CopyFile("my_test_file.txt", "my_new_text_file.txt")
	day4.ReadFromFile("my_new_text_file.txt")
*/

/*
dir, _ := os.Getwd()
	//fmt.Println(dir)
	files, err := day4.ListFiles(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
*/

/*
u := day4.User{Name: "Alice", Email: "alice@example.com", Age: 30, IsActive: true}
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	jsonBytes := string(data)
	fmt.Println(jsonBytes)

	jsonString := `{"name":"Bob","email":"bob@example.com","age":25,"is_active":false}`

	var newUser day4.User

	json.Unmarshal([]byte(jsonString), &newUser)

	fmt.Println(newUser)
*/

/*
users := []day4.User{
		{Name: "Alice", Email: "alice@example.com", Age: 30, IsActive: true},
		{Name: "Bob", Email: "bob@example.com", Age: 25, IsActive: false},
		{Name: "John", Email: "john@example.com", Age: 31, IsActive: true},
	}

	jsonBytes, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("users.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	var loadUsers []day4.User

	data, err := os.ReadFile("users.json")
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &loadUsers)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(loadUsers)
*/

/*
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
	http.HandleFunc("/404", notFoundHandler)
	http.HandleFunc("/500", internalServerErrorHandler)

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

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusNotFound)
	//fmt.Fprintf(w, "Page not found")
	http.Error(w, "Page not found", http.StatusNotFound)
}

func internalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(http.StatusInternalServerError)
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

*/

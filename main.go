package main

import (
	"encoding/json"
	"fmt"
	"glp/day4"
)

func main() {
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

}

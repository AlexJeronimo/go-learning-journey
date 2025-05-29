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

}

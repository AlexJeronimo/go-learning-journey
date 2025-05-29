package day4

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email,omitempty"`
	Age      int    `json:"age"`
	IsActive bool   `json:"is_active"`
}

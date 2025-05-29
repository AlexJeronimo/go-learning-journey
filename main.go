package main

import (
	"glp/day4"
)

func main() {
	day4.CopyFile("my_test_file.txt", "my_new_text_file.txt")
	day4.ReadFromFile("my_new_text_file.txt")
}

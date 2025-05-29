package main

import (
	"fmt"
	"glp/day4"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	//fmt.Println(dir)
	files, err := day4.ListFiles(dir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}

package main

import (
	"errors"
	"fmt"
	"glp/calculator"
)

func main() {
	defer fmt.Println("Closing resource...")
	add := calculator.Addition{}
	res, _ := add.Apply(8, 14)
	fmt.Println(res)

	dev := calculator.Division{}
	res, err := dev.Apply(16, 0)
	if err != nil {
		if val, ok := err.(*calculator.DevideByZeroError); ok {
			fmt.Println("Our type Error")
			fmt.Println(val.Message)
		} else {
			fmt.Println("Not our type Error")
		}
	} else {
		fmt.Println(res)
	}

	fmt.Println(errors.Is(err, calculator.ErrDivideByZero))
}

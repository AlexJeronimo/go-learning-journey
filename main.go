package main

import (
	"fmt"
	"glp/calculator"
	"glp/utils"
)

func main() {
	res, err := calculator.Calculate(2, 5, "-")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", res)
	}
	pi := calculator.GetPi()
	fmt.Println(pi)

	fmt.Println(utils.IsEven(8))
	fmt.Println(utils.IsEven(-5))

}

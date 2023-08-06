package main

import (
	"fmt"
)

type data struct {
	op1      int
	op2      int
	operator int
}

func main() {
	var str string
	_, err := fmt.Scanln(&str)
	if err != nil {
		return
	}

	res := parser(str)

}

func parser(str string) data {

}

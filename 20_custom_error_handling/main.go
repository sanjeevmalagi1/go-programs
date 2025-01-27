package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("DivideByZeroError")
	}

	return a / b, nil
}

func main() {

	res, err := Divide(5, 2)
	if err != nil {
		fmt.Printf("%s", err)
	}

	res, err = Divide(5, 0)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%d", res)

}

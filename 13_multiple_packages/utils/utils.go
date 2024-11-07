package utils

import "fmt"

func privateFunc() {
	fmt.Println("Private")
}

func PublicFunc() {
	fmt.Println("Public")
}

func RunAnotherFunc() {
	anotherFunc()
}

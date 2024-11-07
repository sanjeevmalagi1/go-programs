package main

import (
	"fmt"
)

func PrintFormat[MyType any](arg MyType) string {

	fmt.Printf("%T \n", arg)

	return fmt.Sprintf("%T", arg)
}

func main() {
	fmt.Println("Start")
	PrintFormat("Sanjeev")
	PrintFormat(15)
	return
}

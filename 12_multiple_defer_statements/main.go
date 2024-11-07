package main

import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("1st Defer")

	defer fmt.Println("2nd Defer")

	return
}

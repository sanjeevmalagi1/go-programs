package main

import (
	"fmt"

	uts "github.com/sanjeevmalagi1/go_programs/13_multiple_packages/utils"
)

func main() {
	fmt.Println("Start")
	uts.PublicFunc()
	uts.RunAnotherFunc()
	return
}

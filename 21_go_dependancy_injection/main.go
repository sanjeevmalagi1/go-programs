package main

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(string)
}

type FileLogger struct {
}

type ConsoleLogger struct {
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(message)

	if err != nil {
		panic(err)
	}

	file.Close()
}

func (f ConsoleLogger) Log(message string) {
	fmt.Printf("%s", message)
}

func Log(logger Logger, message string) {
	logger.Log(message)
}

func main() {

	fileLogger := FileLogger{}
	consoleLogger := ConsoleLogger{}

	Log(fileLogger, "Hi")
	Log(consoleLogger, "Hi")

}

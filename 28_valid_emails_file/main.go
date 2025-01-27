package main

import (
	"os"
	"regexp"
	"strings"
)

var validEmail = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)

func ReadEmailsFile() string {
	fileName := "emails.txt"

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func WriteValidEmailsFile(valid_emails []string) {
	file, err := os.Create("valid_emails.txt")
	if err != nil {
		panic(err)
	}

	var fileOutput string = ""
	for _, email := range valid_emails {
		fileOutput += email + "\n"
	}

	_, e := file.WriteString(fileOutput)
	if e != nil {
		panic(e)
	}

}

func filterValidEmails(emails []string) []string {
	var filteredEmails []string
	for _, email := range emails {
		if isValidEmail(email) {
			filteredEmails = append(filteredEmails, email)
		}
	}

	return filteredEmails
}

func isValidEmail(email string) bool {
	return validEmail.MatchString(email)
}

func main() {
	data := ReadEmailsFile()
	emails := strings.Split(data, "\n")

	filteredEmails := filterValidEmails(emails)

	WriteValidEmailsFile(filteredEmails)
}

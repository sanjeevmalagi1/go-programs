package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

func main() {
	employees := []Employee{
		{ID: 1, Name: "John Doe", Age: 30, Salary: 50000},
		{ID: 3, Name: "Zmily Davis", Age: 35, Salary: 55000},
		{ID: 2, Name: "Aane Smith", Age: 25, Salary: 55000},
		{ID: 4, Name: "Michael Brown", Age: 40, Salary: 70000},
	}

	sort.Slice(employees, func(i, j int) bool {
		if employees[i].Salary == employees[j].Salary {
			return employees[i].Name < employees[j].Name
		}

		return employees[i].Salary < employees[j].Salary
	})

	fmt.Println(employees)

}

package entity

import (
	"strconv"
)

type Person struct {
	Name string
	CPF  string
	Age  int
}

func PersonToStringSlice(person Person) []string {
	return []string{person.Name, person.CPF, strconv.Itoa(person.Age)}
}

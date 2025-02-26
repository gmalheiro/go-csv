package controller

import (
	"encoding/csv"
	"os"

	"gmalheiro.com/goCSV/entity"
)

func WriteCSV() error {
	person := entity.Person{
		Name: "Gabs",
		CPF:  "000.000.000-00",
		Age:  21,
	}

	file, err := os.OpenFile("./data/hr.csv", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	/*
		os.O_APPEND -> Add new lines without deleting the existing ones
		os.O_WRONLY -> Allows only writing
		os.O_CREATE -> Creates the file if it does not exist
		0644 -> File permissions (reading and writing for the owner, and reading for others)
	*/

	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Escrever a nova linha no CSV
	err = writer.Write(entity.PersonToStringSlice(person))
	if err != nil {
		return err
	}

	return nil
}

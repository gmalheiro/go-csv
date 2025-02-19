package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"gmalheiro.com/goCSV/model"
)

func main() {
	file, err := os.Open("./data/hr.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()

	if err != nil {
		panic(err)
	}

	people := make([]model.Person, len(records))

	for i, record := range records {
		age, _ := strconv.Atoi(record[2])
		person := model.Person{
			Name: record[0],
			CPF:  record[1],
			Age:  age,
		}

		people[i] = person
	}

	fmt.Println(people)
}

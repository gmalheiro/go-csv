package controller

import (
	"encoding/csv"
	"os"
)

func WriteCSV() {
	file, err := os.Create("subscribers.csv")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	records := [][]string{
		{"Gabs Malheiro", "gabsmalheiro@email.com"},
		{"Teste teste", "teste@email.com"},
	}
  

  writer := csv.NewWriter(file)

  defer writer.Flush()

  for _, record := range records {
    if err := writer.Write(record); err != nil {
      panic(err)
    }
  }
}

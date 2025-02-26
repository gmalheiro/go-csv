package main

import (
	"fmt"

	"gmalheiro.com/goCSV/controller"
)

func main() {
	people := controller.ReadCSV()
	fmt.Println(*people)
	controller.WriteCSV()
}

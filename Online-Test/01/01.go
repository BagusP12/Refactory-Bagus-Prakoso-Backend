package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type Gadgets struct {
	Name     string `json:NAME`
	Category string `json:CATEGORY`
	Price    string `json:PRICE`
}

func main() {
	csv_file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csv_file.Close()
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var gdts Gadgets
	var gadgets []Gadgets
	for _, rec := range records {
		gdts.Name = rec[0]
		gdts.Category = rec[1]
		gdts.Price = rec[2]

		gadgets = append(gadgets, gdts)
	}
	json_data, err := json.Marshal(gadgets)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(json_data))

	json_file, err := os.Create("data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	json_file.Write(json_data)
	json_file.Close()
}

package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadFile() ([][]string, error) {
	f, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	return csvReader.ReadAll()
}

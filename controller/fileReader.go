package controller

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"
)

func readFromCsvFile(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to open file", path, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as csv", path, err)
	}

	return records
}

func ReadFromFile(path string) []byte {
	fileData, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return fileData
}
package controller

import (
	"fmt"
	"snapp-trip/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
	err error
)

func ConnectToPostgre() {
	// dbURL := "postgress://pg:Pouya@Postgre@localhost:5432"
	dsn := "host=localhost user=postgres password=****** dbname=snapp-trip port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic (err)
	}
}

func InitializeTables() {
	importAirlinesToTable()
	importAgenciesToTable()
	importCitiesToTable()
	importSuppliersToTable()
}

func importAirlinesToTable() {
	db.AutoMigrate(&model.Airline{})

	path := "./resources/airline.csv"
	airlineString := readFromCsvFile(path)
	for _, array := range airlineString {
		airline := model.Airline{Code: array[0], EnName: array[1], FaName: array[2]}
		fmt.Println(airline)
		db.Create(&airline)
	}
}

func importAgenciesToTable() {
	db.AutoMigrate(&model.Agency{})

	path := "./resources/agency.csv"
	agencyArray := readFromCsvFile(path)
	for _, array := range agencyArray {
		agency := model.Agency{Name: array[2]}
		db.Create(&agency)
	}
}

func importCitiesToTable() {
	db.AutoMigrate(&model.City{})

	path := "./resources/city.csv"
	citiesCsv := readFromCsvFile(path)
	for _, array := range citiesCsv {
		city := model.City{Code: array[2], FaName: array[4]}
		db.Create(&city)
	}
}

func importSuppliersToTable() {
	db.AutoMigrate(&model.Supplier{})

	path := "./resources/supplier.csv"
	pathesCsv := readFromCsvFile(path)
	for _, array := range pathesCsv {
		supplier := model.Supplier{EnName: array[2], FaName: array[4]}
		db.Create(&supplier)
	}
}
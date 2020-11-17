package main

import (
	"github.com/cmorales95/golang-db-gorm/models"
	"github.com/cmorales95/golang-db-gorm/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	_ = storage.DB().AutoMigrate(
		&models.Product{},
		&models.InvoiceHeader{},
		&models.InvoiceItem{},
	)

	invoice := models.InvoiceHeader{
		Client: "Cristian Morales",
		InvoiceItems: []models.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
}


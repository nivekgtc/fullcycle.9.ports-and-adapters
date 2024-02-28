package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/nivekgtc/hexagonal/adapters/db"
	"github.com/nivekgtc/hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")

	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("Product Example", 30)

	productService.Enable(product)

}
package main

import (
	"database/sql"
)

func getSupremeProducts() []ProductSupremeJson {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM supreme")
	var products []ProductSupremeJson
	for rows.Next(){
		var product ProductSupremeJson
		rows.Scan(&product.Name, &product.Category, &product.Colors)
		products = append(products, product)
	}
	return products
}
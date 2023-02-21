package main

import (
	"database/sql"
	"fmt"
)

func addProductPalace(products []string, category string) {
	database, err := sql.Open("sqlite3", "DB/Bot.db")
	if err != nil {
		fmt.Println(err)
	   }
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS palace (productName TEXT, category TEXT)")
	statement.Exec()
	for _ ,product := range products {
		statementInsert, _ := database.Prepare("INSERT INTO palace (productName, category) VALUES (?, ?)")
		statementInsert.Exec(product, category)
	}
}

func getPalaceProducts() []ProductPalaceJson {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM palace")
	var products []ProductPalaceJson
	for rows.Next(){
		var product ProductPalaceJson
		rows.Scan(&product.Name, &product.Category)
		products = append(products, product)
	}
	return products
}
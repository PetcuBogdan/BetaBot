package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func saveCard (card CardJson) {
	database, err := sql.Open("sqlite3", "DB/Bot.db")	
	if err != nil {
		fmt.Println(err)
	   }
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS card (id INTEGER PRIMARY KEY, nameCard TEXT, cardNumber TEXT, expirationDate TEXT, cvv TEXT)")
	statement.Exec()
	statementInsert, _ := database.Prepare("INSERT INTO card (id, nameCard, cardNumber, expirationDate, cvv) VALUES (?, ?, ?, ?, ?)")
	statementInsert.Exec(card.Id, card.NameCard, card.CardNumber, card.ExpirationDate, card.Cvv)
}

func editCard (card CardJson) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("INSERT INTO card (nameCard, cardNumber, expirationDate, cvv) WHERE id = ? VALUES (?, ?, ?, ?) ")
	statement.Exec(card.Id, card.NameCard, card.CardNumber, card.ExpirationDate, card.Cvv)
}

func deleteCard (id int) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("DELETE FROM card WHERE id=?")
	statement.Exec(id)
}

func getCards() []CardJson {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM card")
	var cards []CardJson
	for rows.Next(){
		var card CardJson
		rows.Scan(&card.Id, &card.NameCard, &card.CardNumber, &card.ExpirationDate, &card.Cvv)
		cards = append(cards, card)
	}
	return cards
}

func getSpecCard(id int) Card {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM card")
	var card Card
	for rows.Next(){
		var card2 Card
		rows.Scan(&card2.Id, &card2.NameCard, &card2.CardNumber, &card2.ExpirationDate, &card2.Cvv)
		if card2.Id == id{
			card = card2
		}
	}
	return card
}

func saveProfile (profile ProfileJson) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS profile (id INTEGER PRIMARY KEY, lname TEXT, fname TEXT, email TEXT, phone TEXT, address TEXT, address2 TEXT, postcode TEXT, city TEXT, county TEXT, country TEXT)")
	statement.Exec()
	statementInsert, _ := database.Prepare("INSERT INTO profile (id, lname, fname, email, phone, address, address2, postcode, city, county, country) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	statementInsert.Exec(profile.Id, profile.Lname, profile.Fname, profile.Email, profile.Phone, profile.Address, profile.Address2, profile.Postcode, profile.City, profile.County, profile.Country)
}

func editProfile(profile ProfileJson) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	statementEdit, _ := database.Prepare("INSERT INTO profile (lname, fname, email, phone, address, address2, postcode, city, county, country) WHERE id = ? VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	statementEdit.Exec(profile.Id, profile.Lname, profile.Fname, profile.Email, profile.Phone, profile.Address, profile.Address2, profile.Postcode, profile.City, profile.County, profile.Country)
}

func deleteProfile(id int) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("DELETE FROM profile WHERE id=? ")
	statement.Exec(id)
}

func getProfiles() []ProfileJson{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM profile")
	var profiles []ProfileJson
	for rows.Next(){
		var profile ProfileJson
		rows.Scan(&profile.Id, &profile.Lname, &profile.Fname, &profile.Email, &profile.Phone, &profile.Address, &profile.Address2, &profile.Postcode, &profile.City, &profile.County, &profile.Country)
		profiles = append(profiles, profile)
	}
	return profiles
}

func getSpecProfile(id int) Profile{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM profile")
	var profile Profile
	for rows.Next(){
		var profile2 Profile
		rows.Scan(&profile2.Id, &profile2.Lname, &profile2.Fname, &profile2.Email, &profile2.Phone, &profile2.Address, &profile2.Address2, &profile2.Postcode, &profile2.City, &profile2.County, &profile2.Country)
		if profile2.Id == id{
			profile = profile2
		}
	}
	return profile
}

func saveTask (task TaskJson) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS task (id INTEGER PRIMARY KEY, name TEXT, shop TEXT, productName TEXT, category TEXT, size TEXT, color TEXT, profile INTEGER, card INTEGER, status TEXT)")
	statement.Exec()
	statementInsert, _ := database.Prepare("INSERT INTO task (id, name, shop, productName, category, size, color, profile, card, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	statementInsert.Exec(task.Id, task.Name, task.Shop, task.ProductName, task.Category, task.Size, task.Color, task.Profile, task.Card, task.Status)
}

func editTask (task TaskJson) {
 	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
 	statementInsert, _ := database.Prepare("INSERT INTO task (name, shop, productName, category, size, color, profile, card, status) WHERE id = ? VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
 	statementInsert.Exec(task.Id, task.Name, task.Shop, task.ProductName, task.Category, task.Size, task.Color, task.Profile, task.Card, task.Status)
 }

func deleteTask (id int) {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")	
	statement, _ := database.Prepare("DELETE FROM task WHERE id=?")
	statement.Exec(id)
}

func getTasks() []TaskJson {
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM task")
	var tasks []TaskJson
	for rows.Next(){
		var task TaskJson
		rows.Scan(&task.Id, &task.Name, &task.Shop, &task.ProductName, &task.Category, &task.Size, &task.Color, &task.Profile, &task.Card, &task.Status)
		tasks = append(tasks, task)
	}
	return tasks
}

func getSpecTask(id int) Task{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM task")
	var task Task
	for rows.Next(){
		var task2 Task
		rows.Scan(&task2.Id, &task2.Name, &task2.Shop, &task2.ProductName, &task2.Category, &task2.Size, &task2.Color, &task2.Profile, &task2.Card, &task2.Status)
		if task2.Id == id{
			task = task2
		}
	}
	return task
}

func getIdCard() int{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM card")
	id := 0
	for rows.Next(){
		var card CardJson
		rows.Scan(&card.Id, &card.NameCard, &card.CardNumber, &card.ExpirationDate, &card.Cvv)
		if card.Id > id{
			id = card.Id
		}
	}
	return id
}

func getIdProfile() int{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM profile")
	id := 0
	for rows.Next(){
		var profile ProfileJson
		rows.Scan(&profile.Id, &profile.Lname, &profile.Fname, &profile.Email, &profile.Phone, &profile.Address, &profile.Address2, &profile.Postcode, &profile.City, &profile.County, &profile.Country)
		if profile.Id > id {
			id = profile.Id
		}
	}
	return id
}

func getIdTask() int{
	database, _ := sql.Open("sqlite3", "DB/Bot.db")
	rows, _ := database.Query("SELECT * FROM task")
	id := 0
	for rows.Next(){
		var task Task
		rows.Scan(&task.Id, &task.Name, &task.Shop, &task.ProductName, &task.Category, &task.Size, &task.Color, &task.Profile, &task.Card, &task.Status)
		if task.Id > id{
			id = task.Id
		}
	}
	return id
}


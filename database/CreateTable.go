package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTable(dbFilePath string) {
	// Открытие базы данных SQLite
	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		fmt.Println("Ошибка открытия базы данных:", err)
		return
	}
	defer db.Close()

	// Отправка SQL запроса на создание таблицы
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		title TEXT,
		comment TEXT,
		repeat TEXT
	);
    `
	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println("Ошибка выполнения SQL запроса:", err)
		return
	}

	fmt.Println("Таблица успешно создана.")

}

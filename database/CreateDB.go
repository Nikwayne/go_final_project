package database

import (
	"log"
	"os"
)

// CreateDatabase создает файл базы данных SQLite без создания таблицы
func CreateDatabase(dbFilePath string) error {
	// Проверяем существование файла базы данных
	if _, err := os.Stat(dbFilePath); os.IsNotExist(err) {
		// Если файл не существует, создаем его
		_, err := os.Create(dbFilePath)
		if err != nil {
			return err
		}

		log.Printf("Файл базы данных %s создан.", dbFilePath)

	} else {
		log.Printf("Файл базы данных %s уже существует.", dbFilePath)
	}
	return nil
}

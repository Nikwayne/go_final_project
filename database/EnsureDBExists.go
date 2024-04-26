package database

import (
	"log"           // Импортируем пакет для логирования ошибок
	"os"            // Импортируем пакет для работы с операционной системой
	"path/filepath" // Импортируем пакет для работы с путями к файлам
)

// CreateDB - функция для создания базы данных
func EnsureDBExists() {
	DataBasePath := "database/scheduler.db"

	// Получаем путь к исполняемому файлу
	appPath, err := os.Executable()
	if err != nil { // Если произошла ошибка при получении пути к файлу
		log.Fatal(err) // Записываем ошибку в лог и завершаем программу
	}

	// Создаем путь к файлу базы данных
	dbFile := filepath.Join(filepath.Dir(appPath), DataBasePath)

	// Проверяем существование файла базы данных
	_, err = os.Stat(dbFile)

	if err != nil { // Если произошла ошибка при проверке существования файла
		CreateDatabase(DataBasePath)
	}
	CreateTable(DataBasePath)
}

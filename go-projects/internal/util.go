package internal

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ReadFile(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error opening csv file: %s", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error reading csv file: %s", err))
	}
	return records, nil
}

func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to connect to DB: %s", err))
	}
	fmt.Println("Database created successfully...")
	return db, nil
}

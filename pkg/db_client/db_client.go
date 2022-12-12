package db_client

import (
	"OnboardingExercise/config"
	"database/sql"
	"fmt"
)

// NewDBConnection returns a new db connection
func NewDBConnection(dbConnection config.DBConnection) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnection.Host, dbConnection.Port, dbConnection.User, dbConnection.Password, dbConnection.DBName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	fmt.Println("Connected to the DB successfully")
	return db, nil
}

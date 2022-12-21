package db_client

import (
	"OnboardingExercise/cmd/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// NewDBClient returns a new db client object
func NewDBClient(dbConnectionInfo config.DBConnectionInfo) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConnectionInfo.Host, dbConnectionInfo.Port, dbConnectionInfo.User, dbConnectionInfo.Password, dbConnectionInfo.DBName)

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

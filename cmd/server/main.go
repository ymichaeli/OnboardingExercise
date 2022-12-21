package main

import (
	"OnboardingExercise/cmd/config"
	"OnboardingExercise/pkg/api"
	"OnboardingExercise/pkg/db_client"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	configFile, err := config.LoadConfig()
	if err != nil {
		panic(errors.Wrap(err, "config loading failed"))
	}

	db, err := db_client.NewDBClient(configFile.DBConfig)
	if err != nil {
		panic(errors.Wrap(err, "db initialized failed"))
	}

	server := api.NewServer(db)
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.Wrap(err, "Couldn't create server")))
	}

	if err := server.Start(configFile.ServerInfo); err != nil {
		panic(errors.Wrap(err, "Couldn't start server"))
	}
}

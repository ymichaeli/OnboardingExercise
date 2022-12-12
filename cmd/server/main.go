package main

import (
	"OnboardingExercise/pkg/api"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	server, err := api.NewServer()
	if err != nil {
		panic(fmt.Sprintf("%+v", errors.Wrap(err, "Couldn't create server")))
	}

	domain := "localhost"
	port := 8080

	if err := server.Start(domain, port); err != nil {
		panic(errors.Wrap(err, "Couldn't start server"))
	}
}

package main

import (
	"OnboardingExercise/pkg/api"
	"github.com/pkg/errors"
)

func main() {
	server := api.NewServer()
	domain := "localhost"
	port := 8080

	if err := server.Start(domain, port); err != nil {
		panic(errors.Wrap(err, "Couldn't start server"))
	}
}

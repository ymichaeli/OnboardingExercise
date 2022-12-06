package main

import (
	"OnboardingExercise/pkg/api"
	"github.com/pkg/errors"
)

func main() {
	server := api.NewServer()

	if err := server.Start(); err != nil {
		panic(errors.WithStack(errors.Wrap(err, "Couldn't start server")))
	}
}

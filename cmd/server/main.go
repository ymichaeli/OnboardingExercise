package main

import (
	"OnboardingExercise/pkg/api"
)

func main() {
	server, err := api.NewServer()

	if err != nil {
		panic("Couldn't create server")
	}

	if err = server.Start(); err != nil {
		panic("Couldn't start server")
	}
}

package lifecycle_service

import (
	lifecycle_repository "OnboardingExercise/pkg/repository/lifecycle"
)

type Service struct {
	repository lifecycle_repository.Repository
}

func NewService(repository lifecycle_repository.Repository) Service {
	return Service{repository: repository}
}

func (service Service) IsAlive() bool {
	return true
}

func (service Service) IsReady() bool {
	return service.repository.IsReady()
}

type LifecycleService interface {
	IsAlive() bool
	IsReady() bool
}

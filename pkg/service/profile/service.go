package profile_service

import (
	"OnboardingExercise/pkg/repository/profile"
	"github.com/google/uuid"
)

type Service struct {
	repository profile_repository.Repository
}

func NewService(repository profile_repository.Repository) Service {
	return Service{repository: repository}
}

func (service *Service) GetAllProfiles() []profile_repository.Profile {
	return service.repository.GetAllProfiles()
}

func (service *Service) GetProfileByUserID(userId string) (profile_repository.Profile, error) {
	return service.repository.GetProfileByUserID(userId)
}

func (service *Service) CreateProfile(newProfile profile_repository.Profile) profile_repository.Profile {
	newProfile.UserId = uuid.New().String()
	service.repository.CreateProfile(newProfile)

	return newProfile
}

func (service *Service) UpdateProfile(updatedProfile profile_repository.Profile) error {
	return service.repository.UpdateProfile(updatedProfile)
}

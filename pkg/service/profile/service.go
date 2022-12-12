package profile_service

import (
	"OnboardingExercise/pkg/api/models"
	"OnboardingExercise/pkg/repository/profile"
	"github.com/google/uuid"
)

// Service implements CRUD functions for profiles connection between the handler and the repository
type Service struct {
	repository profile_repository.Repository
}

func NewService(repository profile_repository.Repository) Service {
	return Service{repository: repository}
}

func (service *Service) GetAllProfiles() ([]profile_repository.Profile, error) {
	return service.repository.GetAllProfiles()
}

func (service *Service) GetProfileByUserID(userId string) (profile_repository.Profile, error) {
	return service.repository.GetProfileByUserID(userId)
}

func (service *Service) CreateProfile(newProfile api_models.Profile) (profile_repository.Profile, error) {
	profileToCreate := profile_repository.Profile{
		UserId:        uuid.New().String(),
		UserName:      newProfile.UserName,
		FullName:      newProfile.FullName,
		Bio:           newProfile.Bio,
		ProfilePicURL: newProfile.ProfilePicURL,
	}

	return profileToCreate, service.repository.CreateProfile(profileToCreate)
}

func (service *Service) UpdateProfile(updatedProfile api_models.Profile, userId string) error {
	return service.repository.UpdateProfile(profile_repository.Profile{
		UserId:        userId,
		UserName:      updatedProfile.UserName,
		FullName:      updatedProfile.FullName,
		Bio:           updatedProfile.Bio,
		ProfilePicURL: updatedProfile.ProfilePicURL,
	})
}

package profile_repository

import (
	"OnboardingExercise/pkg/custom_errors"
	"fmt"
)

// Repository Note: I thought about another solution for the address changed case
// in a regular case the Repository won't contain something that should change
// it will contain a constant db client
// so maybe a better solution will be to extract the profiles slice from the repository and make it global?
type Repository struct {
	profiles []Profile // will be changed to db client
}

func NewDAL(profiles []Profile) Repository {
	return Repository{profiles: profiles}
}

func (repository *Repository) GetAllProfiles() []Profile {
	return repository.profiles
}

func (repository *Repository) GetProfileByUserID(userId string) (Profile, error) {
	for _, profile := range repository.GetAllProfiles() {
		if profile.UserId == userId {
			return profile, nil
		}
	}

	return Profile{}, custom_errors.NotFoundError{Err: fmt.Sprintf("user %s does not exist", userId)}
}

func (repository *Repository) CreateProfile(newProfile Profile) {
	repository.profiles = append(repository.profiles, newProfile)
}

func (repository *Repository) UpdateProfile(updatedProfile Profile) error {
	for i, profile := range repository.GetAllProfiles() {
		if profile.UserId == updatedProfile.UserId {
			repository.profiles[i] = updatedProfile
			return nil
		}
	}
	return custom_errors.NotFoundError{Err: fmt.Sprintf("user %s does not exist", updatedProfile.UserId)}
}

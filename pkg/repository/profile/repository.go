package profile_repository

import (
	"OnboardingExercise/pkg/custom_errors"
	"database/sql"
	"fmt"
	"github.com/kisielk/sqlstruct"
)

// Repository implements CRUD functions on profiles table
type Repository struct {
	client *sql.DB
}

func NewRepository(client *sql.DB) Repository {
	return Repository{client: client}
}

func (repository *Repository) GetAllProfiles() (profiles []Profile, err error) {
	rows, err := repository.client.Query(fmt.Sprintf("SELECT %s FROM profiles", sqlstruct.Columns(Profile{})))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var profile Profile
		err = sqlstruct.Scan(&profile, rows)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}

	return profiles, nil
}

func (repository *Repository) GetProfileByUserID(userId string) (Profile, error) {
	query := fmt.Sprintf(`SELECT %s FROM profiles where user_id = $1`, sqlstruct.Columns(Profile{}))
	rows, err := repository.client.Query(query, userId)
	if err != nil {
		return Profile{}, err
	}

	defer rows.Close()

	if rows.Next() {
		var profile Profile
		if err = sqlstruct.Scan(&profile, rows); err != nil {
			return Profile{}, err
		}
		return profile, nil
	}
	return Profile{}, custom_errors.NewNotFoundError(fmt.Sprintf("user %s does not exist", userId), "")
}

func (repository *Repository) CreateProfile(profile Profile) (string, error) {
	insertQuery := `insert into profiles (username, full_name, bio, profile_pic_url) VALUES ($1, $2, $3, $4) RETURNING user_id`
	var userId string
	if err := repository.client.QueryRow(insertQuery, profile.UserName, profile.FullName, profile.Bio, profile.ProfilePicURL).Scan(&userId); err != nil {
		return "", err
	}
	return userId, nil
}

func (repository *Repository) UpdateProfile(updatedProfile Profile) error {
	updateQuery := `update profiles set username = $1, full_name = $2, bio = $3, profile_pic_url = $4 where user_id = $5`
	res, err := repository.client.Exec(updateQuery, updatedProfile.UserName, updatedProfile.FullName, updatedProfile.Bio, updatedProfile.ProfilePicURL, updatedProfile.UserId)

	if err != nil {
		return err
	}

	// this section is a little problematic - in case rowAffected is failed there is no good way to return response to the client
	// because the update query succeeded, but we are not sure that the userId exists in the database
	if n, _ := res.RowsAffected(); n == 0 {
		return custom_errors.NewNotFoundError(fmt.Sprintf("user %s does not exist", updatedProfile.UserId), "")
	}

	return nil
}

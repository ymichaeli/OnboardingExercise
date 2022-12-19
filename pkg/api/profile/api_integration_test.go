package profile_api

import (
	"OnboardingExercise/config"
	api_models "OnboardingExercise/pkg/api/models"
	"OnboardingExercise/pkg/db_client"
	profile_repository "OnboardingExercise/pkg/repository/profile"
	profile_service "OnboardingExercise/pkg/service/profile"
	"bytes"
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/qustavo/dotsql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type APIIntegrationTestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Handler  Handler
	DB       *sql.DB
}

func (s *APIIntegrationTestSuite) SetupSuite() {
	configFile, err := config.LoadConfig()
	configFile.DBConfig.DBName = "onboarding_db_test"
	if err != nil {
		panic(errors.Wrap(err, "config loading failed"))
	}

	db, err := db_client.NewDBClient(configFile.DBConfig)
	if err != nil {
		panic(errors.Wrap(err, "db initialized failed"))
	}
	s.DB = db
}

func (s *APIIntegrationTestSuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Context, _ = gin.CreateTestContext(s.Recorder)
	s.Handler = NewHandler(profile_service.NewService(profile_repository.NewRepository(s.DB)))

	s.executeSqlFile("./schema/create_profile_table.sql", "create-profile-table")
}

func (s *APIIntegrationTestSuite) TearDownTest() {
	s.executeSqlFile("./schema/delete_profile_table.sql", "delete-profile-table")
}

func (s *APIIntegrationTestSuite) executeSqlFile(file string, section string) {
	dot, err := dotsql.LoadFromFile(file)
	if err != nil {
		panic("couldn't create profile table")
	}
	if _, err := dot.Exec(s.DB, section); err != nil {
		panic("couldn't create profile table")
	}
}

func (s *APIIntegrationTestSuite) TestCreateProfile() {
	// Arrange
	profile := api_models.Profile{
		UserName:      "newUserName",
		FullName:      "newFullName",
		Bio:           "newBio",
		ProfilePicURL: "https://host/new-url.png",
	}
	marshaledProfile, err := json.Marshal(profile)
	if err != nil {
		panic("couldn't marshal profile")
	}

	req, _ := http.NewRequest(http.MethodPost, "/profiles", bytes.NewBuffer(marshaledProfile))
	req.Header.Set("Content-Type", "application/json")
	s.Context.Request = req

	// Act
	s.Handler.CreateProfile(s.Context)

	// Assert Response
	assert.Equal(s.T(), http.StatusCreated, s.Recorder.Code)

	// Assert Insertion
	var returnedProfile api_models.Profile
	if json.Unmarshal(s.Recorder.Body.Bytes(), &returnedProfile) != nil {
		panic("Couldn't unmarshal returned object")
	}
	dbProfile, err := s.Handler.service.GetProfileByUserID(returnedProfile.UserId)
	if err != nil {
		panic("couldn't get user id")
	}

	assert.Equal(s.T(), true, reflect.DeepEqual(api_models.Profile(dbProfile), returnedProfile))
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(APIIntegrationTestSuite))
}

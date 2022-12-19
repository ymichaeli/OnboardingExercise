package profile_service

import (
	"OnboardingExercise/pkg/api/lifecycle"
	"OnboardingExercise/pkg/service/lifecycle/lifecycle_service_mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type APITestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
	Service  *lifecycle_service_mock.Service
}

func (s *APITestSuite) SetupTest() {
	s.Recorder = httptest.NewRecorder()
	s.Context, _ = gin.CreateTestContext(s.Recorder)
	s.Service = new(lifecycle_service_mock.Service)
}

func (s *APITestSuite) TestServerIsReady() {
	// Arrange
	req, _ := http.NewRequest(http.MethodGet, "/is-ready", nil)
	s.Context.Request = req
	s.Service.Mock.On("IsReady", mock.Anything).Return(true)
	handler := lifecycle.NewHandler(s.Service)

	// Act
	handler.Readiness(s.Context)

	// Assert
	assert.Equal(s.T(), http.StatusOK, s.Recorder.Code)
	assert.Equal(s.T(), `{"status":"ok"}`, s.Recorder.Body.String())
}

func (s *APITestSuite) TestServerIsNotReady() {
	// Arrange
	req, _ := http.NewRequest(http.MethodGet, "/is-ready", nil)
	s.Context.Request = req
	s.Service.Mock.On("IsReady", mock.Anything).Return(false)
	handler := lifecycle.NewHandler(s.Service)

	// Act
	handler.Readiness(s.Context)

	// Assert
	assert.Equal(s.T(), http.StatusServiceUnavailable, s.Recorder.Code)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}

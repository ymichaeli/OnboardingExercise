package lifecycle

import (
	"OnboardingExercise/pkg/service/lifecycle"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MiddlewareTestSuite struct {
	suite.Suite
	Recorder *httptest.ResponseRecorder
	Context  *gin.Context
}

func (suite *MiddlewareTestSuite) SetupTest() {
	suite.Recorder = httptest.NewRecorder()
	suite.Context, _ = gin.CreateTestContext(suite.Recorder)
}

func (suite *MiddlewareTestSuite) TestRegularError() {
	req, _ := http.NewRequest(http.MethodGet, "/is-alive", nil)

	suite.Context.Request = req

	handler := NewHandler(lifecycle_service.NewService())
	handler.health(suite.Context)

	assert.Equal(suite.T(), http.StatusOK, suite.Recorder.Code)
	assert.Equal(suite.T(), `{"status":"ok"}`, suite.Recorder.Body.String())
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

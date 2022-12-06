package middlewares

import (
	"OnboardingExercise/pkg/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	suite.Context.Errors[0] = suite.Context.Error(errors.New("some error"))
	ErrorHandler(suite.Context)

	assert.Equal(suite.T(), http.StatusInternalServerError, suite.Recorder.Result().StatusCode)
}

func (suite *MiddlewareTestSuite) TestNotFoundError() {
	suite.Context.Errors[0] = suite.Context.Error(custom_errors.NotFoundError{Err: "not found"})
	ErrorHandler(suite.Context)

	assert.Equal(suite.T(), http.StatusNotFound, suite.Recorder.Result().StatusCode)
}

func (suite *MiddlewareTestSuite) TestBadRequestError() {
	suite.Context.Errors[0] = suite.Context.Error(custom_errors.BadRequestError{Err: "bad request"})
	ErrorHandler(suite.Context)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.Recorder.Result().StatusCode)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}

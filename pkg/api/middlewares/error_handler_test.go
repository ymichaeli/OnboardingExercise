package middlewares

import (
	"OnboardingExercise/pkg/custom_errors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegularError(t *testing.T) {
	errorMessage := "some error"
	err := errors.New(errorMessage)

	status, _ := HandleError(err)

	assert.Equal(t, http.StatusInternalServerError, status)
}

func TestNotFoundError(t *testing.T) {
	publicMessage := "didn't found myResource"
	err := custom_errors.NewNotFoundError(publicMessage, "")

	status, _ := HandleError(err)

	assert.Equal(t, http.StatusNotFound, status)
}

func TestBadRequestError(t *testing.T) {
	publicMessage := "field Bio is required"
	err := custom_errors.NewBadRequestError(publicMessage, "")

	status, message := HandleError(err)

	assert.Equal(t, http.StatusBadRequest, status)
	assert.Equal(t, gin.H{"error": publicMessage}, message)
}

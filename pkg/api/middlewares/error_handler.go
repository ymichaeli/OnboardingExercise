package middlewares

import (
	"OnboardingExercise/pkg/custom_errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err // to simplify I assume there is only one Error

		// log the full error with stacktrace
		fmt.Printf("%+v\n", errors.WithStack(err))

		if errors.As(err, &custom_errors.NotFoundError{}) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if errors.As(err, &custom_errors.BadRequestError{}) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// in case of an unknown error return a general error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	}
}

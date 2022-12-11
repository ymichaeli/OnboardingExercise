package middlewares

import (
	"OnboardingExercise/pkg/custom_errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
)

// ErrorHandlerMiddleware returns a relevant response error to the client in case of an error
// using mapping from known custom error to status code and a message
func ErrorHandlerMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors[0].Err // to simplify I assume there is only one Error
		status, body := HandleError(err)

		c.JSON(status, body)
	}
}

func HandleError(err error) (int, interface{}) {

	// log the full error with stacktrace
	fmt.Printf("%+v\n", errors.WithStack(err))

	var notFoundError custom_errors.NotFoundError
	if errors.As(err, &notFoundError) {
		return http.StatusNotFound, gin.H{"error": notFoundError.DisplayMessage()}
	}

	var badRequestError custom_errors.BadRequestError
	if errors.As(err, &badRequestError) {
		return http.StatusBadRequest, gin.H{"error": badRequestError.DisplayMessage()}
	}

	return http.StatusInternalServerError, gin.H{"error": "Internal Server Error"}
}

package profile

import (
	"OnboardingExercise/pkg/custom_errors"
	"OnboardingExercise/pkg/repository/profile"
	"OnboardingExercise/pkg/service/profile"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service profile_service.Service
}

func NewHandler(service profile_service.Service) Handler {
	return Handler{service: service}
}

func (handler *Handler) GetAllProfiles(c *gin.Context) {
	profiles := handler.service.GetAllProfiles()

	c.JSON(http.StatusOK, profiles)
}

func (handler *Handler) GetProfileByUserID(c *gin.Context) {
	userId := c.Param("userId")
	userProfile, err := handler.service.GetProfileByUserID(userId)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (handler *Handler) CreateProfile(c *gin.Context) {
	var newProfile profile_repository.Profile
	if err := c.ShouldBind(&newProfile); err != nil {
		_ = c.Error(custom_errors.BadRequestError{Err: err.Error()})
		return
	}

	newProfile = handler.service.CreateProfile(newProfile)
	c.JSON(http.StatusCreated, newProfile)

}

func (handler *Handler) UpdateProfile(c *gin.Context) {
	var newProfile profile_repository.Profile
	if err := c.ShouldBind(&newProfile); err != nil {
		_ = c.Error(custom_errors.BadRequestError{Err: err.Error()})
		return
	}
	newProfile.UserId = c.Param("userId")

	if err := handler.service.UpdateProfile(newProfile); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

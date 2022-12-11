package profile_api

import (
	"OnboardingExercise/pkg/api/models"
	"OnboardingExercise/pkg/custom_errors"
	"OnboardingExercise/pkg/service/profile"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler implements profiles CRUD functions using profile_service.Service
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
	var params api_models.ProfileUserParams
	if err := c.ShouldBindUri(&params); err != nil {
		_ = c.Error(custom_errors.NewBadRequestError(err.Error(), ""))
	}
	userProfile, err := handler.service.GetProfileByUserID(params.UserId)

	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, userProfile)
}

func (handler *Handler) CreateProfile(c *gin.Context) {
	var newProfile api_models.Profile
	if err := c.ShouldBind(&newProfile); err != nil {
		_ = c.Error(custom_errors.NewBadRequestError(err.Error(), ""))
		return
	}

	createdProfile := handler.service.CreateProfile(newProfile)
	c.JSON(http.StatusCreated, createdProfile)

}

func (handler *Handler) UpdateProfile(c *gin.Context) {
	var newProfile api_models.Profile
	if err := c.ShouldBind(&newProfile); err != nil {
		_ = c.Error(custom_errors.NewBadRequestError(err.Error(), ""))
		return
	}

	var params api_models.ProfileUserParams
	if err := c.ShouldBindUri(&params); err != nil {
		_ = c.Error(custom_errors.NewBadRequestError(err.Error(), ""))
	}

	if err := handler.service.UpdateProfile(newProfile, params.UserId); err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

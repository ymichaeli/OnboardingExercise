package api_models

type ProfileUserParams struct {
	UserId string `uri:"userId" binding:"required"`
}

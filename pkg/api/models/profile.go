package api_models

type Profile struct {
	UserName      string `json:"userName" binding:"required"`
	FullName      string `json:"fullName" binding:"required"`
	Bio           string `json:"bio" binding:"required"`
	ProfilePicURL string `json:"profilePicURL" binding:"required"`
}

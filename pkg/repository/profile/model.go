package profile_repository

// Profile struct contains UserId which is not required because we generate it when the client wants to create a new user
// is it ok, or I should create a partial Profile struct without UserId?
type Profile struct {
	UserId        string `json:"userId"`
	UserName      string `json:"userName" binding:"required"`
	FullName      string `json:"fullName" binding:"required"`
	Bio           string `json:"bio" binding:"required"`
	ProfilePicURL string `json:"profilePicURL" binding:"required"`
}

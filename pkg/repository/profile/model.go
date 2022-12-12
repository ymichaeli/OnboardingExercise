package profile_repository

type Profile struct {
	UserId        string `sql:"user_id"`
	UserName      string `sql:"username"`
	FullName      string `sql:"full_name"`
	Bio           string `sql:"bio"`
	ProfilePicURL string `sql:"profile_pic_url"`
}

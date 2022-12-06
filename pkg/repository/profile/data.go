package profile_repository

// Profiles variable is going to be deprecated and the data will be fetched from db instead
var Profiles = []Profile{
	{
		UserId:        "uuid",
		UserName:      "yuvalmich",
		FullName:      "Yuval Michaeli",
		Bio:           "this is my beautiful bio",
		ProfilePicURL: "http://somehost/something.png",
	},
}

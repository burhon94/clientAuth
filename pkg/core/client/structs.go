package client

type SignIn struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type NewClientStruct struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	EMail      string `json:"e_mail"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
}

type EditClientPass struct {
	Id      int64  `json:"id"`
	OldPass string `json:"old_pass"`
	NewPass string `json:"new_pass"`
}

type EditClientAvatar struct {
	Id        int64  `json:"id"`
	AvatarUrl string `json:"avatar_url"`
}

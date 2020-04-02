package client

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

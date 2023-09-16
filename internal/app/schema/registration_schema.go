package schema

type RegisterReq struct {
	Fullname string `validate:"required" json:"fullname"`
	Password string `validate:"required,alphanum,min=8" json:"password"`
	Email    string `validate:"required,email" json:"email"`
}

type GetUserResp struct {
	Fullname string `json:"fullname"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

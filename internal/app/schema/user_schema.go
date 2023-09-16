package schema

type GetUsersResp struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type GetUserReq struct {
	UserID int `validate:"required,eq=3" json:"user_id"`
}

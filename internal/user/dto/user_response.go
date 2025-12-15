package dto

type Response struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

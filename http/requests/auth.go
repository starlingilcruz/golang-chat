package requests


type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpParams struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
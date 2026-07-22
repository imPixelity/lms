package user

type createUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type updateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

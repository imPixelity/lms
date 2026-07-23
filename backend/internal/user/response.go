package user

type userResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func newUserResponse(u *User) userResponse {
	return userResponse{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}

type userListResponse struct {
	Users []userResponse `json:"users"`
	Total int            `json:"total"`
}

func NewUserListResponse(users []User) userListResponse {
	out := make([]userResponse, len(users))
	for i, u := range users {
		out[i] = newUserResponse(&u)
	}
	return userListResponse{Users: out, Total: len(out)}
}

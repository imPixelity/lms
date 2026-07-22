package user

type userResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func newUserResponse(u *user) userResponse {
	return userResponse{
		ID:       u.id,
		Email:    u.email,
		Username: u.username,
	}
}

type userListResponse struct {
	Users []userResponse `json:"users"`
	Total int            `json:"total"`
}

func NewUserListResponse(users []user) userListResponse {
	out := make([]userResponse, len(users))
	for i, u := range users {
		out[i] = newUserResponse(&u)
	}
	return userListResponse{Users: out, Total: len(out)}
}

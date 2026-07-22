package user

type UserResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUserResponse(u *user) UserResponse {
	return UserResponse{
		ID:       u.id,
		Email:    u.email,
		Username: u.username,
	}
}

type UserListResponse struct {
	Users []UserResponse `json:"users"`
	Total int            `json:"total"`
}

func NewUserListResponse(users []user) UserListResponse {
	out := make([]UserResponse, len(users))
	for i, u := range users {
		out[i] = NewUserResponse(&u)
	}
	return UserListResponse{Users: out, Total: len(out)}
}

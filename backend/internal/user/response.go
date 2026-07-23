package user

type UserResponse struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func NewUserResponse(user *User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}

type UsersResponse struct {
	Users      []UserResponse `json:"users"`
	HasMore    bool           `json:"hasMore"`
	NextCursor int64          `json:"nextCursor"`
}

func NewUsersResponse(users []User, hasMore bool, nextCursor int64) UsersResponse {
	out := make([]UserResponse, 0, len(users))
	for _, v := range users {
		out = append(out, NewUserResponse(&v))
	}
	return UsersResponse{
		Users:      out,
		HasMore:    hasMore,
		NextCursor: nextCursor,
	}
}

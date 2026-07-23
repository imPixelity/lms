package user

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	return nil
}

func (r *CreateUserRequest) ToModel() *User {
	return &User{
		Email: r.Email,
		Username: r.Username,
		Password: r.Password,
	}
}

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (r *UpdateUserRequest) Validate() error {
	return nil
}

func (r *UpdateUserRequest) ToModel() *User {
	return &User{
		Email: r.Email,
		Username: r.Username,
	}
}

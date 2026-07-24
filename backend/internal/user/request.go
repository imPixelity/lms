package user

import (
	"errors"
	"fmt"
	"strings"
)

var emailCharset [256]bool
var usernameCharset [256]bool

func init() {
	const alnum = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789"

	for i := range len(alnum) {
		emailCharset[alnum[i]] = true
		usernameCharset[alnum[i]] = true
	}

	for _, c := range "@.-" {
		emailCharset[byte(c)] = true
	}

	for _, c := range "_.-" {
		emailCharset[byte(c)] = true
	}
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if err := r.validateEmail(); err != nil {
		return err
	}

	if err := r.validateUsername(); err != nil {
		return err
	}

	if err := r.validatePassword(); err != nil {
		return err
	}

	return nil
}
func (r *CreateUserRequest) validateEmail() error {
	r.Email = strings.ToLower(strings.TrimSpace(r.Email))

	if r.Email == "" {
		return errors.New("email is required")
	}

	if len(r.Email) > 128 {
		return errors.New("email is too long")
	}

	for i := 0; i < len(r.Email); i++ {
		if !emailCharset[r.Email[i]] {
			return fmt.Errorf("email contains invalid character: %q", r.Email[i])
		}
	}

	at := strings.IndexByte(r.Email, '@')
	if at <= 0 || at != strings.LastIndexByte(r.Email, '@') || at == len(r.Email)-1 {
		return errors.New("email must contain exactly one '@' with content on both sides")
	}

	local := r.Email[:at]
	domain := r.Email[at+1:]

	if local[0] == '.' || local[len(local)-1] == '.' || strings.Contains(local, "..") {
		return errors.New("email local part is invalid")
	}

	if !strings.Contains(domain, ".") || domain[0] == '.' || domain[len(domain)-1] == '.' || strings.Contains(domain, "..") {
		return errors.New("email domain part is invalid")
	}

	return nil
}

func (r *CreateUserRequest) validateUsername() error {
	r.Username = strings.ToLower(strings.TrimSpace(r.Username))

	if len(r.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	if len(r.Username) > 30 {
		return errors.New("username must be at most 20 characters")
	}

	for i := 0; i < len(r.Username); i++ {
		if !usernameCharset[r.Username[i]] {
			return fmt.Errorf("username contains invalid character: %q", r.Username[i])
		}
	}

	return nil
}

func (r *CreateUserRequest) validatePassword() error {
	r.Password = strings.TrimSpace(r.Password)

	if len(r.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	if len(r.Password) > 72 {
		return errors.New("password must be at most 72 characters")
	}

	return nil
}

func (r *CreateUserRequest) ToModel() *User {
	return &User{
		Email:    r.Email,
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
		Email:    r.Email,
		Username: r.Username,
	}
}

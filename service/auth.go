package service

import "time"

type AuthSignUpResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type AuthService interface {
	SignUp(email, name, password string) (*AuthSignUpResponse, error)
	SignIn(email string, password string) (*int, error)
}

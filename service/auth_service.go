package service

import (
	"gobasic/repository"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return authService{userRepo: userRepo}
}

func (s authService) SignUp(email, name, password string) (*AuthSignUpResponse, error) {

	user, err := s.userRepo.Create(email, name, password)
	if err != nil {
		return nil, err
	}

	authSignUpRes := AuthSignUpResponse{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return &authSignUpRes, nil
}

func (s authService) SignIn(email string, password string) (*int, error) {

	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	id := int(user.ID)

	return &id, nil
}
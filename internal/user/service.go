package user

import (
	"errors"
	"fmt"
	"sportsync/internal/auth"
	"sportsync/internal/user/dto"
)

var userCredentialError = errors.New("Invalid User Credential")

type service struct {
	repo       Repository
	jwtService auth.JwtService
}

func NewService(repo Repository, jwtService auth.JwtService) *service {
	return &service{
		repo, jwtService,
	}
}

func (s *service) CreateUser(req dto.CreateUserRequest) (*dto.RegisterUserResponse, error) {
	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	user.HashPassword()

	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return &dto.RegisterUserResponse{
		Success: true,
		Message: "User registered successfully",
		Data: dto.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil

}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, userCredentialError
	}

	if err := user.checkPassword(req.Password); err != nil {
		return nil, userCredentialError
	}


	token, err := s.jwtService.GenerateJwtToken(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token : %w", err)
	}
	return &dto.LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginData{
			Token: token,
			User: dto.UserLoginDetails{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Role:  user.Role,
			},
		},
	}, nil

}

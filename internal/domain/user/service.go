package user

import (
	"SpotSync/internal/auth"
	"SpotSync/internal/domain/user/dto"
	"errors"
	"fmt"
)


var ErrInvalidCredentials = errors.New("invalid credentials")


type service struct{
	repo Repository
	jwtService auth.JWTService
}

func NewService (repo Repository, jwtservice auth.JWTService) *service{
	return &service{
		repo: repo,
		jwtService: jwtservice,
	}
}

func (s *service) CreateUser(req dto.CreateRequest) (*dto.CreateUserResponse, error){
	user := User{
		Name : req.Name,
		Email: req.Email,
		Role: req.Role,
		Password: req.Password,
	}

	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &dto.CreateUserResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
	


}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.LoginUserResponse, error){
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}
	
	if user == nil {
		return nil, ErrInvalidCredentials
	}

	err = user.checkPassword(req.Password)

	if err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := s.jwtService.GenerateToken(user.ID, user.Email, user.Name, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return &dto.LoginUserResponse{
		Token: token,
		User: dto.User{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
		},
	}, nil


}





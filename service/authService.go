package service

import (
	"github.com/Bogdanov-G/authorization_app/domain"
	"github.com/Bogdanov-G/authorization_app/dto"
	"github.com/Bogdanov-G/authorization_app/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(tr domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{tr}
}

func (s DefaultAuthService) Login(logReq dto.LoginRequest) (*string, *errs.AppError) {
	login, appErr := s.repo.FindUser(logReq.Username, logReq.Password)
	if appErr != nil {
		return nil, appErr
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

// func (s DefaultAuthService) Register(regReq dto.LoginRequest) (*dto.LoginResponse, error) {
// }

// func (s DefaultAuthService) Verify(verReq dto.LoginRequest) (*dto.LoginResponse, error) {
// }

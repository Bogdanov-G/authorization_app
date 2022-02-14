package service

import (
	"github.com/Bogdanov-G/authorization_app/domain"
	"github.com/Bogdanov-G/authorization_app/dto"
	"github.com/Bogdanov-G/authorization_app/errs"
	"github.com/golang-jwt/jwt"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, *errs.AppError)
	Verify(map[string]string) (bool, *errs.AppError)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func NewAuthService(tr domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{
		tr,
		domain.NewRolePermissions(),
	}
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

func (s DefaultAuthService) Verify(urlParams map[string]string) (bool, *errs.AppError) {
	token, err := stringToJWT(urlParams["token"])
	if err != nil {
		return false, errs.NewUnauthorizedError(err.Error())
	}
	if !token.Valid {
		return false, errs.NewValidationError("invalid token")
	}

	mapClaims := token.Claims.(jwt.MapClaims)
	claims, err := domain.BuildClaims(mapClaims)
	if err != nil {
		return false, errs.NewUnauthorizedError(err.Error())
	}

	if !claims.IsRequestVerified(urlParams) {
		return false, errs.NewValidationError("access denied")
	}
	return s.rolePermissions.IsAuthorized(claims.Role, urlParams["routeName"]), nil
}

func stringToJWT(str string) (*jwt.Token, error) {
	token, err := jwt.Parse(str, func(t *jwt.Token) (interface{}, error) {
		return []byte(domain.SigningKeySample), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

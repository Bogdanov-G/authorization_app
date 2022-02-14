package domain

import (
	"database/sql"
	"time"

	"github.com/Bogdanov-G/authorization_app/errs"
	"github.com/Bogdanov-G/authorization_app/logger"
	"github.com/golang-jwt/jwt"
)

const TokenDuration = time.Hour
const SigningKeySample = "SignInKeySample"

type Login struct {
	Username   string          `db:"username"`
	CustomerId sql.NullString  `db:"customer_id"`
	Accounts   []sql.NullInt64 `db:"accounts_ids"`
	Role       string          `db:"role"`
}

func (l Login) GenerateToken() (*string, *errs.AppError) {
	var claims jwt.MapClaims
	if l.Role == "employee" {
		claims = l.makeEmployeeClaims()
	} else {
		claims = l.makeCustomerClaims()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(SigningKeySample))
	if err != nil {
		logger.Error("token signing error: " + err.Error())
		return nil, errs.NewUnexpectedError("internal server error")
	}
	return &ss, nil
}

func (l Login) makeCustomerClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"username":    l.Username,
		"customer_id": l.CustomerId.String,
		"accounts":    l.Accounts,
		"role":        l.Role,
		"expiration":  time.Now().Add(TokenDuration).Unix(),
	}
}

func (l Login) makeEmployeeClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"username":   l.Username,
		"role":       l.Role,
		"expiration": time.Now().Add(TokenDuration).Unix(),
	}
}

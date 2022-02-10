package domain

import (
	"database/sql"

	"github.com/lib/pq"

	"github.com/Bogdanov-G/authorization_app/errs"
	"github.com/Bogdanov-G/authorization_app/logger"
)

type AuthRepository struct {
	pool *sql.DB
}

type AuthRepository_I interface {
	FindUser(string, string) (Login, error)
}

func NewAuthRepository(pool *sql.DB) AuthRepository {
	return AuthRepository{pool}
}

func (s AuthRepository) FindUser(name, pass string) (*Login, *errs.AppError) {
	var login Login

	row := s.pool.QueryRow(
		`SELECT username,
				u.customer_id,
				ARRAY_AGG(account_id) AS accounts_ids,
				role
		FROM users u
		LEFT JOIN accounts a ON (u.customer_id = a.customer_id)
		WHERE username = $1 AND password = $2
		GROUP BY u.username ;`, name, pass)
	err := row.Scan(&login.Username, &login.CustomerId, pq.Array(&login.Accounts), &login.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("wrong username and/or password, or user doesn't exist")
		} else {
			logger.Error("error searching for user: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &login, nil
}

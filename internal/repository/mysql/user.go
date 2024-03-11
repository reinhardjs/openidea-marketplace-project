package mysql

import (
	"context"
	"database/sql"

	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/user"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) user.Repository {
	return &userRepository{conn}
}

func (m *userRepository) Register(ctx context.Context, user *domain.User) (err error) {
	query := "INSERT user SET name=? , username=? , password=?"
	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, user.Name, user.Username, user.Password)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	user.ID = lastID
	return
}

package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/openidea-marketplace/domain/entities"
	"github.com/openidea-marketplace/user"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) user.Repository {
	return &userRepository{conn}
}

func (m *userRepository) Insert(ctx context.Context, user *entities.User) error {
	query := "INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id"
	row := m.Conn.QueryRowContext(ctx, query, user.Username, user.Name, user.Password)

	var lastID int64
	err := row.Scan(&lastID)

	// PostgreSQL IDs start from 1
	if lastID > 0 {
		user.ID = lastID
		return nil
	}

	return err
}

func (m *userRepository) FindByUsername(ctx context.Context, username string) (user entities.User, err error) {
	query := "SELECT * FROM users WHERE username = $1"
	row := m.Conn.QueryRowContext(ctx, query, username)

	err = row.Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = fmt.Errorf("user with username %s not found", username)
		}
	}

	return
}

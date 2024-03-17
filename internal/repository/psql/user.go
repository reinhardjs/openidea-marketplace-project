package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/openidea-marketplace/domain"
	"github.com/openidea-marketplace/domain/entities"
	"github.com/openidea-marketplace/user"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) user.Repository {
	return &userRepository{conn}
}

func (m *userRepository) Insert(ctx context.Context, user *entities.User) error {
	fmt.Print(user)
	query := "INSERT INTO users (username, name, password) VALUES ($1, $2, $3) RETURNING id"
	row := m.Conn.QueryRowContext(ctx, query, user.Username, user.Name, "password")

	var lastID int64
	err := row.Scan(&lastID)

	// PostgreSQL IDs start from 1
	if lastID > 0 {
		user.ID = lastID
		return nil
	}

	return err
}

func (m *userRepository) fetch(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	rows, err = m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	return rows, nil
}

func (m *userRepository) FindByUsernameAndPassword(ctx context.Context, username, password string) (user *entities.User, err error) {
	query := `SELECT * FROM users WHERE username = ? AND password = ?`

	rows, err := m.fetch(ctx, query, username, password)
	if err != nil {
		return nil, err
	}

	result := make([]entities.User, 0)
	for rows.Next() {
		t := entities.User{}
		banks := make([]entities.Bank, 0)

		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.Password,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		t.Banks = banks

		result = append(result, t)
	}

	if len(result) > 0 {
		user = &result[0]
	} else {
		return nil, domain.ErrNotFound
	}

	return
}

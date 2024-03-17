package entities

import "time"

type User struct {
	ID        int64
	Name      string
	Username  string
	Password  string
	Banks     []Bank
	CreatedAt time.Time
	UpdatedAt time.Time
}

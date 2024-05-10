package model

import (
	"database/sql"
	"time"
)

const (
	Role_USER  Role = 0
	Role_ADMIN Role = 1
)

type Role int32

type User struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      Role         `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserCreate struct {
	Name            string `db:"name"`
	Email           string `db:"email"`
	Password        string `db:"password"`
	PasswordConfirm string `db:"password_confirm"`
	Role            Role   `db:"role"`
}

type UserUpdate struct {
	ID    int64  `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
}

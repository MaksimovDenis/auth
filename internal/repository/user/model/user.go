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
	ID        int64
	Name      string
	Email     string
	Role      Role
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserCreate struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string
	Role            Role
}

type UserUpdate struct {
	ID    int64
	Name  string
	Email string
}

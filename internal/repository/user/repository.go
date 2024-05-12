package user

import (
	"context"
	"errors"
	"fmt"
	"github.com/MaksimovDenis/auth/internal/client/db"
	"github.com/MaksimovDenis/auth/internal/model"
	"github.com/MaksimovDenis/auth/internal/repository"
	"github.com/MaksimovDenis/auth/internal/repository/user/converter"
	modelRepo "github.com/MaksimovDenis/auth/internal/repository/user/model"
	sq "github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

const (
	tableName = "users"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role" // Используем правильное имя столбца
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"

	tableName2 = "roles"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GET",
		QueryRaw: query,
	}

	var user modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Create(ctx context.Context, user *model.UserCreate) (int64, error) {

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, user.Password, user.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	fmt.Println(query)
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userID int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userID)
	if err != nil {
		log.Printf("failed to insert user: %v", err)
		return 0, status.Errorf(codes.Internal, "Internal server error")
	}

	return userID, nil
}

func (r *repo) Update(ctx context.Context, user *model.UserUpdate) (*empty.Empty, error) {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, user.Name).
		Set(emailColumn, user.Email).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &empty.Empty{}, nil
}

func (r *repo) Delete(ctx context.Context, id int64) (*empty.Empty, error) {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "note_repository.Delete",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	rowsAffected := res.RowsAffected()

	if rowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &empty.Empty{}, nil
}

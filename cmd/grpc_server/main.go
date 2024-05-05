package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	desc "github.com/MaksimovDenis/auth/pkg/userAPI_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const grpcPort = 50051

const (
	dbDSN = "host=localhost port=54321 dbname=note user=note-user password=note-password sslmode=disable"
)

var t = time.Now()

type server struct {
	desc.UnimplementedUserAPIV1Server
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {

	builderSelect := sq.Select("id", "name", "email", "role", "created_at", "updated_at").
		From("user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID}).
		Limit(1)

	query, args, err := builderSelect.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	rows, err := pool.Query(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to select user: %v", err)
	}

	var id int64
	var name string
	var email string
	var created_at *timestamppb.Timestamp
	var updated_at *timestamppb.Timestamp

	var roleString string
	var role userAPI_v1.Role

	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &roleString, &created_at, &updated_at)
		if err != nil {
			log.Fatalf("failed to scan user: %v", err)
		}
	}

	switch roleString {
	case "USER":
		role = userAPI_v1.Role_USER
	case "ADMIN":
		role = userAPI_v1.Role_ADMIN
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: created_at,
			UpdatedAt: updated_at,
		},
	}, nil
}

// CREATE...
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	builderInsert := sq.Insert("user").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(gofakeit.BeerName(), gofakeit.Email(), "admin", "USER").
		Suffix("RETURNING id")

	query, args, err := builderInsert.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	var userID int
	err = pool.QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	return &desc.CreateResponse{
		Id: int64(userID),
	}, nil
}

// PUT
func (s *server) Update(ctx context.Context, req *desc.UpdateRequest) (*empty.Empty, error) {

	builderUpdate := sq.Update("user").
		PlaceholderFormat(sq.Dollar).
		Set("name", gofakeit.BeerName()).
		Set("email", gofakeit.Email()).
		Where(sq.Eq{"id": userID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err := pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to update user: %v", err)
	}

	log.Printf("updated %d rows", res.RowsAffected())

	return &empty.Empty{}, nil
}

// DELETE
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {

	builderDelete := sq.Delete("user").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": userID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		log.Fatalf("failed to build query: %v", err)
	}

	res, err := pool.Exec(ctx, query, args...)
	if err != nil {
		log.Fatalf("failed to delete user: %v", err)
	}

	if res.RowsAffected() == 0 {
		return nil, errors.New("user not found")
	}

	return &empty.Empty{}, nil
}

func main() {

	ctx := context.Background()

	// Создаем пул соединений с базой данных
	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserAPIV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

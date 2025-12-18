package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/SubhaNITS150/dob-calculator/db/sqlc"
	"github.com/google/uuid"
)

type UserRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		q: sqlc.New(db),
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	name string,
	dob time.Time,
) (sqlc.CreateUserRow, error) {

	return r.q.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) Update(
	ctx context.Context,
	id uuid.UUID,
	name string,
	dob time.Time,
) (sqlc.UpdateUserRow, error) {

	return r.q.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
}

func (r *UserRepository) Delete(
	ctx context.Context,
	id uuid.UUID,
) error {

	return r.q.DeleteUser(ctx, id)
}

func (r *UserRepository) List(
	ctx context.Context,
) ([]sqlc.ListUsersRow, error) {

	return r.q.ListUsers(ctx)
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*sqlc.GetUserByIDRow, error) {
	user, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}



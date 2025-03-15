package repo

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"pract2/internal/models"
)

const (
	SingUpUserQuery   = `INSERT INTO users(username, password) VALUES ($1, $2);`
	IdentityUserQuery = `SELECT * FROM users WHERE username = $1;`
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) SingUp(ctx context.Context, username, password string) error {
	cmdTag, err := r.pool.Exec(ctx, SingUpUserQuery, username, password)
	fmt.Println(cmdTag)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) SingIn(ctx context.Context, username string) (models.User, error) {
	row := r.pool.QueryRow(ctx, IdentityUserQuery, username)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

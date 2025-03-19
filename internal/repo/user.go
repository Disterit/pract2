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
	DeleteUserQuery   = `DELETE FROM users WHERE id = $1;`
)

type userRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) User {
	return &userRepository{pool}
}

func (r *userRepository) SingUp(ctx context.Context, username, password string) error {
	cmdTag, err := r.pool.Exec(ctx, SingUpUserQuery, username, password)
	fmt.Println(cmdTag)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) SingIn(ctx context.Context, username string) (models.User, error) {
	row := r.pool.QueryRow(ctx, IdentityUserQuery, username)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, DeleteUserQuery, id)
	if err != nil {
		return err
	}
	return nil
}

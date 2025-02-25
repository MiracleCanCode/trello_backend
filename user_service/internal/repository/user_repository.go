package repository

import (
	"context"
	"fmt"

	"github.com/clone_trello/services/user_service/models"
	"github.com/jackc/pgx/v5"
)

type User struct {
	db *pgx.Conn
}

func NewUser(db *pgx.Conn) *User {
	return &User{
		db: db,
	}
}

func (u *User) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	sqlQuery := `
        INSERT INTO users (name, login, id, avatar, password) 
        VALUES ($1, $2, $3, $4)
        RETURNING name, login, id, avatar
    `

	row := u.db.QueryRow(
		ctx,
		sqlQuery,
		user.Name,
		user.Login,
		user.Id,
		user.Avatar,
		user.Password,
	)

	createdUser := models.User{}

	err := row.Scan(
		&createdUser.Name,
		&createdUser.Login,
		&createdUser.Id,
		&createdUser.Avatar,
		&createdUser.Password,
	)

	if err != nil {
		return nil, fmt.Errorf("CreateUser: failed create user: %w", err)
	}

	return &createdUser, nil
}

func (u *User) GetUserByLogin(ctx context.Context, login string) (*models.User, error) {
	sqlQuery := "SELECT id, login, name, avatar, password FROM users WHERE login = $1"

	row := u.db.QueryRow(ctx, sqlQuery, login)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Login, &user.Name, &user.Avatar, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("GetUserByLogin: failed scan user data to user model struct: %w", err)
	}

	return user, nil
}
func (u *User) GetUserById(ctx context.Context, id string) (*models.User, error) {
	sqlQuery := "SELECT id, login, name, avatar, password FROM users WHERE id = $1"

	row := u.db.QueryRow(ctx, sqlQuery, id)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Login, &user.Name, &user.Avatar, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("GetUserById: failed scan user data to user model struct: %w", err)
	}

	return user, nil
}
func (u *User) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	sqlQuery := "UPDATE users SET login = $1, name = $2, avatar = $3 WHERE login = $4"
	row := u.db.QueryRow(ctx, sqlQuery, user.Login, user.Name, user.Avatar, user.Login)

	updatedUser := &models.User{}
	if err := row.Scan(&updatedUser.Id, &updatedUser.Login, &updatedUser.Name, &updatedUser.Avatar, &updatedUser.Password); err != nil {
		return nil, fmt.Errorf("UpdateUser: failed update user data: %w", err)
	}

	return updatedUser, nil
}

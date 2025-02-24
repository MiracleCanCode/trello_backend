package usecase

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context)
}

type User struct {
	repo UserRepository
}

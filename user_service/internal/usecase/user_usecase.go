package usecase

import (
	"context"

	"github.com/MiracleCanCode/common_libary_trello/pkg/logger"
	pb "github.com/MiracleCanCode/trello_protos/pkg/api"
	"github.com/clone_trello/services/user_service/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByLogin(ctx context.Context, login string) (*models.User, error)
	GetUserById(ctx context.Context, id string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
}

type User struct {
	repo   userRepository
	logger *zap.Logger
	pb.UnimplementedUserServiceServer
}

func NewUser(repo userRepository, ctx context.Context) *User {
	logger := logger.GetLogger(ctx)
	return &User{
		repo:   repo,
		logger: logger,
	}
}

func (u *User) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	generatedID := u.generateID()
	userModel := &models.User{
		Id:       generatedID,
		Login:    in.User.Login,
		Name:     in.User.Name,
		Password: in.User.Password,
	}
	user, err := u.repo.CreateUser(ctx, userModel)
	if err != nil {
		u.logger.Error("Failed create user", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Ошибка, создания пользователя: %w", err)
	}

	return &pb.CreateUserResponse{
		User: models.MapToGRPCUser(user),
	}, nil
}
func (u *User) GetUserByLogin(ctx context.Context, in *pb.GetUserByLoginRequest) (*pb.GetUserByLoginResponse, error) {
	user, err := u.repo.GetUserByLogin(ctx, in.Login)
	if err != nil {
		u.logger.Error("Failed get user by login", zap.Error(err), zap.String("user_login", in.Login))
		return nil, status.Errorf(codes.Internal, "Ошибка получения пользователя по логину: %w", err)
	}

	pointerUserByLoginGRPC := &pb.UserByLogin{
		Id:     user.Id,
		Login:  user.Login,
		Name:   user.Name,
		Avatar: user.Avatar,
	}
	return &pb.GetUserByLoginResponse{
		User: pointerUserByLoginGRPC,
	}, nil
}
func (u *User) GetUserById(ctx context.Context, in *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := u.repo.GetUserById(ctx, in.Id)
	if err != nil {
		u.logger.Error("Failed get user by id", zap.Error(err), zap.String("user_id", in.Id))
		return nil, status.Errorf(codes.NotFound, "Ошибка получения пользователя по id=%s:%w", in.Id, err)
	}

	return &pb.GetUserByIdResponse{
		User: models.MapToGRPCUser(user),
	}, nil
}

func (u *User) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	user := models.MapToModelUser(in.User)
	updatedUser, err := u.repo.UpdateUser(ctx, user)
	if err != nil {
		u.logger.Error("Failed update user", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "Ошибка обновления данных пользователя: %w", err)
	}

	return &pb.UpdateUserResponse{
		User: models.MapToGRPCUser(updatedUser),
	}, nil
}

func (u *User) generateID() string {
	return uuid.New().String()
}

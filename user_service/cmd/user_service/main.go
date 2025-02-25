package main

import (
	"context"

	"github.com/MiracleCanCode/common_libary_trello/pkg/logger"
	loggerInstance "github.com/MiracleCanCode/example_configuration_logger"
	"github.com/clone_trello/services/user_service/config"
	"github.com/clone_trello/services/user_service/internal/grpc"
	"github.com/clone_trello/services/user_service/internal/repository"
	"github.com/clone_trello/services/user_service/internal/usecase"
	"github.com/clone_trello/services/user_service/pkg/storage/postgres"
	"go.uber.org/zap"
)

func main() {
	loggerCfg := loggerInstance.DefaultLoggerConfig()
	log := loggerInstance.Logger(loggerCfg)
	cfg, err := config.MustLoad()
	if err != nil {
		log.Error("Failed load config", zap.Error(err))
		return
	}

	ctx := logger.WithLogger(context.Background(), log)
	pg := postgres.New(cfg.DSN)
	connPostgres, err := pg.Conn(ctx)
	if err != nil {
		log.Error("Failed connect to postgresql db", zap.Error(err))
		return
	}

	defer connPostgres.Close(ctx)

	userRepository := repository.NewUser(connPostgres)
	userUsecase := usecase.NewUser(userRepository, ctx)

	server := grpc.New(userUsecase)
	serverRegisterMessage, err := server.Conn(cfg.ADDR)
	if err != nil {
		log.Error("Failed create conn to grpc server", zap.Error(err))
		return
	}

	log.Info(serverRegisterMessage)
}

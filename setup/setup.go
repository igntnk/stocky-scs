package setup

import (
	grpcapp "github.com/igntnk/stocky-scs/grpc"
	mongorepo "github.com/igntnk/stocky-scs/repository/mongo"
	"github.com/igntnk/stocky-scs/service"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func GRPCServer() *grpc.Server {
	return grpcServer
}

func Init(db *mongo.Database, isReplicaSet bool, logger zerolog.Logger) error {
	var (
		userRepo = mongorepo.NewUserRepository(db, isReplicaSet, logger)
		userServ = service.NewUserService(logger, userRepo)
	)

	grpcServer = grpc.NewServer()
	grpcapp.RegisterUserServer(grpcServer, logger, userServ)

	return nil
}

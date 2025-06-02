package setup

import (
	"context"
	grpcapp "github.com/igntnk/stocky-scs/grpc"
	"github.com/igntnk/stocky-scs/repository"
	mongorepo "github.com/igntnk/stocky-scs/repository/mongo"
	"github.com/igntnk/stocky-scs/service"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func GRPCServer() *grpc.Server {
	return grpcServer
}

func SetupDefaultData(ctx context.Context, db *mongo.Database) error {
	userRepo := db.Collection(repository.UserCollection)

	_, err := userRepo.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{Key: " name", Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}

	return nil
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

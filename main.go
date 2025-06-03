package main

import (
	"github.com/igntnk/stocky-scs/config"
	"github.com/igntnk/stocky-scs/grpc"
	"github.com/igntnk/stocky-scs/pkg/client"
	"github.com/igntnk/stocky-scs/setup"
	"go.mongodb.org/mongo-driver/mongo/description"
	"os/signal"
	"syscall"

	"context"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	ctx := context.Background()

	cfg := config.Get(logger)

	db, topology, err := client.NewClient(ctx, cfg.Database, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	logger.Info().Msg("Connection is established. Mongo use topology: " + topology.String())

	isReplicaSet := false
	if topology.Kind()&description.ReplicaSet == description.ReplicaSet {
		isReplicaSet = true
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = setup.Init(db, isReplicaSet, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}

	grpcServ := grpc.New(setup.GRPCServer(), cfg.Server.GrpcPort, logger)

	go func() {
		grpcServ.MustRun()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	grpcServ.Stop()
}

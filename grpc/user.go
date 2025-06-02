package grpc

import (
	"context"
	"github.com/igntnk/stocky-scs/proto/pb"
	"github.com/igntnk/stocky-scs/service"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type productServer struct {
	pb.UnimplementedUserServiceServer
	Logger zerolog.Logger
	serv   service.UserService
}

func RegisterUserServer(server *grpc.Server, logger zerolog.Logger, userService service.UserService) {
	pb.RegisterUserServiceServer(server, &productServer{Logger: logger, serv: userService})
}

func (s *productServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.IdResponse, error) {
	return s.serv.CreateUser(ctx, req)
}

func (s *productServer) BlockUser(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	return s.serv.BlockUser(ctx, req)
}

func (s *productServer) UnblockUser(ctx context.Context, req *pb.IdRequest) (*pb.IdResponse, error) {
	return s.serv.UnblockUser(ctx, req)
}

func (s *productServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.IdResponse, error) {
	return s.serv.UpdateUser(ctx, req)
}

func (s *productServer) GetById(ctx context.Context, req *pb.IdRequest) (*pb.UserModel, error) {
	return s.serv.GetById(ctx, req)
}

func (s *productServer) GetAllUsers(ctx context.Context, req *emptypb.Empty) (*pb.GetAllUsersResponse, error) {
	return s.serv.GetAllUsers(ctx)
}

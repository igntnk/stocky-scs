package service

import (
	"context"
	"github.com/igntnk/stocky-scs/models"
	"github.com/igntnk/stocky-scs/proto/pb"
	"github.com/igntnk/stocky-scs/repository"
	"github.com/rs/zerolog"
	"time"
)

type UserService interface {
	CreateUser(ctx context.Context, user *pb.CreateUserRequest) (*pb.IdResponse, error)
	BlockUser(ctx context.Context, user *pb.IdRequest) (*pb.IdResponse, error)
	UnblockUser(ctx context.Context, user *pb.IdRequest) (*pb.IdResponse, error)
	UpdateUser(ctx context.Context, user *pb.UpdateUserRequest) (*pb.IdResponse, error)
	GetById(ctx context.Context, user *pb.IdRequest) (*pb.UserModel, error)
	GetAllUsers(ctx context.Context) (*pb.GetAllUsersResponse, error)
}

func NewUserService(logger zerolog.Logger, repo repository.UserRepository) UserService {
	return &userService{
		Logger: logger,
		repo:   repo,
	}
}

type userService struct {
	Logger zerolog.Logger
	repo   repository.UserRepository
}

func (u userService) CreateUser(ctx context.Context, user *pb.CreateUserRequest) (*pb.IdResponse, error) {
	res, err := u.repo.CreateUser(ctx, &models.User{
		Name:           user.GetName(),
		Description:    user.GetDescription(),
		DocumentType:   user.GetDocumentType(),
		DocumentNumber: user.GetDocumentNumber(),
		CreationDate:   time.Now().String(),
		Blocked:        false,
		AuthId:         user.AuthId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.IdResponse{Id: res}, nil
}

func (u userService) BlockUser(ctx context.Context, user *pb.IdRequest) (*pb.IdResponse, error) {
	res, err := u.repo.BlockUser(ctx, user.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.IdResponse{Id: res}, nil
}

func (u userService) UnblockUser(ctx context.Context, user *pb.IdRequest) (*pb.IdResponse, error) {
	res, err := u.repo.UnblockUser(ctx, user.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.IdResponse{Id: res}, nil
}

func (u userService) UpdateUser(ctx context.Context, user *pb.UpdateUserRequest) (*pb.IdResponse, error) {
	res, err := u.repo.UpdateUser(ctx, &models.User{
		Name:           user.GetName(),
		Description:    user.GetDescription(),
		DocumentType:   user.GetDocumentType(),
		DocumentNumber: user.GetDocumentNumber(),
		Id:             user.GetId(),
	})

	if err != nil {
		return nil, err
	}

	return &pb.IdResponse{Id: res}, nil
}

func (u userService) GetById(ctx context.Context, user *pb.IdRequest) (*pb.UserModel, error) {
	res, err := u.repo.GetById(ctx, user.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.UserModel{
		Id:             res.Id,
		Name:           res.Name,
		Description:    res.Description,
		DocumentType:   res.DocumentType,
		DocumentNumber: res.DocumentNumber,
		CreationDate:   res.CreationDate,
		Blocked:        res.Blocked,
		AuthId:         res.AuthId,
	}, nil
}

func (u userService) GetAllUsers(ctx context.Context) (*pb.GetAllUsersResponse, error) {
	res, err := u.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]*pb.UserModel, len(res))
	result := &pb.GetAllUsersResponse{}
	for i, re := range res {
		users[i] = &pb.UserModel{
			Id:             re.Id,
			Name:           re.Name,
			Description:    re.Description,
			DocumentType:   re.DocumentType,
			DocumentNumber: re.DocumentNumber,
			CreationDate:   re.CreationDate,
			Blocked:        re.Blocked,
			AuthId:         re.AuthId,
		}
	}

	result.Users = users
	return result, nil
}

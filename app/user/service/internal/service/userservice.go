package service

import (
	"context"

	pb "graduate-pro/api/user/v1"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServiceService() *UserServiceService {
	return &UserServiceService{}
}

func (s *UserServiceService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	return &pb.GetUserByIdReply{}, nil
}
func (s *UserServiceService) GetUserByName(ctx context.Context, req *pb.GetUserByNameRequest) (*pb.GetUserByNameReply, error) {
	return &pb.GetUserByNameReply{}, nil
}
func (s *UserServiceService) SearchUserByName(ctx context.Context, req *pb.SearchUserByNameRequest) (*pb.SearchUserByNameReply, error) {
	return &pb.SearchUserByNameReply{}, nil
}
func (s *UserServiceService) SaveUser(ctx context.Context, req *pb.SaveUserRequest) (*pb.SaveUserReply, error) {
	return &pb.SaveUserReply{}, nil
}
func (s *UserServiceService) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserReply, error) {
	return &pb.RemoveUserReply{}, nil
}

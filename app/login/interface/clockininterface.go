package service

import (
	"context"

	pb "graduate-pro/api/login/interface/v1"
)

type ClockinInterfaceService struct {
	pb.UnimplementedClockinInterfaceServer
}

func NewClockinInterfaceService() *ClockinInterfaceService {
	return &ClockinInterfaceService{}
}

func (s *ClockinInterfaceService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{}, nil
}
func (s *ClockinInterfaceService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	return &pb.LoginReply{}, nil
}
func (s *ClockinInterfaceService) ClockinOnWork(ctx context.Context, req *pb.ClockinOnWorkRequest) (*pb.ClockinOnWorkReply, error) {
	return &pb.ClockinOnWorkReply{}, nil
}
func (s *ClockinInterfaceService) ClockinOffWork(ctx context.Context, req *pb.ClockinOffWorkRequest) (*pb.ClockinOffWorkReply, error) {
	return &pb.ClockinOffWorkReply{}, nil
}

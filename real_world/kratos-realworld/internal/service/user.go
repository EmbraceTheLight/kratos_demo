package service

import (
	"context"
	"kratos-realworld/internal/biz"

	pb "kratos-realworld/api/realworld/v1"
)

type RealWorldService struct {
	pb.UnimplementedRealWorldServer

	uc *biz.UserUsecase
	sc *biz.SocialUsecase
}

func NewRealWorldService(uc *biz.UserUsecase, sc *biz.SocialUsecase) *RealWorldService {
	return &RealWorldService{
		uc: uc,
		sc: sc,
	}
}

func (s *RealWorldService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *pb.EmptyRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *pb.FollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *pb.UnFollowUserRequest) (*pb.ProfileReply, error) {
	return &pb.ProfileReply{}, nil
}

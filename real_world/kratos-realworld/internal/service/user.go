package service

import (
	"context"
	"kratos-realworld/internal/biz"

	v1 "kratos-realworld/api/realworld/v1"
)

type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.UserUsecase
	sc *biz.SocialUsecase
}

func NewRealWorldService(uc *biz.UserUsecase, sc *biz.SocialUsecase) *RealWorldService {
	return &RealWorldService{
		uc: uc,
		sc: sc,
	}
}

func (s *RealWorldService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) Register(ctx context.Context, req *v1.RegisterRequest) (*v1.UserReply, error) {
	user, err := s.uc.Register(
		ctx,
		req.User.Username,
		req.User.Email,
		req.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{
		User: &v1.UserReply_User{
			Username: user.Username,
			Token:    user.Token,
		},
	}, nil
}
func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.EmptyRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnFollowUser(ctx context.Context, req *v1.UnFollowUserRequest) (*v1.ProfileReply, error) {
	return &v1.ProfileReply{}, nil
}

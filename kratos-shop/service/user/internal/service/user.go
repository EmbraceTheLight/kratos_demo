package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	user, err := u.uc.Create(ctx, &biz.User{
		Mobile:   req.Mobile,
		Password: req.Password,
		NickName: req.NickName,
	})
	if err != nil {
		return nil, err
	}
	userInfoResp := &pb.CreateUserReply{
		UserInfo: &pb.UserInfo{
			Id:       user.ID,
			Mobile:   user.Mobile,
			Password: user.Password,
			NickName: user.NickName,
			Gender:   user.Gender,
			Role:     int32(user.Role),
			Birthday: user.Birthday,
		},
	}
	return userInfoResp, nil
}
func (u *UserService) GetUserList(context.Context, *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
}
func (u *UserService) GetUserByMobile(context.Context, *pb.GetUserByMobileRequest) (*pb.GetUserByMobileReply, error) {
}
func (u *UserService) GetUserById(context.Context, *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
}
func (u *UserService) UpdateUser(context.Context, *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
}
func (u *UserService) CheckPassword(context.Context, *pb.CheckPasswordRequest) (*pb.CheckPasswordReply, error) {
}

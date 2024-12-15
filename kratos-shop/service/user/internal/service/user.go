package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
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

func (u *UserService) GetUserList(ctx context.Context, req *pb.GetUserListRequest) (*pb.GetUserListReply, error) {
	list, total, err := u.uc.List(ctx, int(req.PageNumber), int(req.PageSize))
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetUserListReply{
		Total: total,
	}
	for _, user := range list {
		userInfoResp := toUserInfo(user)
		rsp.Data = append(rsp.Data, userInfoResp)
	}
	return rsp, nil
}

func (u *UserService) GetUserByMobile(ctx context.Context, req *pb.GetUserByMobileRequest) (*pb.GetUserByMobileReply, error) {
	user, err := u.uc.UserByMobile(ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetUserByMobileReply{
		UserInfo: toUserInfo(user),
	}
	return rsp, nil
}

func (u *UserService) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdReply, error) {
	user, err := u.uc.UserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	rsp := &pb.GetUserByIdReply{
		UserInfo: toUserInfo(user),
	}
	return rsp, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	birthDay := time.Unix(int64(req.Birthday), 0)
	ok, err := u.uc.UpdateUser(ctx, &biz.User{
		ID:       req.Id,
		Gender:   req.Gender,
		Birthday: &birthDay,
		NickName: req.NickName,
	})

	rsp := &pb.UpdateUserReply{IsOk: ok}
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (u *UserService) CheckPassword(ctx context.Context, req *pb.CheckPasswordRequest) (*pb.CheckPasswordReply, error) {
	check, err := u.uc.CheckPassword(ctx, req.Password, req.EncryptedPassword)
	if err != nil {
		return nil, err
	}
	return &pb.CheckPasswordReply{Success: check}, nil
}

func toUserInfo(user *biz.User) *pb.UserInfo {
	ret := &pb.UserInfo{
		Id:       user.ID,
		Mobile:   user.Mobile,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		ret.Birthday = uint64(user.Birthday.Unix())
	}
	return ret
}

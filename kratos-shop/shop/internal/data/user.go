package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	userService "shop/api/service/user/v1"
	"shop/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u userRepo) CreateUser(c context.Context, user *biz.User) (*biz.User, error) {
	createUser, err := u.data.uc.CreateUser(c, &userService.CreateUserRequest{
		NickName: user.NickName,
		Password: user.Password,
		Mobile:   user.Mobile,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       createUser.UserInfo.Id,
		NickName: createUser.UserInfo.NickName,
		Mobile:   createUser.UserInfo.Mobile,
	}, nil
}

func (u userRepo) UserByMobile(c context.Context, mobile string) (*biz.User, error) {
	byMobile, err := u.data.uc.GetUserByMobile(c, &userService.GetUserByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		return nil, err
	}

	return &biz.User{
		Mobile:   byMobile.UserInfo.Mobile,
		ID:       byMobile.UserInfo.Id,
		NickName: byMobile.UserInfo.NickName,
	}, nil
}

func (u userRepo) UserById(c context.Context, id int64) (*biz.User, error) {
	user, err := u.data.uc.GetUserById(c, &userService.GetUserByIdRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       id,
		Mobile:   user.UserInfo.Mobile,
		NickName: user.UserInfo.NickName,
		Gender:   user.UserInfo.Gender,
		Role:     int(user.UserInfo.Role),
	}, nil
}

func (u userRepo) CheckPassword(c context.Context, password, encryptedPassword string) (bool, error) {
	if byMobile, err := u.data.uc.CheckPassword(c, &userService.CheckPasswordRequest{Password: password, EncryptedPassword: encryptedPassword}); err != nil {
		return false, err
	} else {
		return byMobile.Success, nil
	}
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

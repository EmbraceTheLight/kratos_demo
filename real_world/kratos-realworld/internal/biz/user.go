package biz

import "context"

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type User struct {
	Username string
}

type UserUsecase struct {
	userRepo UserRepo
}

func NewUserUsecase(repo UserRepo) *UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (uc *UserUsecase) Register(ctx context.Context, user *User) error {
	if err := uc.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

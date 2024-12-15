package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int64
	Mobile    string
	Password  string
	NickName  string
	Birthday  *time.Time
	Gender    string
	Role      int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

// 新增mock数据命令
//
//go:generate mockgen -destination=../mocks/mrepo/user.go -package=mrepo . UserRepo
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) (u *User, err error)
	ListUser(ctx context.Context, pageNum, pageSize int64) (users []*User, total int, err error)
	UserByMobile(ctx context.Context, mobile string) (user *User, err error)
	UserByID(ctx context.Context, id int64) (user *User, err error)
	UpdateUser(ctx context.Context, user *User) (isSuccess bool, err error)
	CheckPassword(ctx context.Context, password, encryptedPassword string) (check bool, err error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Create(ctx context.Context, user *User) (*User, error) {
	return uc.repo.CreateUser(ctx, user)
}

func (uc *UserUsecase) List(ctx context.Context, pageNum, pageSize int64) ([]*User, int, error) {
	return uc.repo.ListUser(ctx, pageNum, pageSize)
}

func (uc *UserUsecase) GetByMobile(ctx context.Context, mobile string) (*User, error) {
	return uc.repo.UserByMobile(ctx, mobile)
}

func (uc *UserUsecase) GetByID(ctx context.Context, id int64) (*User, error) {
	return uc.repo.UserByID(ctx, id)
}

func (uc *UserUsecase) Update(ctx context.Context, user *User) (bool, error) {
	return uc.repo.UpdateUser(ctx, user)
}

func (uc *UserUsecase) CheckPassword(ctx context.Context, password, encryptedPassword string) (bool, error) {
	return uc.repo.CheckPassword(ctx, password, encryptedPassword)
}

package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"kratos-realworld/internal/conf"
	"kratos-realworld/internal/pkg/middleware/auth"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserBasic struct {
	Username string
	Email    string
	Bio      string
	Image    string
}

type UserLogin struct {
	*UserBasic
	Token string
}
type User struct {
	*UserBasic
	PasswordHash string
}

type UserUsecase struct {
	userRepo UserRepo
	jwt      *conf.JWT
	logger   *log.Helper
}

func NewUserUsecase(repo UserRepo, jwt *conf.JWT, logger log.Logger) *UserUsecase {
	return &UserUsecase{
		userRepo: repo,
		jwt:      jwt,
		logger:   log.NewHelper(logger),
	}
}

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	u := &User{
		UserBasic: &UserBasic{
			Username: username,
			Email:    email,
		},
		PasswordHash: hashPassword(password),
	}
	if err := uc.userRepo.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	return &UserLogin{
		UserBasic: u.UserBasic,
		Token:     uc.generateToken(u.Username),
	}, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {
	u, err := uc.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New("invalid password")
	}
	return &UserLogin{
		UserBasic: &UserBasic{
			Username: u.Username,
			Email:    u.Email,
			Bio:      u.Bio,
			Image:    u.Image,
		},
		Token: uc.generateToken(u.Username),
	}, nil
}

func (uc *UserUsecase) generateToken(username string) string {
	return auth.GenerateToken(uc.jwt.Secret, username)
}

func hashPassword(pwd string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}
	return true
}

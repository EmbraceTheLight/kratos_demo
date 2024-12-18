package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt2 "github.com/golang-jwt/jwt/v5"
	pb "shop/api/shop/v1"
	"shop/internal/conf"
	"shop/internal/pkg/captcha"
	"shop/internal/pkg/middleware/auth"
	"time"
)

var (
	ErrUsernameInvalid     = errors.New("username invalid")
	ErrPasswordInvalid     = errors.New("password invalid")
	ErrCaptchaInvalid      = errors.New("verification code error")
	ErrMobileInvalid       = errors.New("mobile invalid")
	ErrUserNotFound        = errors.New("user not found")
	ErrLoginFailed         = errors.New("login failed")
	ErrGenerateTokenFailed = errors.New("generate token failed")
	ErrAuthFailed          = errors.New("authentication failed")
)

type User struct {
	ID        int64
	Mobile    string
	NickName  string
	Password  string
	Birthday  int64
	Gender    string
	Role      int
	CreatedAt time.Time
}

type UserRepo interface {
	CreateUser(c context.Context, user *User) (*User, error)
	UserByMobile(c context.Context, mobile string) (*User, error)
	UserById(c context.Context, id int64) (*User, error)
	CheckPassword(c context.Context, password, encryptedPassword string) (bool, error)
}

type UserUsecase struct {
	repo       UserRepo
	log        *log.Helper
	signingKey string // 这个字段是为了生成 token 时可以直接取配置文件里的配置
}

func NewUserUsecase(repo UserRepo, logger log.Logger, conf *conf.Auth) *UserUsecase {
	return &UserUsecase{
		repo:       repo,
		log:        log.NewHelper(log.With(logger, "module", "usecase/shop")),
		signingKey: conf.JwtKey,
	}
}

// GetCaptcha 生成验证码
func (uc *UserUsecase) GetCaptcha(ctx context.Context) (*pb.CaptchaReply, error) {
	captchaInfo, err := captcha.GetCaptcha(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.CaptchaReply{
		CaptchaId: captchaInfo.CaptchaID,
		Picture:   captchaInfo.Picture,
		Answer:    captchaInfo.Answer,
	}, nil
}

// UserDetailByID 获取用户详情
func (uc *UserUsecase) UserDetailByID(ctx context.Context) (*pb.UserDetailReply, error) {
	//通过context上下文获取用户id
	var userId int64
	if claims, ok := jwt.FromContext(ctx); ok {
		c := claims.(jwt2.MapClaims)
		if c["ID"] == nil {
			return nil, ErrAuthFailed
		}
		userId = int64(c["ID"].(float64))
	}

	user, err := uc.repo.UserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &pb.UserDetailReply{
		Id:       user.ID,
		Mobile:   user.Mobile,
		NickName: user.NickName,
	}, nil
}

func (uc *UserUsecase) PasswordLogin(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	if len(req.Mobile) <= 0 {
		return nil, ErrMobileInvalid
	}
	if len(req.Password) <= 0 {
		return nil, ErrPasswordInvalid
	}

	//验证验证码是否正确
	if !captcha.Store.Verify(req.CaptchaId, req.Captcha, true) {
		return nil, ErrCaptchaInvalid
	}

	if user, err := uc.repo.UserByMobile(ctx, req.Mobile); err != nil {
		return nil, ErrUserNotFound
	} else {
		//用户存在，检查密码
		if isPsdCorrect, pasErr := uc.repo.CheckPassword(ctx, req.Password, user.Password); pasErr != nil {
			return nil, ErrPasswordInvalid
		} else { //密码正确
			if isPsdCorrect {
				claims := &auth.CustomClaims{
					ID:          user.ID,
					NickName:    user.NickName,
					AuthorityID: user.Role,
					RegisteredClaims: jwt2.RegisteredClaims{
						NotBefore: jwt2.NewNumericDate(time.Now()),                          //签名生效时间
						ExpiresAt: jwt2.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), //token过期时间: 30天
						Issuer:    "zey",
					},
				}

				token, err := auth.CreateToken(claims, uc.signingKey)
				if err != nil {
					return nil, ErrGenerateTokenFailed
				}
				return &pb.LoginReply{
					Id:       user.ID,
					Mobile:   user.Mobile,
					Username: user.NickName,
					Token:    token,
					ExpireAt: claims.ExpiresAt.Unix(),
				}, nil
			} else { //密码错误
				return nil, ErrLoginFailed
			}
		}
	}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	nu, err := newUser(req.Mobile, req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	createUser, err := uc.repo.CreateUser(ctx, nu)
	if err != nil {
		return nil, err
	}
	claims := &auth.CustomClaims{
		ID:          createUser.ID,
		NickName:    createUser.NickName,
		AuthorityID: createUser.Role,
		RegisteredClaims: jwt2.RegisteredClaims{
			NotBefore: jwt2.NewNumericDate(time.Now()),                          //签名生效时间
			ExpiresAt: jwt2.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)), //token过期时间: 30天
			Issuer:    "zey",
		},
	}
	token, err := auth.CreateToken(claims, uc.signingKey)
	if err != nil {
		return nil, ErrGenerateTokenFailed
	}

	return &pb.RegisterReply{
		Id:       createUser.ID,
		Mobile:   createUser.Mobile,
		Username: createUser.NickName,
		Token:    token,
		ExpireAt: claims.ExpiresAt.Unix(),
	}, nil
}

func newUser(mobile, username, password string) (*User, error) {
	if len(mobile) <= 0 {
		return nil, ErrMobileInvalid
	}
	if len(username) <= 0 {
		return nil, ErrUsernameInvalid
	}
	if len(password) <= 0 {
		return nil, ErrPasswordInvalid
	}
	return &User{
		Mobile:   mobile,
		NickName: username,
		Password: password,
	}, nil
}

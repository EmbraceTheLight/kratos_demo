package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "shop/api/shop/v1"
	"shop/internal/biz"
)

type ShopService struct {
	pb.UnimplementedShopServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewShopService(uc *biz.UserUsecase, logger log.Logger) *ShopService {
	return &ShopService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/shop")),
	}
}

func (s *ShopService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	return s.uc.CreateUser(ctx, req)
}

func (s *ShopService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	return s.uc.PasswordLogin(ctx, req)
}

func (s *ShopService) Captcha(ctx context.Context, req *emptypb.Empty) (*pb.CaptchaReply, error) {
	return s.uc.GetCaptcha(ctx)
}

func (s *ShopService) Detail(ctx context.Context, req *emptypb.Empty) (*pb.UserDetailReply, error) {
	return s.uc.UserDetailByID(ctx)
}

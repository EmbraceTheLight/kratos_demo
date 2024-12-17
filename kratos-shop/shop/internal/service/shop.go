package service

import (
	"github.com/go-kratos/kratos/v2/log"
	pb "shop/api/shop/v1"
)

type ShopService struct {
	pb.UnimplementedShopServer

	uc  *biz.UserUsecase
	log *log.Helper
}

func NewShopService(uc *biz.UserUsecase, logger log.Logger) *ShopService {
	return &ShopService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

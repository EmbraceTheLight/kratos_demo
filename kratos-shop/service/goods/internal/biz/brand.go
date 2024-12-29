package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type BrandRepo interface {
	IsBrandByID(context.Context, int32) (*domain.Brand, error)
}

type BrandUsecase struct {
	repo BrandRepo
	log  *log.Helper
}

func NewBrandUsecase(repo BrandRepo, logger log.Logger) *BrandUsecase {
	return &BrandUsecase{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

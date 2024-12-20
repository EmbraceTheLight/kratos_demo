package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "goods/api/goods/v1"
	"goods/internal/biz"
)

type GoodsService struct {
	v1.UnimplementedGoodsServer
	category *biz.CategoryUsecase
	logger   *log.Helper
}

func NewCategoryService(categoryBiz *biz.CategoryUsecase, logger log.Logger) *GoodsService {
	return &GoodsService{
		category: categoryBiz,
		logger:   log.NewHelper(log.With(logger, "module", "service/category")),
	}
}

func (gs *GoodsService) CreateCategory(ctx context.Context, req *v1.CreateCategoryReq) (*v1.CreateCategoryResp, error) {
	result, err := gs.category.CreateCategory(ctx, &biz.CategoryInfo{
		Name:           req.Name,
		ParentCategory: req.ParentCategory,
		Level:          req.Level,
		IsTab:          req.IsTab,
		Sort:           req.Sort,
	})

	if err != nil {
		return nil, err
	}
	return &v1.CreateCategoryResp{
		Id:             result.ID,
		Name:           result.Name,
		ParentCategory: result.ParentCategory,
		Level:          result.Level,
		IsTab:          result.IsTab,
		Sort:           result.Sort,
	}, nil
}

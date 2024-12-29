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
	types    *biz.GoodsTypeUsecase
	goods    *biz.GoodsUsecase
	attr     *biz.GoodsAttrUsecase
	spec     *biz.SpecificationUsecase
	logger   *log.Helper
}

func NewGoodsService(
	category *biz.CategoryUsecase,
	types *biz.GoodsTypeUsecase,
	attr *biz.GoodsAttrUsecase,
	goods *biz.GoodsUsecase,
	spec *biz.SpecificationUsecase,
	logger log.Logger) *GoodsService {
	return &GoodsService{
		category: category,
		types:    types,
		goods:    goods,
		attr:     attr,
		spec:     spec,
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

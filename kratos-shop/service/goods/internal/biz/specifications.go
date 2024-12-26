package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type SpecificationRepo interface {
	CreateSpecification(context.Context, *domain.Specification) (int64, error)
	CreateSpecificationValue(context.Context, int64, []*domain.SpecificationValue) error
}

type SpecificationUsecase struct {
	repo  SpecificationRepo
	gRepo GoodsTypeRepo
	tx    Transaction
	log   *log.Helper
}

func NewSpecificationUsecase(repo SpecificationRepo, gRepo GoodsTypeRepo, tx Transaction, logger log.Logger) *SpecificationUsecase {
	return &SpecificationUsecase{
		repo:  repo,
		gRepo: gRepo,
		tx:    tx,
		log:   log.NewHelper(logger),
	}
}

func (s *SpecificationUsecase) CreateSpecification(ctx context.Context, r *domain.Specification) (int64, error) {
	var (
		id  int64
		err error
	)

	//domain下编写的判断type id是否为空
	if r.IsTypeIDEmpty() {
		return 0, errors.New("请选择商品类型进行绑定")
	}

	//判断规格名称是否为空
	if r.IsValueEmpty() {
		return 0, errors.New("请添加商品规格参数")
	}

	//查询有没有这个商品类型
	_, err = s.gRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return id, err
	}

	//引入事务
	err = s.tx.ExecTx(ctx, func(ctx context.Context) error {
		//创建规格
		id, err = s.repo.CreateSpecification(ctx, r)
		if err != nil {
			return err
		}

		//插入规格对应的值
		err = s.repo.CreateSpecificationValue(ctx, id, r.SpecificationValues)
		if err != nil {
			return err
		}
		return nil
	})
	return id, err
}

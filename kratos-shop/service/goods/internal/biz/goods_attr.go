package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GoodsAttrRepo interface {
	CreateGoodsGroupAttr(context.Context, *domain.AttrGroup) (*domain.AttrGroup, error)
	IsExistsGroupByID(context.Context, int64) (*domain.AttrGroup, error)
	CreateGoodsAttr(context.Context, *domain.GoodsAttr) (*domain.GoodsAttr, error)
	CreateGoodsAttrValue(context.Context, []*domain.GoodsAttrValue) ([]*domain.GoodsAttrValue, error)
	ListByIds(ctx context.Context, id ...int64) (domain.GoodsAttrList, error)
}

type GoodsAttrUsecase struct {
	repo     GoodsAttrRepo
	typeRepo GoodsTypeRepo
	tx       Transaction
	log      *log.Helper
}

func NewGoodsAttrUsecase(repo GoodsAttrRepo, typeRepo GoodsTypeRepo, tx Transaction, logger log.Logger) *GoodsAttrUsecase {
	return &GoodsAttrUsecase{
		repo:     repo,
		typeRepo: typeRepo,
		tx:       tx,
		log:      log.NewHelper(logger),
	}
}

func (u *GoodsAttrUsecase) CreateAttrGroup(ctx context.Context, r *domain.AttrGroup) (*domain.AttrGroup, error) {
	// 查询是否指定了商品类型ID
	if r.IsTypeIDEmpty() {
		return nil, errors.New("请选择商品类型进行绑定")
	}

	// 根据商品类型ID检查商品类型是否存在
	_, err := u.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	// 创建属性组
	attr, err := u.repo.CreateGoodsGroupAttr(ctx, r)
	if err != nil {
		return nil, err
	}
	return attr, nil
}

func (u *GoodsAttrUsecase) CreateAttrValue(ctx context.Context, r *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	var (
		attrInfo  *domain.GoodsAttr
		attrValue []*domain.GoodsAttrValue
		err       error
	)
	//检查商品类型ID是否为空
	if r.IsTypeIDEmpty() {
		return nil, errors.New("请选择商品类型进行绑定")
	}

	//检查商品类型是否存在
	_, err = u.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, err
	}

	//检查属性组是否存在
	_, err = u.repo.IsExistsGroupByID(ctx, r.GroupID)
	if err != nil {
		return nil, err
	}

	//引入事务，创建商品属性和属性值
	err = u.tx.ExecTx(ctx, func(ctx context.Context) error {
		//创建商品属性信息
		attrInfo, err = u.repo.CreateGoodsAttr(ctx, r)
		if err != nil {
			return err
		}

		//获取并遍历要添加的商品属性值
		var values []*domain.GoodsAttrValue
		for _, v := range r.GoodsAttrValue {
			if v.IsValueEmpty() {
				return errors.New("商品属性值不能为空")
			}
			res := &domain.GoodsAttrValue{
				AttrID:  attrInfo.ID,
				GroupID: attrInfo.GroupID,
				Value:   v.Value,
			}
			values = append(values, res)
		}
		//为商品属性添加属性值
		attrValue, err = u.repo.CreateGoodsAttrValue(ctx, values)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return &domain.GoodsAttr{
		ID:             attrInfo.ID,
		TypeID:         attrInfo.TypeID,
		GroupID:        attrInfo.GroupID,
		Title:          attrInfo.Title,
		Sort:           attrInfo.Sort,
		Status:         attrInfo.Status,
		Desc:           attrInfo.Desc,
		GoodsAttrValue: attrValue,
	}, nil
}

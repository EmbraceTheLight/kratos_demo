package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// GoodsType 商品类型表
type GoodsType struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:商品类型名称" json:"name"`
	TypeCode  string         `gorm:"type:varchar(50);not null;comment:商品类型编码" json:"type_code"`
	AliasName string         `gorm:"type:varchar(50);not null;comment:商品类型别名" json:"alias_name"`
	IsVirtual bool           `gorm:"comment:是否是虚拟商品显示;default:false" json:"is_virtual"`
	Desc      string         `gorm:"type:varchar(50);not null;comment:商品类型描述" json:"desc"`
	Sort      int32          `gorm:"comment:类型排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type goodsTypeRepo struct {
	data *Data
	log  *log.Helper
}

func NewGoodsTypeRepo(data *Data, logger log.Logger) biz.GoodsTypeRepo {
	return &goodsTypeRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *goodsTypeRepo) CreateGoodsType(ctx context.Context, req *domain.GoodsType) (int64, error) {
	gt := GoodsType{
		Name:      req.Name,
		TypeCode:  req.TypeCode,
		AliasName: req.AliasName,
		IsVirtual: req.IsVirtual,
		Desc:      req.Desc,
		Sort:      req.Sort,
	}
	result := g.data.db.Save(&gt)
	return gt.ID, result.Error
}

func (g *goodsTypeRepo) IsExistsByID(ctx context.Context, typeID int64) (*domain.GoodsType, error) {
	var goodsType GoodsType
	if res := g.data.db.First(&goodsType, typeID); res.RowsAffected == 0 {
		return nil, errors.New("商品类型不存在")
	}
	return goodsType.ToDomain(), nil
}

func (g *GoodsType) ToDomain() *domain.GoodsType {
	return &domain.GoodsType{
		ID:        g.ID,
		Name:      g.Name,
		TypeCode:  g.TypeCode,
		AliasName: g.AliasName,
		IsVirtual: g.IsVirtual,
		Desc:      g.Desc,
		Sort:      g.Sort,
	}
}

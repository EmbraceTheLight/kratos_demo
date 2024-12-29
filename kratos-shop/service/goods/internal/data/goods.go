package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
)

// Goods 商品表
type Goods struct {
	BaseFields
	CategoryID int32 `gorm:"index:category_id;type:int;comment:分类ID;not null"`
	BrandsID   int32 `gorm:"index:brand_id;type:int;comment:品牌ID ;not null"`
	TypeID     int64 `gorm:"index:type_id;type:int;comment:商品类型ID ;not null"`

	Name            string   `gorm:"type:varchar(100);not null;comment:商品名称"`
	AliasName       string   `gorm:"type:varchar(100);not null;comment:商品别名"`
	GoodsSn         string   `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsTags       string   `gorm:"type:varchar(100);not null;comment:商品标签"`
	MarketPrice     int64    `gorm:"type:int;default:0;not null;comment:商品展示价格"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品封面图"`
	GoodsImages     GormList `gorm:"type:varchar(1000);not null;comment:商品的介绍图"` // 切片类型转为 json 到数据库，取出来是切片类型

	OnSale   bool  `gorm:"default:false;comment:是否上架;not null "`
	ShipFree bool  `gorm:"default:false;comment:是否免运费; not null"`
	ShipID   int32 `gorm:"type:int;comment:运费模版ID;not null"`
	IsNew    bool  `gorm:"default:false;comment:是否新品;not null"`
	IsHot    bool  `gorm:"comment:是否热卖商品;default:false;not null"`

	ClickNum int64 `gorm:"default:0;type:int; comment 商品详情点击数"`
	SoldNum  int64 `gorm:"default:0;type:int; comment 商品销售数"`
	FavNum   int64 `gorm:"default:0;type:int; comment 商品收藏数"`

	// 售前服务信息、售后服务信息、商品促销活动信息
}

type goodsRepo struct {
	data *Data
	log  *log.Helper
}

func (g *goodsRepo) CreateGoods(ctx context.Context, goods *domain.Goods) (*domain.Goods, error) {
	product := &Goods{
		CategoryID:      goods.CategoryID,
		BrandsID:        goods.BrandsID,
		TypeID:          goods.TypeID,
		Name:            goods.Name,
		AliasName:       goods.AliasName,
		GoodsSn:         goods.GoodsSn,
		GoodsTags:       goods.GoodsTags,
		MarketPrice:     goods.MarketPrice,
		GoodsBrief:      goods.GoodsBrief,
		GoodsFrontImage: goods.GoodsFrontImage,
		GoodsImages:     goods.GoodsImages,
		OnSale:          goods.OnSale,
		ShipFree:        goods.ShipFree,
		ShipID:          goods.ShipID,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
	}
	result := g.data.DB(ctx).Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product.ToDomain(), nil
}

// NewGoodsRepo .
func NewGoodsRepo(data *Data, logger log.Logger) biz.GoodsRepo {
	return &goodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *Goods) ToDomain() *domain.Goods {
	return &domain.Goods{
		ID:              g.ID,
		CategoryID:      g.CategoryID,
		BrandsID:        g.BrandsID,
		TypeID:          g.TypeID,
		Name:            g.Name,
		AliasName:       g.AliasName,
		GoodsSn:         g.GoodsSn,
		GoodsTags:       g.GoodsTags,
		MarketPrice:     g.MarketPrice,
		GoodsBrief:      g.GoodsBrief,
		GoodsFrontImage: g.GoodsFrontImage,
		GoodsImages:     g.GoodsImages,
		OnSale:          g.OnSale,
		ShipFree:        g.ShipFree,
		ShipID:          g.ShipID,
		IsNew:           g.IsNew,
		IsHot:           g.IsHot,
		ClickNum:        g.ClickNum,
		SoldNum:         g.SoldNum,
		FavNum:          g.FavNum,
	}
}

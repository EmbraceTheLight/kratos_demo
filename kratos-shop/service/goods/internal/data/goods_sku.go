package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
)

// GoodsSku 商品SKU 表
type GoodsSku struct {
	BaseFields
	GoodsID        int64  `gorm:"index:goods_id;type:int;comment:商品ID;not null"`
	GoodsSn        string `gorm:"type:varchar(100);not null;comment:商品编号"`
	GoodsName      string `gorm:"type:varchar(100);not null;comment:商品名称"`
	SkuName        string `gorm:"type:varchar(100);comment:SKU名称;not null"`
	SkuCode        string `gorm:"type:varchar(100);comment:SKUCode;not null"`
	BarCode        string `gorm:"type:varchar(100);comment:条码;not null"`
	Price          int64  `gorm:"type:int;comment:商品售价;not null"`
	PromotionPrice int64  `gorm:"type:int;comment:商品促销售价;not null"`
	Points         int64  `gorm:"type:int;comment:赠送积分;not null"`
	RemarksInfo    string `gorm:"type:varchar(100);comment:备注信息;not null"`
	Pic            string `gorm:"type:varchar(500);not null;comment:规格参数对应的图片" json:"pic"`
	OnSale         bool   `gorm:"comment:是否上架;default:false;not null"`
	AttrInfo       string `gorm:"type:varchar(2000);comment:商品属性信息JSON;not null"`
	Inventory      int64  `gorm:"type:int;comment:商品SKU库存冗余字段;not null"`
}

// GoodsSpecificationSku 商品规格和商品Sku关联表
type GoodsSpecificationSku struct {
	BaseFields
	SkuID           int64  `gorm:"index:sku_id;type:int;comment:商品SKU_ID;not null"`
	SkuCode         string `gorm:"type:varchar(100);comment:商品SKU_Code;not null"`
	SpecificationId int64  `gorm:"index:specification_id;type:int;comment:商品规格ID;not null"`
	ValueId         int64  `gorm:"index:value_id;type:int;comment:商品规格值表ID;not null"`
}

type goodsSkuRepo struct {
	data *Data
	log  *log.Helper
}

func (r goodsSkuRepo) Create(ctx context.Context, req *domain.GoodsSku) (*domain.GoodsSku, error) {
	sku := &GoodsSku{
		GoodsID:        req.GoodsID,
		GoodsSn:        req.GoodsSn,
		GoodsName:      req.GoodsName,
		SkuName:        req.SkuName,
		SkuCode:        req.SkuCode,
		BarCode:        req.BarCode,
		Price:          req.Price,
		PromotionPrice: req.PromotionPrice,
		Points:         req.Points,
		RemarksInfo:    req.RemarksInfo,
		Pic:            req.Pic,
		OnSale:         req.OnSale,
		AttrInfo:       req.AttrInfo,
		Inventory:      req.Inventory,
	}

	if err := r.data.DB(ctx).Save(sku).Error; err != nil {
		return nil, err
	}
	return sku.ToDomain(), nil
}

func (r goodsSkuRepo) CreateSkuRelation(ctx context.Context, req []*domain.GoodsSpecificationSku) error {
	var info []*GoodsSpecificationSku
	for _, sku := range req {
		i := GoodsSpecificationSku{
			SkuID:           sku.SkuID,
			SkuCode:         sku.SkuCode,
			SpecificationId: sku.SpecificationID,
			ValueId:         sku.ValueID,
		}
		info = append(info, &i)
	}
	result := r.data.DB(ctx).Model(&GoodsSpecificationSku{}).Save(&info)
	return result.Error
}

func NewGoodsSkuRepo(data *Data, logger log.Logger) biz.GoodsSkuRepo {
	return &goodsSkuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (sku *GoodsSku) ToDomain() *domain.GoodsSku {
	return &domain.GoodsSku{
		ID:             sku.ID,
		GoodsID:        sku.GoodsID,
		GoodsSn:        sku.GoodsSn,
		GoodsName:      sku.GoodsName,
		SkuName:        sku.SkuName,
		SkuCode:        sku.SkuCode,
		BarCode:        sku.BarCode,
		Price:          sku.Price,
		PromotionPrice: sku.PromotionPrice,
		Points:         sku.Points,
		RemarksInfo:    sku.RemarksInfo,
		Pic:            sku.Pic,
		Inventory:      sku.Inventory,
		OnSale:         sku.OnSale,
		AttrInfo:       sku.AttrInfo,
	}
}

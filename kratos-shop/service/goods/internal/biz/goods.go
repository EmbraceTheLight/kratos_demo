package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type GoodsRepo interface {
	CreateGoods(ctx context.Context, goods *domain.Goods) (*domain.Goods, error)
	GoodsListByIDs(context.Context, ...int64) ([]*domain.Goods, error)
}

type GoodsUsecase struct {
	repo              GoodsRepo
	tr                Transaction
	skuRepo           GoodsSkuRepo
	categoryRepo      CategoryRepo
	brandRepo         BrandRepo
	typeRepo          GoodsTypeRepo
	specificationRepo SpecificationRepo
	goodsAttrRepo     GoodsAttrRepo
	esGoodsRepo       EsGoodsRepo
	inventoryRepo     InventoryRepo
	log               *log.Helper
}

func NewGoodsUsecase(
	repo GoodsRepo,
	tr Transaction,
	skuRepo GoodsSkuRepo,
	categoryRepo CategoryRepo,
	brandRepo BrandRepo,
	typeRepo GoodsTypeRepo,
	specificationRepo SpecificationRepo,
	goodsAttrRepo GoodsAttrRepo,
	esGoodsRepo EsGoodsRepo,
	inventoryRepo InventoryRepo,
	logger log.Logger) *GoodsUsecase {
	return &GoodsUsecase{
		repo:              repo,
		tr:                tr,
		skuRepo:           skuRepo,
		categoryRepo:      categoryRepo,
		brandRepo:         brandRepo,
		typeRepo:          typeRepo,
		specificationRepo: specificationRepo,
		goodsAttrRepo:     goodsAttrRepo,
		esGoodsRepo:       esGoodsRepo,
		inventoryRepo:     inventoryRepo,
		log:               log.NewHelper(logger),
	}

}

func (u *GoodsUsecase) CreateGoods(ctx context.Context, r *domain.Goods) (*domain.GoodsInfoResponse, error) {
	var (
		err     error
		goods   *domain.Goods
		esGoods *domain.ESGoods
	)
	// 判断品牌是否存在
	brand, err := u.brandRepo.IsBrandByID(ctx, r.BrandsID)
	if err != nil {
		return nil, err
	}

	cate, err := u.categoryRepo.GetCategoryByID(ctx, r.CategoryID)
	if err != nil {
		return nil, errors.New("分类不存在")
	}

	//判断商品类型是否存在
	goodsType, err := u.typeRepo.IsExistsByID(ctx, r.TypeID)
	if err != nil {
		return nil, errors.New("商品类型不存在")
	}

	//判断商品规格和属性是否存在
	for _, sku := range r.Sku {

		//遍历请求的商品规格ID
		var sIDs []int64
		for _, info := range sku.Specification {
			sIDs = append(sIDs, info.SpecificationsID)
		}

		//根据遍历到的规格ID查询规格是否存在
		specList, err := u.specificationRepo.ListByIds(ctx, sIDs...)
		if err != nil {
			return nil, errors.New("商品规格不存在")
		}
		for _, sID := range sIDs {
			info := specList.FindById(sID)
			if info == nil {
				return nil, errors.New("商品规格不存在")
			}
		}

		//遍历请求的商品属性ID
		var attrIDs []int64
		for _, attr := range sku.GroupAttr {
			for _, a := range attr.Attr {
				attrIDs = append(attrIDs, a.AttrID)
			}
		}
		//根据遍历到的属性ID查询属性是否存在
		attrList, err := u.goodsAttrRepo.ListByIds(ctx, attrIDs...)
		if err != nil {
			return nil, errors.New("商品属性不存在")
		}

		//遍历请求的商品属性组
		for _, attrGroup := range sku.GroupAttr {
			//遍历某个属性组的属性
			for _, attr := range attrGroup.Attr {
				attrIDs = append(attrIDs, attr.AttrID)
				notExist := attrList.IsNotExist(attrGroup.GroupID, attr.AttrID)
				if notExist {
					return nil, errors.New("商品属性不存在")
				}

			}
		}
	}

	err = u.tr.ExecTx(ctx, func(ctx context.Context) error {
		goods, err = u.repo.CreateGoods(ctx, &domain.Goods{
			CategoryID:      r.CategoryID,
			BrandsID:        r.BrandsID,
			TypeID:          r.TypeID,
			Name:            r.Name,
			AliasName:       r.AliasName,
			GoodsSn:         r.GoodsSn,
			GoodsTags:       r.GoodsTags,
			MarketPrice:     r.MarketPrice,
			GoodsBrief:      r.GoodsBrief,
			GoodsFrontImage: r.GoodsFrontImage,
			GoodsImages:     r.GoodsImages,
			OnSale:          r.OnSale,
			IsNew:           r.IsNew,
			IsHot:           r.IsHot,
			ShipFree:        r.ShipFree,
			ShipID:          r.ShipID,
		})
		if err != nil {
			return err
		}

		// 更新商品 SKU 表
		for _, v := range r.Sku {
			res := &domain.GoodsSku{
				GoodsID:        goods.ID,
				GoodsSn:        goods.GoodsSn,
				GoodsName:      goods.Name,
				SkuName:        v.SkuName,
				SkuCode:        v.SkuCode,
				BarCode:        v.BarCode,
				Price:          v.Price,
				PromotionPrice: v.PromotionPrice,
				Points:         v.Points,
				RemarksInfo:    v.RemarksInfo,
				Pic:            v.Pic,
				Inventory:      v.Inventory,
				OnSale:         v.OnSale,
			}

			goodsAttr, err := json.Marshal(v.GroupAttr)
			if err != nil {
				return err
			}
			res.AttrInfo = string(goodsAttr)

			// 插入 sku 表
			skuInfo, err := u.skuRepo.Create(ctx, res)
			if err != nil {
				return err
			}

			// 插入库存表
			_, err = u.inventoryRepo.Create(ctx, &domain.Inventory{
				SkuID:     skuInfo.ID,
				Inventory: skuInfo.Inventory,
			})
			if err != nil {
				return err
			}
			// 插入 sku 规格关联关系表
			var skuRelation []*domain.GoodsSpecificationSku
			for _, spec := range v.Specification {
				skuRelation = append(skuRelation, &domain.GoodsSpecificationSku{
					SkuID:           skuInfo.ID,
					SkuCode:         skuInfo.SkuCode,
					SpecificationID: spec.SpecificationsID,
					ValueID:         spec.SpecificationsValueID,
				})
			}

			// 插入商品规格关联关系表
			err = u.skuRepo.CreateSkuRelation(ctx, skuRelation)
			if err != nil {
				return err
			}

			// esModel
			{
				esGoods = new(domain.ESGoods)
				esGoods.Sku = append(esGoods.Sku, domain.EsSku{
					SkuID:    skuInfo.ID,
					SkuName:  skuInfo.SkuName,
					SkuPrice: skuInfo.Price,
				})
				esGoods.BrandsID = brand.ID
				esGoods.BrandName = brand.Name
				esGoods.CategoryID = cate.ID
				esGoods.CategoryName = cate.Name
				esGoods.TypeID = goodsType.ID
				esGoods.TypeName = goodsType.Name
				esGoods.Name = goodsType.Name
				esGoods.ID = goods.ID
				esGoods.OnSale = goods.OnSale
				esGoods.ShipFree = goods.ShipFree
				esGoods.IsNew = goods.IsNew
				esGoods.IsHot = goods.IsHot
				esGoods.Name = goods.Name
				esGoods.GoodsTags = goods.GoodsTags
				esGoods.ClickNum = goods.ClickNum
				esGoods.SoldNum = goods.SoldNum
				esGoods.FavNum = goods.FavNum
				esGoods.MarketPrice = goods.MarketPrice
				esGoods.GoodsBrief = goods.GoodsBrief
			}

			err = u.esGoodsRepo.InsertEsGoods(ctx, esGoods)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &domain.GoodsInfoResponse{
		GoodsID: goods.ID,
	}, nil
}

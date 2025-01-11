package service

import (
	"context"
	pb "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateGoods 创建商品
func (gs *GoodsService) CreateGoods(ctx context.Context, r *pb.CreateGoodsReq) (*pb.CreateGoodsResp, error) {
	var goodsSku []*domain.GoodsSku

	// 处理、填充商品sku
	for _, sku := range r.Sku {
		res := &domain.GoodsSku{
			GoodsName:      r.Name,
			GoodsSn:        r.GoodsSn,
			SkuName:        sku.SkuName,
			BarCode:        sku.BarCode,
			Price:          sku.Price,
			PromotionPrice: sku.PromotionPrice,
			Points:         sku.Points,
			Pic:            sku.Image,
			Inventory:      sku.Inventory,
			OnSale:         r.OnSale,
		}

		// 填充sku中的商品规格信息
		for _, specification := range sku.SpecificationInfo {
			sinfo := &domain.SpecificationInfo{
				SpecificationsID:      specification.SId,
				SpecificationsValueID: specification.VId,
			}
			res.Specification = append(res.Specification, sinfo)
		}

		// 填充sku中的商品属性信息--遍历属性组
		for _, attrGroup := range sku.GroupAttrInfo {
			group := &domain.GroupAttr{
				GroupID:   attrGroup.GroupId,
				GroupName: attrGroup.GroupName,
			}
			// 遍历属性组中的属性
			for _, attr := range attrGroup.AttrInfo {
				a := &domain.Attr{
					AttrID:        attr.AttrId,
					AttrName:      attr.AttrName,
					AttrValueID:   attr.AttrValueId,
					AttrValueName: attr.AttrValueName,
				}
				group.Attr = append(group.Attr, a)
			}
			res.GroupAttr = append(res.GroupAttr, group)
		}
		goodsSku = append(goodsSku, res)
	}
	goodsInfo := &domain.Goods{
		ID:              r.Id,
		CategoryID:      r.CategoryId,
		BrandsID:        r.BrandId,
		TypeID:          r.TypeId,
		Name:            r.Name,
		AliasName:       r.AliasName,
		GoodsSn:         r.GoodsSn,
		GoodsTags:       r.GoodsTags,
		MarketPrice:     r.MarketPrice,
		GoodsBrief:      r.GoodsBrief,
		GoodsFrontImage: r.GoodsFrontImage,
		GoodsImages:     r.GoodsImages,
		OnSale:          r.OnSale,
		ShipFree:        r.ShipFree,
		ShipID:          r.ShipId,
		IsNew:           r.IsNew,
		IsHot:           r.IsHot,
		Sku:             goodsSku,
	}
	result, err := gs.goods.CreateGoods(ctx, goodsInfo)
	if err != nil {
		return nil, err
	}
	return &pb.CreateGoodsResp{Id: result.GoodsID}, nil
}

// GoodsList 通过elastic search查询商品
func (gs *GoodsService) GoodsList(ctx context.Context, r *pb.GoodsListReq) (*pb.GoodsListResp, error) {
	goodsFilter := &domain.ESGoodsFilter{
		ID:          r.Id,
		CategoryID:  r.CategoryId,
		BrandsID:    r.BrandId,
		Keywords:    r.Keywords,
		IsNew:       r.IsNew,
		IsHot:       r.IsHot,
		ClickNum:    r.ClickNum,
		SoldNum:     r.SoldNum,
		FavNum:      r.FavNum,
		MaxPrice:    r.MaxPrice,
		MinPrice:    r.MinPrice,
		Pages:       r.Pages,
		PagePerNums: r.PagePerNums,
	}

	result, err := gs.esGoods.GoodsList(ctx, goodsFilter)
	if err != nil {
		return nil, err
	}
	response := &pb.GoodsListResp{
		TotalCount: result.Total,
	}
	for _, goods := range result.List {
		res := &pb.GoodsListResp_GoodsInfo{
			Id:          goods.ID,
			CategoryId:  goods.CategoryID,
			BrandId:     goods.BrandsID,
			Name:        goods.Name,
			GoodsSn:     goods.GoodsSn,
			ClickNum:    goods.ClickNum,
			SoldNum:     goods.SoldNum,
			FavNum:      goods.FavNum,
			MarketPrice: goods.MarketPrice,
			GoodsBrief:  goods.GoodsBrief,
			GoodsDesc:   goods.GoodsBrief,
			ShipFree:    goods.ShipFree,
			Images:      goods.GoodsFrontImage,
			GoodsImages: goods.GoodsImages,
			IsNew:       goods.IsNew,
			IsHot:       goods.IsHot,
			OnSale:      goods.OnSale,
		}
		response.GoodsInfo = append(response.GoodsInfo, res)
	}
	return response, nil
}

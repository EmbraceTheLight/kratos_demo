package service

import (
	"context"
	pb "goods/api/goods/v1"
	"goods/internal/domain"
)

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

package service

import (
	"context"
	pb "goods/api/goods/v1"
	"goods/internal/domain"
)

func (gs *GoodsService) CreateGoodsType(ctx context.Context, req *pb.GoodsTypeReq) (*pb.GoodsTypeResp, error) {
	id, err := gs.types.GoodsTypeCreate(ctx, &domain.GoodsType{
		Name:      req.Name,
		TypeCode:  req.TypeCode,
		AliasName: req.AliasName,
		IsVirtual: req.IsVirtual,
		Desc:      req.Desc,
		Sort:      req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GoodsTypeResp{
		Id: id,
	}, nil
}

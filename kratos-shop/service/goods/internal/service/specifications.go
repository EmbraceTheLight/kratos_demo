package service

import (
	"context"
	pb "goods/api/goods/v1"
	"goods/internal/domain"
)

func (gs *GoodsService) CreateGoodsSpecification(ctx context.Context, r *pb.SpecificationReq) (*pb.SpecificationResp, error) {
	var value []*domain.SpecificationValue
	if r.SpecificationValues != nil {
		for _, v := range r.SpecificationValues {
			res := &domain.SpecificationValue{
				Value: v.Value,
				Sort:  v.Sort,
			}
			value = append(value, res)
		}
	}
	id, err := gs.spec.CreateSpecification(ctx, &domain.Specification{
		TypeID:              r.TypeId,
		Name:                r.Name,
		Sort:                r.Sort,
		Status:              r.Status,
		IsSKU:               r.IsSku,
		IsSelect:            r.IsSelect,
		SpecificationValues: value,
	})

	if err != nil {
		return nil, err
	}
	return &pb.SpecificationResp{
		Id: id,
	}, nil
}

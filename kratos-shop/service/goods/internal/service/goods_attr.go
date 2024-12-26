package service

import (
	"context"
	pb "goods/api/goods/v1"
	"goods/internal/domain"
)

// CreateAttrGroup 创建商品参数属性分组
func (gs *GoodsService) CreateAttrGroup(ctx context.Context, r *pb.CreateAttrGroupReq) (*pb.CreateAttrGroupResp, error) {
	res, err := gs.attr.CreateAttrGroup(ctx, &domain.AttrGroup{
		TypeID: r.TypeId,
		Title:  r.Title,
		Sort:   r.Sort,
		Status: r.Status,
		Desc:   r.Desc,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateAttrGroupResp{
		Id:     res.ID,
		TypeId: res.TypeID,
		Title:  res.Title,
		Sort:   res.Sort,
		Status: res.Status,
		Desc:   res.Desc,
	}, nil
}

func (gs *GoodsService) CreateAttrValue(ctx context.Context, r *pb.CreateAttrValueReq) (*pb.CreateAttrValueResp, error) {
	var values []*domain.GoodsAttrValue
	for _, v := range r.AttrValue {
		res := &domain.GoodsAttrValue{
			GroupID: r.GroupId,
			Value:   v.Value,
		}
		values = append(values, res)
	}

	info, err := gs.attr.CreateAttrValue(ctx, &domain.GoodsAttr{
		TypeID:         r.TypeId,
		GroupID:        r.GroupId,
		Title:          r.Title,
		Sort:           r.Sort,
		Status:         r.Status,
		Desc:           r.Desc,
		GoodsAttrValue: values,
	})
	if err != nil {
		return nil, err
	}

	var attrValue []*pb.CreateAttrValueResp_AttrValue
	for _, v := range info.GoodsAttrValue {
		result := &pb.CreateAttrValueResp_AttrValue{
			Id:      v.ID,
			AttrId:  v.AttrID,
			GroupId: v.GroupID,
			Value:   v.Value,
		}
		attrValue = append(attrValue, result)
	}
	return &pb.CreateAttrValueResp{
		Id:        info.ID,
		TypeId:    info.TypeID,
		GroupId:   info.GroupID,
		Title:     info.Title,
		Desc:      info.Desc,
		Sort:      info.Sort,
		Status:    info.Status,
		AttrValue: attrValue,
	}, nil
}

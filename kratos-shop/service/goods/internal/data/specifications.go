package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// SpecificationsAttr 规格参数信息表
type SpecificationsAttr struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	TypeID    int64          `gorm:"index:type_id;type:int;comment:商品类型ID;not null"`
	Name      string         `gorm:"type:varchar(250);not null;comment:规格参数名称" json:"name"`
	Sort      int32          `gorm:"comment:规格排序;default:99;not null;type:int" json:"sort"`
	Status    bool           `gorm:"comment:参数状态;default:false" json:"status"`
	IsSKU     bool           `gorm:"comment:是否通用的SKU持有;default:false"  json:"is_sku"`
	IsSelect  bool           `gorm:"comment:是否可查询;default:false" json:"is_select"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type SpecificationsAttrValue struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	AttrID    int64          `gorm:"index:attr_id;type:int;comment:规格ID;not null"`
	Value     string         `gorm:"type:varchar(250);not null;comment:规格参数信息值" json:"value"`
	Sort      int32          `gorm:"comment:规格参数值排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type specificationRepo struct {
	data *Data
	log  *log.Helper
}

func NewSpecificationRepo(data *Data, logger log.Logger) biz.SpecificationRepo {
	return &specificationRepo{
		data: data,
		log:  log.NewHelper(logger),
	}

}

func (s *specificationRepo) CreateSpecification(ctx context.Context, req *domain.Specification) (int64, error) {
	specificationAttr := &SpecificationsAttr{
		TypeID:   req.TypeID,
		Name:     req.Name,
		Sort:     req.Sort,
		Status:   req.Status,
		IsSKU:    req.IsSKU,
		IsSelect: req.IsSelect,
	}
	result := s.data.DB(ctx).Save(specificationAttr)
	return specificationAttr.ID, result.Error
}

func (s *specificationRepo) CreateSpecificationValue(ctx context.Context, attrId int64, req []*domain.SpecificationValue) error {
	var values []*SpecificationsAttrValue

	// 遍历规格值
	for _, v := range req {
		res := &SpecificationsAttrValue{
			AttrID: attrId,
			Value:  v.Value,
			Sort:   v.Sort,
		}
		values = append(values, res)
	}

	// 批量创建规格值
	result := s.data.DB(ctx).Create(&values)
	return result.Error
}

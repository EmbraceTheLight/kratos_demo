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

// GoodsAttrGroup 商品属性分组表
type GoodsAttrGroup struct {
	ID          int64          `gorm:"primarykey;type:int" json:"id"`
	GoodsTypeID int64          `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	Title       string         `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc        string         `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status      bool           `gorm:"comment:状态;default:false;not null"`
	Sort        int32          `gorm:"type:int;comment:商品属性排序字段;not null"`
	CreatedAt   time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// GoodsAttr 商品属性表
type GoodsAttr struct {
	ID          int64          `gorm:"primarykey;type:int" json:"id"`
	GoodsTypeID int64          `gorm:"index:goods_type_id;type:int;comment:商品类型ID;not null"`
	GroupID     int64          `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	Title       string         `gorm:"type:varchar(100);comment:属性名;not null"`
	Desc        string         `gorm:"type:varchar(200);comment:属性描述;default:false;not null"`
	Status      bool           `gorm:"comment:状态;default:false;not null"`
	Sort        int32          `gorm:"type:int;comment:商品属性排序字段;not null"`
	CreatedAt   time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// GoodsAttrValue 商品属性值表
type GoodsAttrValue struct {
	ID        int64          `gorm:"primarykey;type:int" json:"id"`
	AttrId    int64          `gorm:"index:property_name_id;type:int;comment:属性表ID;not null"`
	GroupID   int64          `gorm:"index:attr_group_id;type:int;comment:商品属性分组ID;not null"`
	Value     string         `gorm:"type:varchar(100);comment:属性值;not null"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type goodsAttrRepo struct {
	data *Data
	log  *log.Helper
}

func NewGoodsAttrRepo(data *Data, logger log.Logger) biz.GoodsAttrRepo {
	return &goodsAttrRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (g *goodsAttrRepo) CreateGoodsGroupAttr(ctx context.Context, group *domain.AttrGroup) (*domain.AttrGroup, error) {
	attrGroup := GoodsAttrGroup{
		GoodsTypeID: group.TypeID,
		Title:       group.Title,
		Desc:        group.Desc,
		Status:      group.Status,
		Sort:        group.Sort,
	}
	result := g.data.db.Save(&attrGroup)
	if result.Error != nil {
		return nil, result.Error
	}
	return attrGroup.ToDomain(), nil
}

func (g *goodsAttrRepo) IsExistsGroupByID(ctx context.Context, groupID int64) (*domain.AttrGroup, error) {
	var group GoodsAttrGroup
	if res := g.data.db.First(&group, groupID); res.RowsAffected == 0 {
		return nil, errors.New("商品属性组不存在")
	}
	return group.ToDomain(), nil
}

func (g *goodsAttrRepo) CreateGoodsAttr(ctx context.Context, attr *domain.GoodsAttr) (*domain.GoodsAttr, error) {
	goodsAttr := GoodsAttr{
		GoodsTypeID: attr.TypeID,
		GroupID:     attr.GroupID,
		Title:       attr.Title,
		Desc:        attr.Desc,
		Status:      attr.Status,
		Sort:        attr.Sort,
	}

	//data的DB方法会根据传入的ctx判断是否是事务db，因此会自动执行biz层中的事务
	if err := g.data.DB(ctx).Save(&goodsAttr).Error; err != nil {
		return nil, err
	}
	return goodsAttr.ToDomain(), nil
}

func (g *goodsAttrRepo) CreateGoodsAttrValue(ctx context.Context, values []*domain.GoodsAttrValue) ([]*domain.GoodsAttrValue, error) {
	var attrValues []*GoodsAttrValue
	for _, value := range values {
		attrValue := &GoodsAttrValue{
			AttrId:  value.AttrID,
			GroupID: value.GroupID,
			Value:   value.Value,
		}
		attrValues = append(attrValues, attrValue)
	}
	if err := g.data.DB(ctx).Create(&attrValues).Error; err != nil {
		return nil, err
	}

	var res []*domain.GoodsAttrValue
	for _, v := range attrValues {
		value := v.ToDomain()
		res = append(res, value)
	}
	return res, nil
}

func (grp *GoodsAttrGroup) ToDomain() *domain.AttrGroup {
	return &domain.AttrGroup{
		ID:     grp.ID,
		TypeID: grp.GoodsTypeID,
		Title:  grp.Title,
		Desc:   grp.Desc,
		Status: grp.Status,
		Sort:   grp.Sort,
	}
}

func (attr *GoodsAttr) ToDomain() *domain.GoodsAttr {
	return &domain.GoodsAttr{
		ID:      attr.ID,
		TypeID:  attr.GoodsTypeID,
		GroupID: attr.GroupID,
		Title:   attr.Title,
		Sort:    attr.Sort,
		Status:  attr.Status,
		Desc:    attr.Desc,
	}
}

func (atv *GoodsAttrValue) ToDomain() *domain.GoodsAttrValue {
	return &domain.GoodsAttrValue{
		ID:      atv.ID,
		AttrID:  atv.AttrId,
		GroupID: atv.GroupID,
		Value:   atv.Value,
	}
}

package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"gorm.io/gorm"
	"time"
)

// Brand 商品品牌表
type Brand struct {
	ID        int32          `gorm:"primarykey;type:int" json:"id"`
	Name      string         `gorm:"type:varchar(50);not null;comment:品牌名称" json:"name"`
	Logo      string         `gorm:"type:varchar(200);default:;comment:品牌Logo图片"`
	Desc      string         `gorm:"type:varchar(500);default:;comment:品牌描述"`
	IsTab     bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort      int32          `gorm:"comment:品牌排序;default:99;not null;type:int" json:"sort"`
	CreatedAt time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
type brandRepo struct {
	data *Data
	log  *log.Helper
}

func NewBrandRepo(data *Data, logger log.Logger) biz.BrandRepo {
	return &brandRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *brandRepo) IsBrandByID(ctx context.Context, id int32) (*domain.Brand, error) {
	var b Brand
	if err := r.data.db.Model(&Brand{}).Where(&Brand{ID: id}).First(&b).Error; err != nil {
		return nil, err
	}
	return b.ToDomain(), nil

}

func (p *Brand) ToDomain() *domain.Brand {
	return &domain.Brand{
		ID:    p.ID,
		Name:  p.Name,
		Logo:  p.Logo,
		Desc:  p.Desc,
		IsTab: p.IsTab,
		Sort:  p.Sort,
	}
}

package data

import (
	"context"
	stderr "errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"gorm.io/gorm"
	"time"
)

// Category 商品分类表
type Category struct {
	ID               int32          `gorm:"primarykey;type:int" json:"id"`
	Name             string         `gorm:"type:varchar(50);not null;unique;comment:分类名称" json:"name"`
	ParentCategoryID int32          `json:"parent_id"`
	ParentCategory   *Category      `json:"-"`
	SubCategory      []*Category    `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32          `gorm:"column:level;default:1;not null;type:int;comment:分类的级别" json:"level"`
	IsTab            bool           `gorm:"comment:是否显示;default:false" json:"is_tab"`
	Sort             int32          `gorm:"comment:分类排序;default:99;not null;type:int" json:"sort"`
	CreatedAt        time.Time      `gorm:"column:add_time" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:update_time" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
}
type categoryRepo struct {
	data *Data
	log  *log.Helper
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *categoryRepo) AddCategory(ctx context.Context, req *biz.CategoryInfo) (*biz.CategoryInfo, error) {
	//cMap := map[string]interface{}{}
	//cMap["name"] = req.Name
	//cMap["level"] = req.Level
	//cMap["is_tab"] = req.IsTab
	//cMap["sort"] = req.Sort
	//cMap["add_time"] = time.Now()
	//cMap["update_time"] = time.Now()
	var newCategory Category
	newCategory.Name = req.Name
	newCategory.Level = req.Level
	newCategory.IsTab = req.IsTab
	newCategory.Sort = req.Sort

	//查询父级目录是否存在
	if req.Level != 1 {
		var categories Category
		if res := r.data.db.Debug().First(&categories, req.ParentCategory); res.RowsAffected == 0 {
			return nil, stderr.New("商品不存在")
		}
		newCategory.ParentCategoryID = req.ParentCategory
	}

	result := r.data.db.Model(&Category{}).Create(&newCategory)
	if result.Error != nil {
		return nil, result.Error
	}
	//var value int32
	//value, ok := cMap["parent_category_id"].(int32)
	//if !ok {
	//	value = 0
	//}

	//res := &biz.CategoryInfo{
	//	Name:           cMap["name"].(string),
	//	ParentCategory: value,
	//	Level:          cMap["level"].(int32),
	//	IsTab:          cMap["is_tab"].(bool),
	//	Sort:           cMap["sort"].(int32),
	//}

	res := &biz.CategoryInfo{
		Name:           newCategory.Name,
		ParentCategory: newCategory.ParentCategoryID,
		Level:          newCategory.Level,
		IsTab:          newCategory.IsTab,
		Sort:           newCategory.Sort,
	}
	return res, nil
}

func (r *categoryRepo) GetCategoryByID(ctx context.Context, id int32) (*biz.CategoryInfo, error) {
	var categories Category
	if res := r.data.db.First(&categories, id); res.Error != nil {
		return nil, res.Error
	}
	info := &biz.CategoryInfo{
		ID:             categories.ID,
		Name:           categories.Name,
		ParentCategory: categories.ParentCategoryID,
		Level:          categories.Level,
		IsTab:          categories.IsTab,
		Sort:           categories.Sort,
	}
	return info, nil
}

func (r *categoryRepo) GetCategoryAll(ctx context.Context, level, id int32) ([]interface{}, error) {
	categoryIds := make([]interface{}, 0)
	var subQuery string
	// 把一级级分类下的所有三级分类都拿到
	if level == 1 {
		subQuery = fmt.Sprintf("SELECT id FROM categories WHERE parent_category_id IN (SELECT id FROM categories WHERE parent_category_id=%d)", id)
	} else if level == 2 {
		subQuery = fmt.Sprintf("SELECT id FROM categories WHERE parent_category_id=%d", id)
	} else if level == 3 {
		subQuery = fmt.Sprintf("SELECT id FROM categories WHERE id=%d", id)
	}

	type Result struct {
		ID int32
	}

	var results []Result
	if err := r.data.db.Table("categories").Model(Category{}).Raw(subQuery).Scan(&results).Error; err != nil {
		return nil, errors.InternalServer("CATEGORY_ERROR", err.Error())
	}
	for _, re := range results {
		categoryIds = append(categoryIds, re.ID)
	}
	return categoryIds, nil
}

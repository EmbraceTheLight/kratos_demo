package biz

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/domain"
)

type EsGoodsRepo interface {
	// InsertEsGoods 存储商品信息到es
	InsertEsGoods(context.Context, *domain.ESGoods) error

	// GoodsList 商品列表
	GoodsList(context.Context, *domain.EsSearch) ([]int64, int64, error)
}

type EsGoodsUsecase struct {
	repo         EsGoodsRepo
	grepo        GoodsRepo
	categoryRepo CategoryRepo
	log          *log.Helper
}

func NewEsGoodsUsecase(
	repo EsGoodsRepo,
	grepo GoodsRepo,
	categoryRepo CategoryRepo,
	logger log.Logger) *EsGoodsUsecase {
	return &EsGoodsUsecase{
		repo:         repo,
		grepo:        grepo,
		categoryRepo: categoryRepo,
		log:          log.NewHelper(logger),
	}
}

func (uc *EsGoodsUsecase) GoodsList(ctx context.Context, r *domain.ESGoodsFilter) (*domain.GoodsListResp, error) {
	//组织es查询条件
	var es domain.EsSearch
	es.Query = types.NewQuery()
	es.Query.Bool = types.NewBoolQuery()
	if r.Keywords != "" {
		query := types.NewQuery()
		query.MultiMatch = types.NewMultiMatchQuery()
		query.MultiMatch.Query = r.Keywords
		query.MultiMatch.Fields = []string{"name", "goods_brief", "sku.sku_name"}
		es.Query.Bool.Must = append(es.Query.Bool.Must, *query)
	}
	if r.IsHot {
		query := types.NewQuery()
		query.Term["is_hot"] = types.TermQuery{Value: r.IsHot}
		es.Query.Bool.Must = append(es.Query.Bool.Must, *query)
	}
	if r.ClickNum > 0 {
		so := types.NewSortOptions()
		so.SortOptions["click_num"] = types.FieldSort{Order: &sortorder.Desc}
		es.Sorters = append(es.Sorters, so)
	}
	rq := types.NewNumberRangeQuery()
	nrquery := types.NewQuery()
	if r.MinPrice > 0 {
		tmp := types.Float64(r.MinPrice)
		rq.Gte = &tmp
		nrquery.Range["shop_price"] = rq
	}
	if r.MaxPrice > 0 {
		tmp := types.Float64(r.MinPrice)
		rq.Lte = &tmp
		nrquery.Range["shop_price"] = rq
	}
	es.Query.Bool.Should = append(es.Query.Bool.Should, *nrquery)

	if r.BrandsID > 0 {
		query := types.NewQuery()
		query.Term["brands_id"] = types.TermQuery{Value: r.BrandsID}
		es.Query.Bool.Should = append(es.Query.Bool.Should, *query)
	}
	if r.CategoryID > 0 {
		query := types.NewQuery()
		query.Terms = types.NewTermsQuery()

		//查询分类是否存在
		category, err := uc.categoryRepo.GetCategoryByID(ctx, r.CategoryID)
		if err != nil {
			return nil, err
		}
		//查询分类下的所有子分类ID
		categoryIds, err := uc.categoryRepo.GetCategoryAll(ctx, category.Level, r.CategoryID)
		if err != nil {
			return nil, err
		}

		//将分类ID加入查询条件
		var fvs []types.FieldValue
		for _, v := range categoryIds {
			fvs = append(fvs, types.FieldValue(v))
		}
		query.Terms.TermsQuery["category_id"] = categoryIds
		es.Query.Bool.Should = append(es.Query.Bool.Should, *query)
	}

	//分页相关处理
	switch {
	case r.PagePerNums > 100:
		r.PagePerNums = 100
	case r.PagePerNums < 1:
		r.PagePerNums = 10
	}
	if r.Pages == 0 {
		r.Pages = 1
	}
	es.From = (r.Pages - 1) * r.PagePerNums
	es.Size = r.PagePerNums

	// 去 es repo 中查询获得的商品ID
	res := &domain.GoodsListResp{}
	goodsIds, total, err := uc.repo.GoodsList(ctx, &es)
	if err != nil {
		return nil, err
	}
	res.Total = total

	// 根据es返回的商品ID 获取详细的商品信息
	goodsList, err := uc.grepo.GoodsListByIDs(ctx, goodsIds...)
	if err != nil {
		return nil, err
	}
	res.List = goodsList

	return res, nil
}

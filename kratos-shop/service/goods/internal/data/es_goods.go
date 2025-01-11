package data

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/domain"
	"strconv"
)

type esGoodsRepo struct {
	data *Data
	log  *log.Helper
}

func NewEsGoodsRepo(data *Data, logger log.Logger) biz.EsGoodsRepo {
	return &esGoodsRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// GetIndexName 设计商品索引goods
func (r *esGoodsRepo) GetIndexName() string {
	return "goods"
}

// GetMapping 设计商品索引映射
func (r *esGoodsRepo) GetMapping() string {
	goodsMapping := `
    {
    "mappings": {
        "properties": {
            "id": {
                "type": "integer"
            },
            "brands_id": {
                "type": "integer"
            },
            "category_id": {
                "type": "integer"
            },
            "type_id": {
                "type": "integer"
            },
            "click_num": {
                "type": "integer"
            },
            "fav_num": {
                "type": "integer"
            },
            "is_hot": {
                "type": "boolean"
            },
            "is_new": {
                "type": "boolean"
            },
            "market_price": {
                "type": "integer"
            },
            "name": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "brand_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false
            },
            "category_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false
            },
            "type_name": {
                "type": "keyword",
                "index": false,
                "dec_values": false
            },
            "goods_brief": {
                "type": "text",
                "analyzer": "ik_max_word"
            },
            "on_sale": {
                "type": "boolean"
            },
            "ship_free": {
                "type": "boolean"
            },
            "shop_price": {
                "type": "integer"
            },
            "sold_num": {
                "type": "integer"
            },
            "sku": {
                "type": "nested",
                "sku_id": {
                    "type": "integer"
                },
                "sku_name": {
                    "type": "text",
                    "analyzer": "ik_max_word"
                },
                "sku_price": {
                    "type": "integer"
                }
            }
        }
    }
}`
	return goodsMapping
}

func (r *esGoodsRepo) InsertEsGoods(ctx context.Context, esModel *domain.ESGoods) error {
	//新建mapping和index
	exists, err := r.data.esClient.Indices.Exists(r.GetIndexName()).Do(ctx)
	if err != nil {
		panic(err)
	}
	//索引不存在，创建索引
	if !exists {
		//构建请求，这里是为index创建mapping
		req, err := create.NewRequest().FromJSON(r.GetMapping())
		if err != nil {
			return err
		}
		_, err = r.data.esClient.Indices.Create(r.GetIndexName()).Request(req).Do(ctx)
		if err != nil {
			return err
		}
	}

	//向es中插入数据
	_, err = r.data.esClient.
		Index(r.GetIndexName()).
		Request(esModel).
		Id(strconv.Itoa(int(esModel.ID))).
		Do(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *esGoodsRepo) GoodsList(ctx context.Context, filter *domain.EsSearch) ([]int64, int64, error) {
	r.printEsRawReq(ctx, filter)
	result, err := r.data.esClient.Search().
		Index(r.GetIndexName()).
		Query(filter.Query).
		Sort(filter.Sorters...).
		From(int(filter.From)).
		Size(int(filter.Size)).
		Do(ctx)

	if err != nil {
		return nil, 0, err
	}
	//取出商品ID
	goodsIds := make([]int64, 0)
	for _, value := range result.Hits.Hits {
		goods := &domain.ESGoods{}
		_ = json.Unmarshal(value.Source_, goods)
		goodsIds = append(goodsIds, goods.ID)
	}
	return goodsIds, result.Hits.Total.Value, nil
}

// printEsRawReq 打印es原始请求
func (r *esGoodsRepo) printEsRawReq(ctx context.Context, filter *domain.EsSearch) {
	fmt.Printf("\n===========================  ES Raw Request ==============================\n")
	req, _ := r.data.esClient.Search().
		Index(r.GetIndexName()).
		Query(filter.Query).
		Sort(filter.Sorters...).
		From(int(filter.From)).
		Size(int(filter.Size)).
		HttpRequest(ctx)
	//打印请求方法
	fmt.Printf("%s %s\n", req.Method, req.URL.String())

	//打印请求头
	for k := range req.Header {
		fmt.Printf("%s: %s\n", k, req.Header.Get(k))
	}
	fmt.Println()

	//打印请求体
	var str bytes.Buffer
	b := make([]byte, req.ContentLength)
	_, _ = req.Body.Read(b)
	json.Indent(&str, b, "", "  ")
	fmt.Println(str.String())
	fmt.Printf("\n===========================  ES Raw Request End =============================\n")
}

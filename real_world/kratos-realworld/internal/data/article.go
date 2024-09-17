package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

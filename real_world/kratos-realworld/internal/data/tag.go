package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-realworld/internal/biz"
)

type tagRepo struct {
	data *Data
	log  *log.Helper
}

// NewTagRepo .
func NewTagRepo(data *Data, logger log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

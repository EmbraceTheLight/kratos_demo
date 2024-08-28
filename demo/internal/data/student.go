package data

import (
	"context"
	"demo/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data //用于连接数据库客户端
	log  *log.Helper
}

func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) Save(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	return s, nil
}

func (r *studentRepo) Get(ctx context.Context, s *biz.Student) (*biz.Student, error) {
	return s, nil
}

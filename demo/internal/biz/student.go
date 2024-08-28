package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Student struct {
	ID      string
	Name    string
	Sayname string
}

type StudentRepo interface {
	Save(context.Context, *Student) (*Student, error)
	Get(context.Context, *Student) (*Student, error)
}

// StudentUsercasse 对 Student 的操作加上日志
type StudentUsercasse struct {
	repo StudentRepo
	log  *log.Helper
}

// NewStudentUsercase 初始化 StudentUsercase
func NewStudentUsercase(repo StudentRepo, logger log.Logger) *StudentUsercasse {
	return &StudentUsercasse{
		repo: repo,
		log:  log.NewHelper(logger),
	}
}

/*编写业务逻辑*/
func (uc *StudentUsercasse) CreateStudent(ctx context.Context, stu *Student) (*Student, error) {
	uc.log.WithContext(ctx).Infof("CreateStudent: %v", stu.ID)
	return uc.repo.Save(ctx, stu)
}

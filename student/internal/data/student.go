package data

import (
	"context"
	"student/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type studentRepo struct {
	data *Data
	log  *log.Helper
}

// NewStudentRepo .
func NewStudentRepo(data *Data, logger log.Logger) biz.StudentRepo {
	return &studentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *studentRepo) GetStudent(ctx context.Context, id int32) (*biz.Student, error) {
	var stu biz.Student
	err := r.data.gormDB.Where("id = ?", id).First(&stu).Error
	if err != nil {
		return nil, err
	}

	r.log.WithContext(ctx).Info("gormDB: GetStudent, id: ", id)
	return &biz.Student{
		ID:        stu.ID,
		Name:      stu.Name,
		Info:      stu.Info,
		Status:    stu.Status,
		UpdatedAt: stu.UpdatedAt,
		CreatedAt: stu.CreatedAt,
	}, nil
}
func (r *studentRepo) CreateStudent(ctx context.Context, stu *biz.Student) error {
	err := r.data.gormDB.Create(stu).Error
	if err != nil {
		return err
	}
	r.log.WithContext(ctx).Info("gormDB: CreateStudent")
	return nil
}

func (r *studentRepo) UpdateStudent(ctx context.Context, stu *biz.Student) (*biz.Student, error) {
	err := r.data.gormDB.Where("id = ?", stu.ID).Updates(stu).Error
	if err != nil {
		return nil, err
	}
	r.log.WithContext(ctx).Info("gormDB: UpdateStudent: id: ", stu.ID)
	return stu, nil
}

func (r *studentRepo) DeleteStudent(ctx context.Context, id int32) error {
	err := r.data.gormDB.Where("id = ?", id).Delete(&biz.Student{}).Error
	if err != nil {
		return err
	}
	r.log.WithContext(ctx).Info("gormDB: DeleteStudent: id: ", id)
	return nil
}

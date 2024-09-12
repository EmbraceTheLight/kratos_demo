package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Student is a Greeter model.
type Student struct {
	ID        int32
	Name      string
	Info      string
	Status    int32
	UpdatedAt time.Time
	CreatedAt time.Time
}

// StudentRepo is a Greater repo.
type StudentRepo interface {
	GetStudent(context.Context, int32) (*Student, error)
}

// StudentUsecase is a Student usecase.
type StudentUsecase struct {
	repo StudentRepo
	log  *log.Helper
}

// NewStudentUsecase new a Student usecase.
func NewStudentUsecase(repo StudentRepo, logger log.Logger) *StudentUsecase {
	return &StudentUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *StudentUsecase) Get(ctx context.Context, id int32) (*Student, error) {
	uc.log.WithContext(ctx).Infof("biz.Get: %d", id)
	return uc.repo.GetStudent(ctx, id)
}

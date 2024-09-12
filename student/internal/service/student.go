package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "student/api/student/v1"
	"student/internal/biz"
)

type StudentService struct {
	pb.UnimplementedStudentServer

	student *biz.StudentUsecase
	log     *log.Helper
}

func NewStudentService(stu *biz.StudentUsecase, logger log.Logger) *StudentService {
	return &StudentService{
		student: stu,
		log:     log.NewHelper(logger)}
}

func (s *StudentService) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentReply, error) {
	err := s.student.Create(ctx, &biz.Student{
		Name:   req.Name,
		Info:   req.Info,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateStudentReply{
		Code:    2000,
		Message: "创建学生成功",
	}, nil
}
func (s *StudentService) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentReply, error) {
	_, err := s.student.Update(ctx, &biz.Student{
		ID:     req.Id,
		Name:   req.Name,
		Info:   req.Info,
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateStudentReply{}, nil
}
func (s *StudentService) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentReply, error) {
	err := s.student.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteStudentReply{}, nil
}
func (s *StudentService) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentReply, error) {
	stu, err := s.student.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetStudentReply{
		Id:     stu.ID,
		Status: stu.Status,
		Name:   stu.Name,
	}, nil
}

func (s *StudentService) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentReply, error) {
	return &pb.ListStudentReply{}, nil
}

package service

import (
	"context"
	pb "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.SingleCommentReply, error) {
	return &pb.SingleCommentReply{}, nil
}
func (s *RealWorldService) GetComments(ctx context.Context, req *pb.GetCommentsRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}
func (s *RealWorldService) DeleteComments(ctx context.Context, req *pb.DeleteCommentsRequest) (*pb.MultipleCommentsReply, error) {
	return &pb.MultipleCommentsReply{}, nil
}

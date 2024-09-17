package service

import (
	"context"
	pb "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) GetTags(ctx context.Context, req *pb.EmptyRequest) (*pb.TasLIstReply, error) {
	return &pb.TasLIstReply{}, nil
}

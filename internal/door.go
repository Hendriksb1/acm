package internal

import (
	"context"
	pb "acm/api/pb"
)

type Door struct {
	Id string
	AccessLevel int32
}

func (s *Server) AddDoor (ctx context.Context, req *pb.AddDoorRequest) (*pb.AddDoorResponse, error) {
	// TODO continue here
	return nil, nil
}

func (s *Server) RemoveDoor (ctx context.Context, req *pb.RemoveDoorRequest) (*pb.RemoveDoorResponse, error) {
	// TODO continue here
	return nil, nil
}
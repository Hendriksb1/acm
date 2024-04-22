package internal

import (
	"acm/api/pb"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// User represents a user with a chip card
type User struct {
    Name         string
    ChipCardID   string
    AccessRights int32
}

// AddUser adds a new user to the access control system
func(s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error){

    if req.Name == "" {
        return nil, status.Error(codes.InvalidArgument, "no name")
    }

    if req.ChipCardId == "" {
        return nil, status.Error(codes.InvalidArgument, "no chip card id")
    }

    user := &User{
		Name: req.Name,
		ChipCardID: req.ChipCardId,
		AccessRights: req.AccessRights,
	}

	err := s.DB.AddUser(user)
	if err != nil {
		return nil, err
	}

	return &pb.AddUserResponse{}, nil
}

// CheckAccess checks if a user with the given chip card ID has access rights to a door
func(s *Server) CheckAccess(ctx context.Context, req *pb.CheckAccessRequest) (*pb.CheckAccessResponse, error){

    if req.ChipCardId == "" {
        return nil, status.Error(codes.InvalidArgument, "no id")
    }

    if req.DoorLevel == 0 {
        return nil, status.Error(codes.InvalidArgument, "no door level")
    }

	user, err := s.DB.GetUserByChipCardId(req.ChipCardId)
	if err != nil {
        fmt.Println("access denied: unknown chip card ID")
		return &pb.CheckAccessResponse{
			HasAccess: false,
		}, nil
	}

    if user.AccessRights >= req.DoorLevel {
        fmt.Printf("access granted: welcome, %s!\n", user.Name)
        return &pb.CheckAccessResponse{
			HasAccess: true,
		}, nil
    }

    fmt.Println("access denied: insufficient access rights")
    return  &pb.CheckAccessResponse{
		HasAccess: false,
	}, nil
}

// DeleteUserByChipCardId deletes a user from the database based on the chip card ID.
func(s *Server) DeleteUserByChipCardId(ctx context.Context, req *pb.DeleteUserByChipCardIdRequest) (*pb.DeleteUserByChipCardIdResponse, error){
    if req.ChipCardId == "" {
        return nil, status.Error(codes.InvalidArgument, "no id")
    }

    err := s.DB.DeleteUserByChipCardId(req.ChipCardId)
    if err != nil {
		return nil, err
	}

    return &pb.DeleteUserByChipCardIdResponse{}, nil
}
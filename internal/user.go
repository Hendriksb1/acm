package internal

import (
	"acm/api/pb"
	"context"
	"fmt"
)

// User represents a user with a chip card
type User struct {
    Name         string
    ChipCardID   string
    AccessRights int32
}

// AddUser adds a new user to the access control system
func(s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error){
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

	user, err := s.DB.GetUserByChipCardId(req.ChipCardId)
	if err != nil {
        fmt.Println("access denied: Unknown chip card ID")
		return &pb.CheckAccessResponse{
			HasAccess: false,
		}, nil
	}

    if user.AccessRights >= req.DoorLevel {
        fmt.Printf("access granted: Welcome, %s!\n", user.Name)
        return &pb.CheckAccessResponse{
			HasAccess: true,
		}, nil
    }

    fmt.Println("access denied: Insufficient access rights")
    return  &pb.CheckAccessResponse{
		HasAccess: false,
	}, nil
}
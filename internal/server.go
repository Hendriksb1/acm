package internal

import (
	pb "acm/api/pb"
	"context"
	"fmt"
)

type Server struct{
    pb.UnimplementedAccessControlManagerServer
	AccessControlSystem *AccessControlSystem
}

func ServerInit() *Server {
	return &Server{
		AccessControlSystem: NewAccessControlSystem(),
	}
}

// AddUser adds a new user to the access control system
func(s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error){
    user := User{
		Name: req.Name,
		ChipCardID: req.ChipCardId,
		AccessRights: req.AccessRights,
	}

    s.AccessControlSystem.Users[req.ChipCardId] = user

	return &pb.AddUserResponse{}, nil
}

// CheckAccess checks if a user with the given chip card ID has access rights to a door
func(s *Server) CheckAccess(ctx context.Context, req *pb.CheckAccessRequest) (*pb.CheckAccessResponse, error){

    user, ok := s.AccessControlSystem.Users[req.ChipCardId]
    if !ok {
        fmt.Println("Access denied: Unknown chip card ID")
        return &pb.CheckAccessResponse{
			HasAccess: false,
		}, nil
    }

    if user.AccessRights >= req.DoorLevel {
        fmt.Printf("Access granted: Welcome, %s!\n", user.Name)
        return &pb.CheckAccessResponse{
			HasAccess: true,
		}, nil
    }

    fmt.Println("Access denied: Insufficient access rights")
    return  &pb.CheckAccessResponse{
		HasAccess: false,
	}, nil
}
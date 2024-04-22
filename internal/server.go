package internal

import (
	pb "acm/api/pb"
)

// Server represents the server for the access control manager.
type Server struct{
    pb.UnimplementedAccessControlManagerServer
	DB Database
}

// ServerInit initializes a new instance of Server with the provided database.
func ServerInit(db Database) *Server {
	return &Server{
		DB: db,
	}
}
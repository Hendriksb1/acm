package internal

import (
	pb "acm/api/pb"
)

type Server struct{
    pb.UnimplementedAccessControlManagerServer
	DB Database
}

func ServerInit(db Database) *Server {
	return &Server{
		DB: db,
	}
}
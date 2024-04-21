package main

import (
	pb "acm/api/pb"
	"acm/internal"
	"log"
	"net"

	"os"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

    postgresDB := &internal.PostgresDB{}

	server := internal.ServerInit(postgresDB)
	logger.Info("server initated")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	logger.Info("listening on port", port)

	grpcServer := grpc.NewServer()
	pb.RegisterAccessControlManagerServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Error("failed to serve: %v", err)
		}
	}()

	logger.Info("grpc service registered")
}

package main

import (
	pb "acm/api/pb"
	"acm/internal"
	"context"
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

	server := internal.ServerInit()
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

	_, err = server.AddUser(context.Background(), &pb.AddUserRequest{
		Name:         "Alice",
		ChipCardId:   "123456789",
		AccessRights: int32(pb.AccessLevel_LEVEL_1.Number()),
	})
	if err != nil {
		slog.Warn("failed to add user")
	}
	logger.Info("user added")

	// should fail
	result, err := server.CheckAccess(context.Background(), &pb.CheckAccessRequest{
		ChipCardId: "123456789",
		DoorLevel:  int32(pb.AccessLevel_ADMIN),
	})
	if err != nil {
		slog.Warn("failed to check access")
	}
	slog.Info("check completetd and result is: ", result.HasAccess)

    // should succeed
    result, err = server.CheckAccess(context.Background(), &pb.CheckAccessRequest{
		ChipCardId: "123456789",
		DoorLevel:  int32(pb.AccessLevel_LEVEL_1),
	})
	if err != nil {
		slog.Warn("failed to check access")
	}

	slog.Info("check completetd and result is: ", result.HasAccess)
}

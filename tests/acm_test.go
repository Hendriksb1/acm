// server_test.go

package internal_test

import (
	"acm/api/pb"
	"acm/internal"
	"context"
	"net"
	"os"
	"testing"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var (
	testServer *grpc.Server
	testClient pb.AccessControlManagerClient
)

func TestMain(m *testing.M) {

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	server := internal.ServerInit()
	log.Info("server initated")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("listening on port", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAccessControlManagerServer(grpcServer, server)

	// Start the gRPC server
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("failed to serve", err)
		}
	}()

	// Create a gRPC client connection to the test server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Error("failed to dial", err)
	}
	defer conn.Close()

	testClient = pb.NewAccessControlManagerClient(conn)

	// Run tests
	exitCode := m.Run()

	// Shutdown the server
	grpcServer.Stop()
	os.Exit(exitCode)
}


func TestAddUserAndCheck(t *testing.T) {

	// add user 
	addUserReq := &pb.AddUserRequest{
		Name:         "Anna",
		ChipCardId:   "123",
		AccessRights: int32(pb.AccessLevel_LEVEL_1),
	}
	_, err := testClient.AddUser(context.Background(), addUserReq)
	if err != nil {
		t.Error("failed to add user")
	}

	testCases := []struct {
		ChipCardID string
		DoorLevel  pb.AccessLevel
		Expected   bool
	}{
		// Positive test cases
		{"123", pb.AccessLevel_LEVEL_1, true},

		// Negative test cases
		{"123", pb.AccessLevel_ADMIN, false},

		// Invalid user test cases
		{"999", pb.AccessLevel_LEVEL_1, false},
	}

	for _, c := range testCases {
		checkRequest := &pb.CheckAccessRequest{
			ChipCardId: c.ChipCardID,
			DoorLevel:  int32(c.DoorLevel),
		}

		res, err := testClient.CheckAccess(context.Background(), checkRequest)
		if err != nil {
			t.Error("failed to check access")
		}
		if res.HasAccess != c.Expected {
			t.Error("expected", c.Expected, "but got", res.HasAccess)
		}				
	}

	// 	time=2024-04-21T12:26:05.889+02:00 level=INFO msg="server initated"
	// Access granted: Welcome, Anna!
	// Access denied: Insufficient access rights
	// Access denied: Unknown chip card ID
	// PASS
	// ok      acm/tests       0.720s
}

// Add more test cases as needed

package main

import (
	pb "acm/api/pb"
	"acm/internal"
	"net"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
)

func main() {

	log := slog.New(slog.NewTextHandler(os.Stderr, nil))

	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Error("failed to read env file")
	}

	port := os.Getenv("PORT")
	dbConnString := os.Getenv("DB")

	postgresDB, err := internal.NewPostgresDB(dbConnString, log)
	if err != nil {
		return
	}
	log.Info("db connected", dbConnString)

	err = postgresDB.InnitUserTable(log)
	if err != nil {
		return
	}
	log.Info("users table initated")

	server := internal.ServerInit(postgresDB)
	log.Info("server initated")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("failed to listen", err)
	}
	log.Info("listening on port", port)

	grpcServer := grpc.NewServer()
	pb.RegisterAccessControlManagerServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Error("failed to serve: %v", err)
		}
	}()

	log.Info("grpc service registered")
}

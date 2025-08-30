package main

import (
	runner "Runlet_runners/internal/infrastructure/proto"
	"log/slog"
	"net"
	"os"

	grpcImpl "Runlet_runners/internal/interfaces/grpc"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("cannot start runner server", "error", err)
		os.Exit(1)
	}
	server := grpc.NewServer()
	runner.RegisterRunnerServer(server, &grpcImpl.Server{})
	if err := server.Serve(lis); err != nil {
		slog.Error("cannot run grpc server", "error", err)
		os.Exit(1)
	}
	slog.Info("started runner container", "lang", os.Getenv("LANG"))
}

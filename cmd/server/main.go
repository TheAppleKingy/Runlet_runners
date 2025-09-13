package main

import (
	"Runlet_runners/internal/application/service"
	"Runlet_runners/internal/infrastructure/config"
	"Runlet_runners/internal/infrastructure/implementations"
	grpc_interfaces "Runlet_runners/internal/infrastructure/proto"
	"log/slog"
	"net"
	"os"

	grpcImpl "Runlet_runners/internal/interfaces/grpc"

	"google.golang.org/grpc"
)

func main() {
	config.LoadConfigs()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		slog.Error("cannot start runner server", "error", err)
		os.Exit(1)
	}

	runner := implementations.NewCodeRunner()
	runService := service.NewRunCodeService(runner)

	server := grpc.NewServer()
	grpc_interfaces.RegisterRunnerServer(server, &grpcImpl.Server{RunService: runService})
	slog.Info("started runner container", "lang", os.Getenv("LANG"))
	if err := server.Serve(lis); err != nil {
		slog.Error("cannot run grpc server", "error", err)
		os.Exit(1)
	}
}

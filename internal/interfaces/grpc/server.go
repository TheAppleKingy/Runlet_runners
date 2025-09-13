package grpc

import (
	"Runlet_runners/internal/application/dto"
	"Runlet_runners/internal/application/service"
	runner "Runlet_runners/internal/infrastructure/proto"
	"context"
	"encoding/json"
	"log/slog"
)

type Server struct {
	runner.UnimplementedRunnerServer
	RunService *service.RunCodeService
}

func (s *Server) RunCode(ctx context.Context, request *runner.RunCodeRequest) (*runner.RunCodeResponse, error) {
	var testsData []dto.RunData
	var response runner.RunCodeResponse
	if err := json.Unmarshal(request.Cases, &testsData); err != nil {
		slog.Error("cannot decode cases to run tests", "error", err, "problem_id", request.Problem, "student_id", request.Student, "runner", request.Lang)
		return &response, err
	}
	results := s.RunService.TestSolution(testsData, int(request.Student), int(request.Problem), request.Lang, request.Code)
	r, _ := json.Marshal(results)
	response.Results = r
	return &response, nil
}

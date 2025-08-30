package grpc

import (
	"Runlet_runners/internal/application/dto"
	"Runlet_runners/internal/domain/entities"
	"Runlet_runners/internal/infrastructure"
	runner "Runlet_runners/internal/infrastructure/proto"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

type Server struct {
	runner.UnimplementedRunnerServer
}

func (s *Server) RunCode(ctx context.Context, request *runner.RunCodeRequest) (*runner.RunCodeResponse, error) {
	var testsData []dto.RunTestData
	var response runner.RunCodeResponse
	if err := json.Unmarshal(request.Cases, &testsData); err != nil {
		slog.Error("cannot decode cases to run tests", "error", err, "problem_id", request.Problem, "student_id", request.Student, "runner", request.Lang)
		return &response, err
	}
	var results entities.TestCases
	for _, testCase := range testsData {
		tmpFile, err := os.CreateTemp("/tmp", fmt.Sprintf("runlet-*-s-%d-p-%d.%s", request.Student, request.Problem, request.Lang))
		if err != nil {
			slog.Error("cannot create tempfile", "error", err, "problem_id", request.Problem, "student_id", request.Student, "runner", request.Lang)
			return &response, err
		}
		tmpFile.WriteString(request.Code)

		cmd := exec.Command(infrastructure.RunCommands[request.Lang], tmpFile.Name())
		var output bytes.Buffer
		var errOut bytes.Buffer
		cmd.Stdin = bytes.NewBufferString(testCase.Input)
		cmd.Stdout = &output
		cmd.Stderr = &errOut
		cmd.Run()
		if errOut.String() != "" {
			output = *bytes.NewBufferString("internal error")
		}
		res := entities.TestCase{
			TestNum: testCase.TestNum,
			Input:   testCase.Input,
			Output:  strings.TrimSpace(output.String()),
		}
		results = append(results, res)
		os.Remove(tmpFile.Name())
	}
	r, _ := json.Marshal(results)
	response.Results = r
	return &response, nil
}

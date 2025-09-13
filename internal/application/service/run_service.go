package service

import (
	"Runlet_runners/internal/application/dto"
	"Runlet_runners/internal/domain/entities"
	"Runlet_runners/internal/domain/interfaces"
	"Runlet_runners/internal/infrastructure/implementations"
	"os"
)

type RunCodeService struct {
	Runner interfaces.Runner
}

func NewRunCodeService(runner interfaces.Runner) *RunCodeService {
	return &RunCodeService{
		Runner: runner,
	}
}

func (rs RunCodeService) TestSolution(testsData []dto.RunData, s int, p int, lang string, code string) entities.TestCases {
	f, err := implementations.WriteTemp(s, p, lang, code)
	if err != nil {
		return entities.TestCases{{TestNum: 1, Output: "internal error"}}
	}
	f.Close()
	toExec := f.Name()
	defer os.Remove(toExec)

	results, err := rs.Runner.Run(testsData, toExec, lang)
	if err != nil {
		return entities.TestCases{{TestNum: 1, Output: err.Error()}}
	}

	return results
}

package service

import (
	"Runlet_runners/internal/application/dto"
	"Runlet_runners/internal/domain/entities"
	"Runlet_runners/internal/infrastructure/runtime"
	"fmt"
	"log/slog"
	"os"
	"strings"
)

// writeTemp creates and returns pointer to temporary file containing code for running
func writeTemp(s int, p int, lang string, code string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("/tmp", fmt.Sprintf("runlet-*-s-%d-p-%d.%s", s, p, lang))
	if err != nil {
		return nil, err
	}
	_, err = tmpFile.WriteString(code)
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, fmt.Errorf("cannot save code for running: %w", err)
	}
	return tmpFile, nil
}

func TestSolution(testsData []dto.RunTestData, s int, p int, lang string, code string) entities.TestCases {
	var results entities.TestCases
	f, err := writeTemp(s, p, lang, code)
	if err != nil {
		slog.Error("cannot create tempfile", "error", err, "problem_id", p, "student_id", s, "runner", lang)
		return results
	}
	f.Close()

	toExec := f.Name()
	defer os.Remove(toExec)

	c, ok := runtime.CodeMap[lang]
	if !ok {
		slog.Error("language not registered on runner's side", "problem_id", p, "student_id", s, "runner", lang)
	}

	if c.Compiled {
		toExec, err = runtime.Compile(toExec, lang)
		defer os.Remove(toExec)
		if err != nil {
			slog.Error("cannot compile bin", "error", err, "problem_id", p, "student_id", s, "runner", lang)
			return results
		}
	}

	var output string
	for _, testCase := range testsData {
		output = c.Run(toExec, testCase.Input)
		results = append(results, entities.TestCase{
			TestNum: testCase.TestNum,
			Input:   testCase.Input,
			Output:  strings.TrimRight(output, "\n"),
		})
	}
	return results
}

package implementations

import (
	"Runlet_runners/internal/application/dto"
	"Runlet_runners/internal/domain/entities"
	"Runlet_runners/internal/domain/interfaces"
	"Runlet_runners/internal/infrastructure/config"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

// WriteTemp creates and returns pointer to temporary file containing code for running
func WriteTemp(s int, p int, lang string, code string) (*os.File, error) {
	tmpFile, err := os.CreateTemp("/tmp", fmt.Sprintf("runlet-*-s-%d-p-%d.%s", s, p, lang))
	if err != nil {
		return nil, err
	}
	_, err = tmpFile.WriteString(code)
	if err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		slog.Error("cannot write code to temp file", "error", err, "src", tmpFile.Name())
		return nil, fmt.Errorf("cannot save code for running: %w", err)
	}
	return tmpFile, nil
}

// CodeRunner is implementation of interfaces.Runner
type CodeRunner struct{}

func NewCodeRunner() interfaces.Runner {
	return CodeRunner{}
}

// compile compiles binary based on src via lang compilator and returns name of binary file or provided src if language is interpreted
func (cr CodeRunner) compile(src string, lang string) (string, error) {
	binaryName := strings.TrimSuffix(src, "."+lang)
	langConf, ok := config.LaunguagesConfig.Languages[lang]
	if !ok {
		return "", fmt.Errorf("config for language %s was not registered", lang)
	}
	args := langConf.CompileCommandArgs
	if len(args) == 0 {
		return src, nil
	}

	compileArgs := make([]string, len(args))
	for i, arg := range args {
		switch {
		case strings.Contains(arg, "{src}"):
			compileArgs[i] = strings.ReplaceAll(arg, "{src}", src)
		case strings.Contains(arg, "{bin}"):
			compileArgs[i] = strings.ReplaceAll(arg, "{bin}", binaryName)
		default:
			compileArgs[i] = arg
		}
	}

	cmd := exec.Command(compileArgs[0], compileArgs[1:]...)
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cannot compile bin from %s: %w", src, err)
	}
	return binaryName, nil
}

func (cr CodeRunner) Run(runData []dto.RunData, src string, lang string) (entities.TestCases, error) {
	var err error
	var results entities.TestCases

	fName, err := cr.compile(src, lang)
	if err != nil {
		slog.Error("error compiling", "error", err, "src", src)
		return entities.TestCases{}, fmt.Errorf("error compiling code for lang %s", lang)
	}

	args := config.LaunguagesConfig.Languages[lang].RunCommandArgs
	runArgs := make([]string, len(args))
	for i, arg := range args {
		runArgs[i] = strings.ReplaceAll(arg, "{src}", fName)
	}

	for _, testData := range runData {
		cmd := exec.Command(runArgs[0], runArgs[1:]...)
		var output bytes.Buffer
		cmd.Stdout = &output
		cmd.Stdin = bytes.NewBufferString(testData.Input)
		err = cmd.Run()
		out := strings.TrimRight(output.String(), "\n")
		if err != nil {
			out = err.Error()
		}
		results = append(results, entities.TestCase{
			TestNum: testData.RunNum,
			Input:   testData.Input,
			Output:  out,
		})
	}
	os.Remove(src)
	return results, nil
}

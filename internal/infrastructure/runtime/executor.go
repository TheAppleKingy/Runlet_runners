package runtime

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Compile compiles binary based on src via lang compilator and return name of binary file
func Compile(src string, lang string) (string, error) {
	binaryName := strings.TrimSuffix(src, "."+lang)
	compileCommand := compileCommands[lang](binaryName, src)
	if err := compileCommand.Run(); err != nil {
		return "", fmt.Errorf("cannot compile bin from %s: %w", src, err)
	}
	return binaryName, nil
}

type CodeRunner struct {
	lang     string
	Compiled bool
}

// Run runs programs based on binaries for compiled languages and based on file for interpreted languages.
func (cr CodeRunner) Run(src string, input string) string {
	cmd := runCommands[cr.lang](src)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stdin = bytes.NewBufferString(input)
	if err := cmd.Run(); err != nil {
		return err.Error()
	}
	return output.String()
}

var compileCommands = map[string]func(binaryName string, src string) *exec.Cmd{
	"go": func(binaryName string, src string) *exec.Cmd {
		return exec.Command("go", "build", "-o", binaryName, src)
	},
}

// runCommand is a map representing commands for running code of certain programming language
var runCommands = map[string]func(execFile string) *exec.Cmd{
	"py": func(execFile string) *exec.Cmd {
		return exec.Command("python", execFile)
	},
	"go": func(execFile string) *exec.Cmd {
		return exec.Command(execFile)
	},
	"js": func(execFile string) *exec.Cmd {
		return exec.Command("node", execFile)
	},
}

// CodeMap is a map for CodeRunners. To add new runner create new corresponding container and register CodeRunner structure here
var CodeMap = map[string]CodeRunner{
	"py": {
		lang:     "py",
		Compiled: false,
	},
	"go": {
		lang:     "go",
		Compiled: true,
	},
	"js": {
		lang:     "js",
		Compiled: false,
	},
}

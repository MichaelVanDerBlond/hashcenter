package system

import (
	"bytes"
	"os/exec"
	"strings"
)

type CommandResult struct {
	Command  []string
	ExitCode int
	Output   string
}

func RunCommand(command []string) (*CommandResult, error) {

	if len(command) == 0 {
		return &CommandResult{}, nil
	}

	cmd := exec.Command(command[0], command[1:]...)

	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()

	result := &CommandResult{
		Command:  command,
		Output:   strings.TrimSpace(out.String()),
		ExitCode: 0,
	}

	if err == nil {
		return result, nil
	}

	if exitErr, ok := err.(*exec.ExitError); ok {
		result.ExitCode = exitErr.ExitCode()
		return result, err
	}

	result.ExitCode = -1
	return result, err
}

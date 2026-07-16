package analyze

import (
	"bytes"
	"os/exec"
	"strings"
)

func run(name string, args ...string) ([]byte, error) {

	cmd := exec.Command(name, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if stdout.Len() > 0 {
		return stdout.Bytes(), err
	}

	return []byte(strings.TrimSpace(stderr.String())), err
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

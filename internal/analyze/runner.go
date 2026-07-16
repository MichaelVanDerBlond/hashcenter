package analyze

import (
	"os/exec"
)

func run(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).CombinedOutput()
}

func exists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}

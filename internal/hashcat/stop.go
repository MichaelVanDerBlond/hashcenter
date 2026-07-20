package hashcat

import (
	"errors"
	"syscall"
)

func (r *Runner) Stop() error {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cmd == nil {
		return errors.New("hashcat is not running")
	}

	if r.cmd.Process == nil {
		return errors.New("process is not available")
	}

	return r.cmd.Process.Signal(syscall.SIGTERM)
}

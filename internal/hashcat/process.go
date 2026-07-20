package hashcat

import "os"

func (r *Runner) PID() int {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cmd == nil || r.cmd.Process == nil {
		return 0
	}

	return r.cmd.Process.Pid
}

func (r *Runner) Running() bool {

	return r.PID() != 0
}

func (r *Runner) Process() *os.Process {

	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cmd == nil {
		return nil
	}

	return r.cmd.Process
}

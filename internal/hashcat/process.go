package hashcat

import "os"

func (r *Runner) PID() int {

	process := r.Process()
	if process == nil {
		return 0
	}

	return process.Pid
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

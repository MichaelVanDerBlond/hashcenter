package hashcat

import "os"

func (r *Runner) PID() int {
	if process := r.Process(); process != nil {
		return process.Pid
	}

	return 0
}

func (r *Runner) Running() bool {
	return r.Process() != nil
}

func (r *Runner) Process() *os.Process {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cmd == nil {
		return nil
	}

	return r.cmd.Process
}

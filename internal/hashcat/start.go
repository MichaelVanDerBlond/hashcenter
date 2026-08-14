package hashcat

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"sync"
	"time"
)

func (r *Runner) Start(ctx context.Context, args ...string) (*Result, *Status, error) {
	r.mu.Lock()
	r.done = make(chan struct{})
	r.mu.Unlock()

	parser := NewParser()

	cmd := r.Command(ctx, args...)

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return nil, nil, err
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return nil, nil, err
	}

	if err := cmd.Start(); err != nil {
		return nil, nil, err
	}

	r.mu.Lock()
	r.cmd = cmd
	r.mu.Unlock()

	result := &Result{
		Started: time.Now(),
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	var wg sync.WaitGroup

	copyPipe := func(dst *bytes.Buffer, src io.Reader, parse bool) {
		defer wg.Done()

		scanner := bufio.NewScanner(src)

		for scanner.Scan() {
			line := scanner.Text()

			dst.WriteString(line)
			dst.WriteByte('\n')

			if parse {
				parser.Parse(line)
			}
		}
	}

	wg.Add(2)

	go copyPipe(&stdout, stdoutPipe, true)
	go copyPipe(&stderr, stderrPipe, false)

	go func() {
		result.Error = cmd.Wait()

		wg.Wait()

		result.Finished = time.Now()
		result.Duration = result.Finished.Sub(result.Started)
		result.Stdout = stdout.String()
		result.Stderr = stderr.String()

		if cmd.ProcessState != nil {
			result.ExitCode = cmd.ProcessState.ExitCode()
		}

		stats := parser.Status.Stats()

		if stats.State == "" || stats.State == "starting" {
			parser.Status.SetState("finished")
			stats = parser.Status.Stats()
		}

		result.State = stats.State
		result.Speed = stats.Speed
		result.SpeedHPS = stats.SpeedHPS
		result.Progress = stats.Progress
		result.Recovered = stats.Recovered
		result.ETA = stats.ETA

		// Hashcat uses exit status 1 for a normal exhausted attack.
		// Cracked and Exhausted are successful workflow outcomes,
		// not process failures.
		if result.State == "Cracked" || result.State == "Exhausted" {
			result.Error = nil
		}

		r.mu.Lock()
		r.cmd = nil
		close(r.done)
		r.mu.Unlock()
	}()

	return result, parser.Status, nil
}

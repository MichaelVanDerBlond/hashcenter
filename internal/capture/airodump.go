package capture

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/system"
)

type Airodump struct{}

func (Airodump) Name() string {
	return "airodump-ng"
}

func (Airodump) Capture(job Job) (Result, error) {

	start := time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), job.Duration)
	defer cancel()

	cmd, err := system.Start(
		ctx,
		"airodump-ng",
		"--write", job.Output,
		"--output-format", "pcap",
		job.Interface,
	)
	if err != nil {
		return Result{}, err
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Wait()

	// Если процесс завершился из-за таймаута контекста —
	// это штатное завершение захвата.
	if ctx.Err() == context.DeadlineExceeded {
		err = nil
	}

	if err != nil {
		return Result{}, err
	}

	output := job.Output + "-01.cap"

	if st, statErr := os.Stat(output); statErr != nil {
		return Result{}, fmt.Errorf("capture file not created: %s", output)
	} else if st.Size() == 0 {
		return Result{}, fmt.Errorf("capture file is empty: %s", output)
	}

	return Result{
		Started:  start,
		Finished: time.Now(),
		Duration: time.Since(start),
		Output:   output,
		Success:  true,
	}, nil
}

package attack

import (
	"context"
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
)

func Execute(ctx context.Context, job Job) error {

	run, err := jobs.DefaultManager().Run(ctx, job)
	if err != nil {
		return err
	}

	snapshot := run.Status.Snapshot()

	fmt.Println("Session ID :", run.SessionID)
	fmt.Println("Status     :", snapshot.State)
	fmt.Println("Exit code  :", run.Result.ExitCode)
	fmt.Printf("Duration   : %s\n", run.Result.Duration)
	fmt.Println()

	if run.Result.Stdout != "" {
		fmt.Println("STDOUT")
		fmt.Println("----------------------------------------")
		fmt.Print(run.Result.Stdout)
		fmt.Println()
	}

	if run.Result.Stderr != "" {
		fmt.Println("STDERR")
		fmt.Println("----------------------------------------")
		fmt.Print(run.Result.Stderr)
		fmt.Println()
	}

	return nil
}

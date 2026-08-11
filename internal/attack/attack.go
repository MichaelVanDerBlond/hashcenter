package attack

import (
	"context"
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/jobs"
	"github.com/MichaelVanDerBlond/hashcenter/internal/workflow"
)

func ExecuteResult(ctx context.Context, job Job) (*jobs.Result, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	wf := workflow.NewContext()

	wf.HashFile = job.HashFile
	wf.Dictionary = job.Dictionary
	wf.HashMode = job.HashMode
	wf.AttackMode = job.AttackMode

	wf.Rule = job.Rule
	wf.Mask = job.Mask
	wf.SessionName = job.SessionName
	wf.Device = job.Device
	wf.Workload = job.Workload

	wf.Extra = append([]string(nil), job.Extra...)

	engine := workflow.NewDefaultPipeline()

	if err := engine.Run(ctx, wf); err != nil {
		return nil, err
	}

	if wf.Result == nil {
		return nil, fmt.Errorf("workflow finished without result")
	}

	return wf.Result, nil
}

func Execute(ctx context.Context, job Job) error {
	result, err := ExecuteResult(ctx, job)
	if err != nil {
		return err
	}

	snapshot := result.Status.Snapshot()

	fmt.Println("Session ID :", result.SessionID)
	fmt.Println("Status     :", snapshot.State)
	fmt.Println("Exit code  :", result.Result.ExitCode)
	fmt.Printf("Duration   : %s\n", result.Result.Duration)
	fmt.Println()

	if result.Result.Stdout != "" {
		fmt.Println("STDOUT")
		fmt.Println("----------------------------------------")
		fmt.Print(result.Result.Stdout)
		fmt.Println()
	}

	if result.Result.Stderr != "" {
		fmt.Println("STDERR")
		fmt.Println("----------------------------------------")
		fmt.Print(result.Result.Stderr)
		fmt.Println()
	}

	return nil
}

package attack

import (
	"context"
	"fmt"

	"github.com/MichaelVanDerBlond/hashcenter/internal/workflow"
)

func Execute(ctx context.Context, job Job) error {

	wf := workflow.NewContext()

	wf.HashFile = job.HashFile
	wf.Dictionary = job.Dictionary
	wf.HashMode = job.HashMode
	wf.AttackMode = job.AttackMode
	wf.Extra = append([]string(nil), job.Extra...)

	engine := workflow.NewDefaultPipeline()

	if err := engine.Run(ctx, wf); err != nil {
		return err
	}

	if wf.Result == nil {
		return fmt.Errorf("workflow finished without result")
	}

	snapshot := wf.Result.Status.Snapshot()

	fmt.Println("Session ID :", wf.Result.SessionID)
	fmt.Println("Status     :", snapshot.State)
	fmt.Println("Exit code  :", wf.Result.Result.ExitCode)
	fmt.Printf("Duration   : %s\n", wf.Result.Result.Duration)
	fmt.Println()

	if wf.Result.Result.Stdout != "" {
		fmt.Println("STDOUT")
		fmt.Println("----------------------------------------")
		fmt.Print(wf.Result.Result.Stdout)
		fmt.Println()
	}

	if wf.Result.Result.Stderr != "" {
		fmt.Println("STDERR")
		fmt.Println("----------------------------------------")
		fmt.Print(wf.Result.Result.Stderr)
		fmt.Println()
	}

	return nil
}

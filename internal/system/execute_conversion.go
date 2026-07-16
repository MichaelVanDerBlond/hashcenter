package system

import (
	"fmt"
	"os"
	"path/filepath"
)

func ExecuteConversion(input string) (*ConversionPlan, *CommandResult, error) {

	t := DetectCaptureFile(input)
	plan := BuildConversionPlan(input, t)

	if !plan.ConversionNeeded {
		return &plan, nil, nil
	}

	if !plan.BackendAvailable {
		return &plan, nil, fmt.Errorf("backend %q not found", plan.Backend)
	}

	_ = os.MkdirAll(filepath.Dir(plan.OutputFile), 0755)
	_ = os.Remove(plan.OutputFile)

	result, err := RunCommand(plan.Command)
	if err != nil {
		return &plan, result, err
	}

	if _, err := os.Stat(plan.OutputFile); err != nil {
		return &plan, result, fmt.Errorf("conversion finished but output file not found: %s", plan.OutputFile)
	}

	return &plan, result, nil
}

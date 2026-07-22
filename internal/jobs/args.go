package jobs

import "fmt"

func buildArgs(job Job) []string {
	args := []string{
		"-m", fmt.Sprint(job.HashMode),
		"-a", fmt.Sprint(job.AttackMode),
	}

	if job.SessionName != "" {
		args = append(args, "--session", job.SessionName)
	}

	if job.Device != "" {
		args = append(args, "-d", job.Device)
	}

	if job.Workload > 0 {
		args = append(args, "-w", fmt.Sprint(job.Workload))
	}

	if job.Rule != "" {
		args = append(args, "-r", job.Rule)
	}

	args = append(args, job.HashFile)

	switch job.AttackMode {
	case 3:
		if job.Mask != "" {
			args = append(args, job.Mask)
		}
	default:
		if job.Dictionary != "" {
			args = append(args, job.Dictionary)
		}
		if job.Mask != "" {
			args = append(args, job.Mask)
		}
	}

	args = append(args, job.Extra...)

	return args
}

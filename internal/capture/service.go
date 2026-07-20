package capture

import (
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/device"
	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

func Execute(
	s *session.Session,
	channel int,
	output string,
	duration time.Duration,
) (Result, error) {

	manager, err := device.Default()
	if err != nil {
		s.Fail(err)
		return Result{}, err
	}

	s.Interface = manager.Interface()
	s.Channel = channel

	if err := manager.Validate(); err != nil {
		s.Fail(err)
		return Result{}, err
	}

	if err := manager.EnableMonitor(); err != nil {
		s.Fail(err)
		return Result{}, err
	}

	defer func() {
		_ = manager.DisableMonitor()
	}()

	if err := manager.SetChannel(channel); err != nil {
		s.Fail(err)
		return Result{}, err
	}

	backend := Default()

	s.Backend = backend.Name()

	s.Start()

	job := Job{
		Interface: manager.Interface(),
		Channel:   channel,
		Output:    output,
		Duration:  duration,
		Backend:   backend.Name(),
	}

	result, err := Run(backend, job)
	if err != nil {
		s.Fail(err)
		return Result{}, err
	}

	s.CaptureFile = result.Output
	s.Finish()

	return result, nil
}

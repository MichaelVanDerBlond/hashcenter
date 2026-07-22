package hashcat

import "errors"

func (r *Runner) Wait(result *Result) (*Result, error) {

	if result == nil {
		return nil, errors.New("nil result")
	}

	r.mu.Lock()
	done := r.done
	r.mu.Unlock()

	if done != nil {
		<-done
	}

	return result, result.Error
}

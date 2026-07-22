package session

import "time"

func (s *Session) Start() {
	if s == nil {
		return
	}

	s.State = Running
	s.Started = time.Now()
	s.Finished = time.Time{}
	s.Error = ""
}

func (s *Session) Finish() {
	if s == nil {
		return
	}

	s.State = Finished
	s.Finished = time.Now()
}

func (s *Session) Fail(err error) {
	if s == nil {
		return
	}

	s.State = Failed
	s.Finished = time.Now()

	if err != nil {
		s.Error = err.Error()
	} else {
		s.Error = ""
	}
}

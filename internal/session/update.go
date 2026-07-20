package session

import "time"

func (s *Session) Start() {
	s.State = Running
	s.Started = time.Now()
}

func (s *Session) Finish() {
	s.State = Finished
	s.Finished = time.Now()
}

func (s *Session) Fail(err error) {

	s.State = Failed
	s.Finished = time.Now()

	if err != nil {
		s.Error = err.Error()
	}
}

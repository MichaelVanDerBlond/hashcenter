package hashcatruntime

import (
	"bufio"
	"context"
	"io"
)

type Client struct {
	session *Session
}

func NewClient(session *Session) *Client {
	return &Client{
		session: session,
	}
}

func (c *Client) Consume(ctx context.Context, r io.Reader) error {

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		snapshot, err := ParseStatus(scanner.Bytes())
		if err != nil {
			continue
		}

		c.session.Update(func(s *Snapshot) {

			started := s.Started

			*s = snapshot

			if !started.IsZero() && s.Started.IsZero() {
				s.Started = started
			}
		})
	}

	return scanner.Err()
}

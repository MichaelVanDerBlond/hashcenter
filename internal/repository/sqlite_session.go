package repository

import (
	"database/sql"
	"errors"
	"time"

	"github.com/MichaelVanDerBlond/hashcenter/internal/session"
)

const timeLayout = "2006-01-02 15:04:05"

type SQLiteSessionRepository struct {
	db *sql.DB
}

func NewSQLiteSessionRepository(db *sql.DB) *SQLiteSessionRepository {
	return &SQLiteSessionRepository{
		db: db,
	}
}

func formatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(timeLayout)
}

func parseTime(s string) time.Time {
	if s == "" {
		return time.Time{}
	}

	t, err := time.Parse(timeLayout, s)
	if err != nil {
		return time.Time{}
	}

	return t
}

func (r *SQLiteSessionRepository) Save(s *session.Session) error {
	if s == nil {
		return errors.New("nil session")
	}

	_, err := r.db.Exec(`
INSERT INTO sessions(
	id,
	state,
	started,
	finished,
	interface,
	channel,
	backend,
	capture_file,
	error
)
VALUES(?,?,?,?,?,?,?,?,?)
`,
		s.ID,
		string(s.State),
		formatTime(s.Started),
		formatTime(s.Finished),
		s.Interface,
		s.Channel,
		s.Backend,
		s.CaptureFile,
		s.Error,
	)

	return err
}

func (r *SQLiteSessionRepository) Update(s *session.Session) error {
	if s == nil {
		return errors.New("nil session")
	}

	_, err := r.db.Exec(`
UPDATE sessions
SET
	state=?,
	started=?,
	finished=?,
	interface=?,
	channel=?,
	backend=?,
	capture_file=?,
	error=?
WHERE id=?
`,
		string(s.State),
		formatTime(s.Started),
		formatTime(s.Finished),
		s.Interface,
		s.Channel,
		s.Backend,
		s.CaptureFile,
		s.Error,
		s.ID,
	)

	return err
}

func (r *SQLiteSessionRepository) Get(id string) (*session.Session, error) {

	var s session.Session

	var started string
	var finished string

	err := r.db.QueryRow(`
SELECT
	id,
	state,
	started,
	finished,
	interface,
	channel,
	backend,
	capture_file,
	error
FROM sessions
WHERE id=?
`,
		id,
	).Scan(
		&s.ID,
		&s.State,
		&started,
		&finished,
		&s.Interface,
		&s.Channel,
		&s.Backend,
		&s.CaptureFile,
		&s.Error,
	)

	if err != nil {
		return nil, err
	}

	s.Started = parseTime(started)
	s.Finished = parseTime(finished)

	return &s, nil
}

func (r *SQLiteSessionRepository) List() ([]session.Session, error) {

	rows, err := r.db.Query(`
SELECT
	id,
	state,
	started,
	finished,
	interface,
	channel,
	backend,
	capture_file,
	error
FROM sessions
ORDER BY started DESC
`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sessions []session.Session

	for rows.Next() {

		var s session.Session

		var started string
		var finished string

		err := rows.Scan(
			&s.ID,
			&s.State,
			&started,
			&finished,
			&s.Interface,
			&s.Channel,
			&s.Backend,
			&s.CaptureFile,
			&s.Error,
		)
		if err != nil {
			return nil, err
		}

		s.Started = parseTime(started)
		s.Finished = parseTime(finished)

		sessions = append(sessions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

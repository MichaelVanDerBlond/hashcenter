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
	if r == nil || r.db == nil {
		return errors.New("nil repository")
	}

	if s == nil {
		return errors.New("nil session")
	}

	if s.ID == "" {
		return errors.New("empty session id")
	}

	_, err := r.db.Exec(`
INSERT INTO sessions(
	id,
	state,
	created,
	started,
	finished,
	interface,
	channel,
	backend,
	capture_file,
	error
)
VALUES(?,?,?,?,?,?,?,?,?,?)
`,
		s.ID,
		string(s.State),
		formatTime(s.Created),
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
	if r == nil || r.db == nil {
		return errors.New("nil repository")
	}

	if s == nil {
		return errors.New("nil session")
	}

	if s.ID == "" {
		return errors.New("empty session id")
	}

	_, err := r.db.Exec(`
UPDATE sessions
SET
	state=?,
	created=?,
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
		formatTime(s.Created),
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
	if r == nil || r.db == nil {
		return nil, errors.New("nil repository")
	}

	if id == "" {
		return nil, sql.ErrNoRows
	}

	var s session.Session

	var created string
	var started string
	var finished string

	err := r.db.QueryRow(`
SELECT
	id,
	state,
	created,
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
		&created,
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

	s.Created = parseTime(created)
	s.Started = parseTime(started)
	s.Finished = parseTime(finished)

	return &s, nil
}

func (r *SQLiteSessionRepository) List() ([]session.Session, error) {
	if r == nil || r.db == nil {
		return nil, errors.New("nil repository")
	}

	rows, err := r.db.Query(`
SELECT
	id,
	state,
	created,
	started,
	finished,
	interface,
	channel,
	backend,
	capture_file,
	error
FROM sessions
ORDER BY created DESC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []session.Session

	for rows.Next() {
		var s session.Session

		var created string
		var started string
		var finished string

		if err := rows.Scan(
			&s.ID,
			&s.State,
			&created,
			&started,
			&finished,
			&s.Interface,
			&s.Channel,
			&s.Backend,
			&s.CaptureFile,
			&s.Error,
		); err != nil {
			return nil, err
		}

		s.Created = parseTime(created)
		s.Started = parseTime(started)
		s.Finished = parseTime(finished)

		sessions = append(sessions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

func (r *SQLiteSessionRepository) Delete(id string) error {
	if r == nil || r.db == nil {
		return errors.New("nil repository")
	}

	if id == "" {
		return nil
	}

	_, err := r.db.Exec(`
DELETE FROM sessions
WHERE id=?
`,
		id,
	)

	return err
}

func (r *SQLiteSessionRepository) Exists(id string) (bool, error) {
	if r == nil || r.db == nil {
		return false, errors.New("nil repository")
	}

	if id == "" {
		return false, nil
	}

	var count int

	err := r.db.QueryRow(`
SELECT COUNT(*)
FROM sessions
WHERE id=?
`,
		id,
	).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

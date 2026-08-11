package task

import (
	"database/sql"
	"errors"
	"time"
)

const taskTimeLayout = "2006-01-02 15:04:05"

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	return &SQLiteStore{
		db: db,
	}
}

func taskFormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.Format(taskTimeLayout)
}

func taskParseTime(value string) time.Time {
	if value == "" {
		return time.Time{}
	}

	t, err := time.Parse(taskTimeLayout, value)
	if err != nil {
		return time.Time{}
	}

	return t
}

func (s *SQLiteStore) Save(t *Task) error {
	if s == nil || s.db == nil {
		return errors.New("nil task store")
	}

	if t == nil {
		return errors.New("nil task")
	}

	if t.ID == "" {
		return errors.New("empty task id")
	}

	_, err := s.db.Exec(`
INSERT INTO tasks(
	id,
	session_id,
	type,
	state,
	created,
	started,
	finished,
	hash_file,
	dictionary,
	hash_mode,
	attack_mode,
	rule,
	mask,
	device,
	workload,
	session_name,
	error
)
VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`,
		t.ID,
		t.SessionID,
		string(t.Type),
		string(t.State),
		taskFormatTime(t.Created),
		taskFormatTime(t.Started),
		taskFormatTime(t.Finished),
		t.HashFile,
		t.Dictionary,
		t.HashMode,
		t.AttackMode,
		t.Rule,
		t.Mask,
		t.Device,
		t.Workload,
		t.SessionName,
		t.Error,
	)

	return err
}

func (s *SQLiteStore) Update(t *Task) error {
	if s == nil || s.db == nil {
		return errors.New("nil task store")
	}

	if t == nil {
		return errors.New("nil task")
	}

	if t.ID == "" {
		return errors.New("empty task id")
	}

	_, err := s.db.Exec(`
UPDATE tasks
SET
	session_id=?,
	type=?,
	state=?,
	created=?,
	started=?,
	finished=?,
	hash_file=?,
	dictionary=?,
	hash_mode=?,
	attack_mode=?,
	rule=?,
	mask=?,
	device=?,
	workload=?,
	session_name=?,
	error=?
WHERE id=?
`,
		t.SessionID,
		string(t.Type),
		string(t.State),
		taskFormatTime(t.Created),
		taskFormatTime(t.Started),
		taskFormatTime(t.Finished),
		t.HashFile,
		t.Dictionary,
		t.HashMode,
		t.AttackMode,
		t.Rule,
		t.Mask,
		t.Device,
		t.Workload,
		t.SessionName,
		t.Error,
		t.ID,
	)

	return err
}

func (s *SQLiteStore) Delete(id string) error {
	if s == nil || s.db == nil {
		return errors.New("nil task store")
	}

	if id == "" {
		return nil
	}

	_, err := s.db.Exec(`
DELETE FROM tasks
WHERE id=?
`, id)

	return err
}

func (s *SQLiteStore) Load() ([]*Task, error) {
	if s == nil || s.db == nil {
		return nil, errors.New("nil task store")
	}

	rows, err := s.db.Query(`
SELECT
	id,
	session_id,
	type,
	state,
	created,
	started,
	finished,
	hash_file,
	dictionary,
	hash_mode,
	attack_mode,
	rule,
	mask,
	device,
	workload,
	session_name,
	error
FROM tasks
ORDER BY created ASC
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	for rows.Next() {
		var t Task
		var typ string
		var state string
		var created string
		var started string
		var finished string

		if err := rows.Scan(
			&t.ID,
			&t.SessionID,
			&typ,
			&state,
			&created,
			&started,
			&finished,
			&t.HashFile,
			&t.Dictionary,
			&t.HashMode,
			&t.AttackMode,
			&t.Rule,
			&t.Mask,
			&t.Device,
			&t.Workload,
			&t.SessionName,
			&t.Error,
		); err != nil {
			return nil, err
		}

		t.Type = Type(typ)
		t.State = State(state)
		t.Created = taskParseTime(created)
		t.Started = taskParseTime(started)
		t.Finished = taskParseTime(finished)

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

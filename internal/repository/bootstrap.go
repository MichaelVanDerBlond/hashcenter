package repository

import (
	"os"
	"path/filepath"

	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
	"github.com/MichaelVanDerBlond/hashcenter/internal/scanner"
)

func (r *DictionaryRepository) Bootstrap(dicts []models.Dictionary) error {

	for _, d := range dicts {

		if err := r.Sync(d); err != nil {
			return err
		}

		var favorite int

		err := r.db.QueryRow(`
SELECT favorite
FROM dictionaries
WHERE path=?
`,
			d.Path,
		).Scan(&favorite)

		if err != nil {
			return err
		}

		if favorite == 1 {
			continue
		}

		priority, ok := scanner.DefaultFavorites[filepath.Base(d.Path)]
		if !ok {
			continue
		}

		_, err = r.db.Exec(`
UPDATE dictionaries
SET favorite=1,
priority=?
WHERE path=?
`,
			priority,
			d.Path,
		)

		if err != nil {
			return err
		}

	}

	return nil
}

func (r *DictionaryRepository) RemoveMissing() error {

	rows, err := r.db.Query(`
SELECT path
FROM dictionaries
`)
	if err != nil {
		return err
	}

	defer rows.Close()

	var path string

	for rows.Next() {

		if err := rows.Scan(&path); err != nil {
			return err
		}

		if _, err := os.Stat(path); err == nil {
			continue
		}

		if _, err := r.db.Exec(`
DELETE FROM dictionaries
WHERE path=?
`,
			path,
		); err != nil {
			return err
		}

	}

	return rows.Err()
}

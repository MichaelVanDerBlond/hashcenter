package repository

import (
	"database/sql"
	"errors"

	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
)

type DictionaryRepository struct {
	db *sql.DB
}

func NewDictionaryRepository(db *sql.DB) *DictionaryRepository {
	return &DictionaryRepository{db: db}
}

func (r *DictionaryRepository) Sync(d models.Dictionary) error {
	if r == nil || r.db == nil {
		return errors.New("nil repository")
	}

	_, err := r.db.Exec(`
INSERT INTO dictionaries(path,name,size)
VALUES(?,?,?)
ON CONFLICT(path)
DO UPDATE SET
	name=excluded.name,
	size=excluded.size
`,
		d.Path,
		d.Name,
		d.Size,
	)

	return err
}

func (r *DictionaryRepository) Favorites(search string) ([]models.Dictionary, error) {
	return r.list(true, search)
}

func (r *DictionaryRepository) Others(search string) ([]models.Dictionary, error) {
	return r.list(false, search)
}

func (r *DictionaryRepository) list(favorite bool, search string) ([]models.Dictionary, error) {
	if r == nil || r.db == nil {
		return nil, errors.New("nil repository")
	}

	query := `
SELECT path,name,size,favorite,priority
FROM dictionaries
WHERE favorite=?
`

	args := []any{0}
	if favorite {
		args[0] = 1
	}

	if search != "" {
		query += " AND LOWER(name) LIKE LOWER(?)"
		args = append(args, "%"+search+"%")
	}

	if favorite {
		query += " ORDER BY priority ASC,name ASC"
	} else {
		query += " ORDER BY size ASC,name ASC"
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Dictionary

	for rows.Next() {
		var d models.Dictionary
		var fav int

		if err := rows.Scan(
			&d.Path,
			&d.Name,
			&d.Size,
			&fav,
			&d.Priority,
		); err != nil {
			return nil, err
		}

		d.Favorite = fav == 1
		result = append(result, d)
	}

	return result, rows.Err()
}

func (r *DictionaryRepository) FavoriteCount() (int, error) {

	var count int

	err := r.db.QueryRow(`
SELECT COUNT(*)
FROM dictionaries
WHERE favorite=1
`).Scan(&count)

	return count, err
}

func (r *DictionaryRepository) OtherCount() (int, error) {

	var count int

	err := r.db.QueryRow(`
SELECT COUNT(*)
FROM dictionaries
WHERE favorite=0
`).Scan(&count)

	return count, err
}

func (r *DictionaryRepository) SetFavorite(path string, favorite bool) error {
	if r == nil || r.db == nil {
		return errors.New("nil repository")
	}

	value := 0
	priority := 0

	if favorite {
		value = 1
	}

	_, err := r.db.Exec(`
UPDATE dictionaries
SET
	favorite=?,
	priority=?
WHERE path=?
`,
		value,
		priority,
		path,
	)

	return err
}

func (r *DictionaryRepository) Get(path string) (*models.Dictionary, error) {
	if r == nil || r.db == nil {
		return nil, errors.New("nil repository")
	}

	var d models.Dictionary
	var favoriteValue int

	err := r.db.QueryRow(`
SELECT
	path,
	name,
	size,
	favorite,
	priority
FROM dictionaries
WHERE path=?
`,
		path,
	).Scan(
		&d.Path,
		&d.Name,
		&d.Size,
		&favoriteValue,
		&d.Priority,
	)

	if err != nil {
		return nil, err
	}

	d.Favorite = favoriteValue == 1

	return &d, nil
}

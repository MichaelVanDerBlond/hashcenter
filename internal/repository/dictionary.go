package repository

import (
	"database/sql"

	"github.com/MichaelVanDerBlond/hashcenter/internal/models"
)

type DictionaryRepository struct {
	db *sql.DB
}

func NewDictionaryRepository(db *sql.DB) *DictionaryRepository {
	return &DictionaryRepository{db: db}
}

func (r *DictionaryRepository) Sync(d models.Dictionary) error {

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

	query := `
SELECT path,name,size,favorite,priority
FROM dictionaries
WHERE favorite=1
`

	args := []any{}

	if search != "" {
		query += " AND LOWER(name) LIKE LOWER(?)"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY priority ASC,name ASC"

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

func (r *DictionaryRepository) Others(search string) ([]models.Dictionary, error) {

	query := `
SELECT path,name,size,favorite,priority
FROM dictionaries
WHERE favorite=0
`

	args := []any{}

	if search != "" {
		query += " AND LOWER(name) LIKE LOWER(?)"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY size ASC,name ASC"

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

	value := 0
	if favorite {
		value = 1
	}

	_, err := r.db.Exec(`
UPDATE dictionaries
SET favorite=?
WHERE path=?
`,
		value,
		path,
	)

	return err
}

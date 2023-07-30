package sqlite

import (
	"database/sql"
	"fmt"
	"url-shortener/internal/http-server/handlers/url/getAll"
	"url-shortener/internal/storage"

	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const operation = "storage.sqlite.New"
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS url(
			id INTEGER PRIMARY KEY,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUrl(urlToSave string, alias string) (int64, error) {
	const operation = "storage.sqlite.SaveUrl"

	stmt, err := s.db.Prepare("INSERT INTO url(url, alias) VALUES(?, ?)")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	response, err := stmt.Exec(urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s: %w", operation, storage.ErrorUrlExists)
		}

		return 0, fmt.Errorf("%s: %w", operation, err)
	}

	id, err := response.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id %w", operation, err)
	}

	return id, nil
}

func (s *Storage) GetUrls() ([]getAll.Urls, error) {
	const operation = "storage.sqlite.GetUrls"

	stmt, err := s.db.Prepare("SELECT id, alias, url FROM url;")
	if err != nil {
		return nil, fmt.Errorf("%s: prepare statement: %w", operation, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, fmt.Errorf("%s: query: %w", operation, err)
	}
	defer rows.Close()

	var urls []getAll.Urls

	for rows.Next() {
		var url getAll.Urls
		err := rows.Scan(&url.ID, &url.Alias, &url.URL)
		if err != nil {
			return nil, fmt.Errorf("%s: scan row: %w", operation, err)
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: rows iteration: %w", operation, err)
	}

	return urls, nil
}

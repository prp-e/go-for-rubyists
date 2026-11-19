package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const schema = `
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL CHECK (length(title) BETWEEN 1 AND 120),
    body TEXT NOT NULL CHECK (length(body) BETWEEN 1 AND 10000)
) STRICT;
`

type post struct {
	ID    int64
	Title string
	Body  string
}

func openStore(ctx context.Context, filename string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	// One connection serializes this local blog's reads and writes.
	db.SetMaxOpenConns(1)
	if _, err := db.ExecContext(ctx, schema); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func listPosts(ctx context.Context, db *sql.DB) ([]post, error) {
	rows, err := db.QueryContext(ctx, "SELECT id, title FROM posts ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []post
	for rows.Next() {
		var p post
		if err := rows.Scan(&p.ID, &p.Title); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func getPost(ctx context.Context, db *sql.DB, id int64) (post, error) {
	p := post{ID: id}
	err := db.QueryRowContext(ctx,
		"SELECT title, body FROM posts WHERE id = ?", id,
	).Scan(&p.Title, &p.Body)
	return p, err
}

func changed(result sql.Result, err error) error {
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err == nil && n == 0 {
		return sql.ErrNoRows
	}
	return err
}

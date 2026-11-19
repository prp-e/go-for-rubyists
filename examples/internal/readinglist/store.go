package readinglist

import (
	"errors"
	"strings"
	"sync"
	"unicode/utf8"
)

type Book struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Read  bool   `json:"read"`
}

var ErrInvalidTitle = errors.New("title must contain 1 to 200 Unicode code points")

// Store is an in-memory catalog. Its zero value is ready to use.
type Store struct {
	mu    sync.RWMutex
	books []Book
}

func (s *Store) Add(title string) (Book, error) {
	title = strings.TrimSpace(title)
	if title == "" || !utf8.ValidString(title) || utf8.RuneCountInString(title) > 200 {
		return Book{}, ErrInvalidTitle
	}
	s.mu.Lock()
	defer s.mu.Unlock()

	book := Book{ID: int64(len(s.books)) + 1, Title: title}
	s.books = append(s.books, book)
	return book, nil
}

func (s *Store) List() []Book {
	s.mu.RLock()
	defer s.mu.RUnlock()

	snapshot := make([]Book, len(s.books))
	copy(snapshot, s.books)
	return snapshot
}

func (s *Store) Find(id int64) (Book, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if id <= 0 || id > int64(len(s.books)) {
		return Book{}, false
	}
	return s.books[id-1], true
}

func (s *Store) MarkRead(id int64) (Book, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if id <= 0 || id > int64(len(s.books)) {
		return Book{}, false
	}
	s.books[id-1].Read = true
	return s.books[id-1], true
}

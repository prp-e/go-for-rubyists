package readinglist

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func request(handler http.Handler, method, path, contentType, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	return response
}

func TestReadingList(t *testing.T) {
	handler := Handler(&Store{})
	empty := request(handler, http.MethodGet, "/books", "", "")
	if empty.Code != 200 || empty.Body.String() != "[]\n" {
		t.Fatalf("empty list: %d %q", empty.Code, empty.Body)
	}

	created := request(handler, http.MethodPost, "/books", "application/json", `{"title":"  Kindred  "}`)
	if created.Code != 201 || created.Header().Get("Location") != "/books/1" {
		t.Fatalf("create: %d %v %s", created.Code, created.Header(), created.Body)
	}
	checkBook := func(response *httptest.ResponseRecorder, wantStatus int, wantRead bool) {
		t.Helper()
		if response.Code != wantStatus {
			t.Fatalf("status = %d; want %d: %s", response.Code, wantStatus, response.Body)
		}
		var book Book
		if err := json.Unmarshal(response.Body.Bytes(), &book); err != nil {
			t.Fatal(err)
		}
		if book != (Book{ID: 1, Title: "Kindred", Read: wantRead}) {
			t.Fatalf("unexpected book: %+v", book)
		}
	}
	checkBook(created, 201, false)
	checkBook(request(handler, http.MethodGet, "/books/1", "", ""), 200, false)
	for range 2 {
		checkBook(request(handler, http.MethodPut, "/books/1/read", "", ""), 200, true)
	}
	checkBook(request(handler, http.MethodGet, "/books/1", "", ""), 200, true)
}

func TestBadRequests(t *testing.T) {
	handler := Handler(&Store{})
	for _, tc := range []struct {
		name, method, path, contentType, body string
		status                                int
	}{
		{"unknown book", "GET", "/books/9", "", "", 404},
		{"invalid ID", "GET", "/books/zero", "", "", 400},
		{"wrong method", "DELETE", "/books/1", "", "", 405},
		{"wrong media type", "POST", "/books", "text/plain", "", 415},
		{"blank title", "POST", "/books", "application/json", `{"title":"  "}`, 400},
		{"missing title", "POST", "/books", "application/json", `{}`, 400},
		{"null title", "POST", "/books", "application/json", `{"title":null}`, 400},
		{"invalid JSON", "POST", "/books", "application/json", `{`, 400},
		{"unknown field", "POST", "/books", "application/json", `{"title":"A","read":true}`, 400},
		{"two objects", "POST", "/books", "application/json", `{"title":"A"} {}`, 400},
		{"long title", "POST", "/books", "application/json", `{"title":"` + strings.Repeat("λ", 201) + `"}`, 400},
		{"large body", "POST", "/books", "application/json", `{"title":"A"}` + strings.Repeat(" ", 4096), 413},
	} {
		t.Run(tc.name, func(t *testing.T) {
			response := request(handler, tc.method, tc.path, tc.contentType, tc.body)
			if response.Code != tc.status {
				t.Fatalf("status = %d; want %d: %s", response.Code, tc.status, response.Body)
			}
		})
	}
	if response := request(handler, "GET", "/books", "", ""); response.Body.String() != "[]\n" {
		t.Fatalf("invalid requests created books: %s", response.Body)
	}
}

func TestConcurrentAddAndSnapshot(t *testing.T) {
	const size = 30
	var store Store
	var wg sync.WaitGroup
	for range size {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := store.Add("Solaris"); err != nil {
				t.Error(err)
			}
			store.List()
		}()
	}
	wg.Wait()
	books := store.List()
	if len(books) != size {
		t.Fatalf("got %d books; want %d", len(books), size)
	}
	for i, book := range books {
		if book.ID != int64(i+1) {
			t.Fatalf("book %d has unexpected ID %d", i, book.ID)
		}
	}
	books[0].Title = "changed"
	if book, ok := store.Find(1); !ok || book.Title != "Solaris" {
		t.Fatalf("snapshot leaked a mutation: %+v", book)
	}
}

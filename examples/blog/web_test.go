package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
)

func visit(h http.Handler, method, path string, form url.Values, crossOrigin bool) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(form.Encode()))
	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if crossOrigin {
		req.Header.Set("Sec-Fetch-Site", "cross-site")
	}
	response := httptest.NewRecorder()
	h.ServeHTTP(response, req)
	return response
}

func expect(t *testing.T, response *httptest.ResponseRecorder, status int, text string) {
	t.Helper()
	if response.Code != status || !strings.Contains(response.Body.String(), text) {
		t.Fatalf("response = %d %q; want %d containing %q", response.Code, response.Body.String(), status, text)
	}
}

func TestBlog(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "test.db")
	db, err := openStore(context.Background(), filename)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })
	h := routes(db)
	expect(t, visit(h, "GET", "/", nil, false), 200, "No posts yet")
	expect(t, visit(h, "GET", "/posts/new", nil, false), 200, "New post")
	expect(t, visit(h, "GET", "/static/blog.css", nil, false), 200, "white-space")
	expect(t, visit(h, "POST", "/posts", url.Values{"title": {"draft"}}, false), 422, "draft")
	expect(t, visit(h, "POST", "/posts", nil, false), 415, "Expected an HTML form")
	form := url.Values{"title": {"Go & Ruby"}, "body": {"<em>Plain text</em>"}}
	expect(t, visit(h, "POST", "/posts", form, true), 403, "")
	created := visit(h, "POST", "/posts", form, false)
	expect(t, created, 303, "")
	path := created.Header().Get("Location")
	if path != "/posts/1" {
		t.Fatalf("Location = %q", path)
	}
	expect(t, visit(h, "GET", "/", nil, false), 200, "Go &amp; Ruby")
	expect(t, visit(h, "GET", path, nil, false), 200, "&lt;em&gt;Plain text&lt;/em&gt;")
	expect(t, visit(h, "GET", path+"/edit", nil, false), 200, "Go &amp; Ruby")
	form.Set("title", "Revised title")
	expect(t, visit(h, "POST", path+"/edit", form, false), 303, "")
	expect(t, visit(h, "GET", path+"/delete", nil, false), 405, "")

	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
	reopened, err := openStore(context.Background(), filename)
	if err != nil {
		t.Fatal(err)
	}
	defer reopened.Close()
	h = routes(reopened)
	expect(t, visit(h, "GET", path, nil, false), 200, "Revised title")
	expect(t, visit(h, "POST", path+"/delete", nil, true), 403, "")
	expect(t, visit(h, "GET", path, nil, false), 200, "Revised title")
	expect(t, visit(h, "POST", path+"/delete", nil, false), 303, "")
	expect(t, visit(h, "GET", path, nil, false), 404, "Post not found")
	expect(t, visit(h, "GET", "/", nil, false), 200, "No posts yet")
}

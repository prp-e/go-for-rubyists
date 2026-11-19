package main

import (
	"bytes"
	"database/sql"
	"embed"
	"errors"
	"html/template"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed templates/*.html static/blog.css
var assets embed.FS

var pages = template.Must(template.ParseFS(assets, "templates/*.html"))

type blog struct{ db *sql.DB }

type page struct {
	Title   string
	Posts   []post
	Post    post
	Action  string
	Message string
}

func routes(db *sql.DB) http.Handler {
	b := blog{db: db}
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.Use(http.NewCrossOriginProtection().Handler)
	r.Get("/", b.index)
	r.Get("/posts/new", b.newPost)
	r.Post("/posts", b.create)
	r.Route("/posts/{id}", func(r chi.Router) {
		r.Get("/", b.show)
		r.Get("/edit", b.edit)
		r.Post("/edit", b.update)
		r.Post("/delete", b.remove)
	})
	r.Get("/static/blog.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, assets, "static/blog.css")
	})
	return r
}

func (b blog) index(w http.ResponseWriter, r *http.Request) {
	posts, err := listPosts(r.Context(), b.db)
	if err != nil {
		fail(w, err)
		return
	}
	render(w, http.StatusOK, "index", page{Title: "Posts", Posts: posts})
}

func (b blog) newPost(w http.ResponseWriter, r *http.Request) {
	render(w, http.StatusOK, "form", page{Title: "New post", Action: "/posts"})
}

func (b blog) show(w http.ResponseWriter, r *http.Request) {
	p, err := getPost(r.Context(), b.db, postID(r))
	if err != nil {
		fail(w, err)
		return
	}
	render(w, http.StatusOK, "show", page{Title: p.Title, Post: p})
}

func (b blog) edit(w http.ResponseWriter, r *http.Request) {
	p, err := getPost(r.Context(), b.db, postID(r))
	if err != nil {
		fail(w, err)
		return
	}
	render(w, http.StatusOK, "form", editPage(p))
}

func editPage(p post) page {
	return page{Title: "Edit post", Post: p, Action: postURL(p.ID) + "/edit"}
}

func (b blog) create(w http.ResponseWriter, r *http.Request) {
	form := page{Title: "New post", Action: "/posts"}
	if !readForm(w, r, &form) {
		return
	}
	result, err := b.db.ExecContext(r.Context(),
		"INSERT INTO posts (title, body) VALUES (?, ?)", form.Post.Title, form.Post.Body)
	if err != nil {
		fail(w, err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fail(w, err)
		return
	}
	http.Redirect(w, r, postURL(id), http.StatusSeeOther)
}

func (b blog) update(w http.ResponseWriter, r *http.Request) {
	p, err := getPost(r.Context(), b.db, postID(r))
	if err != nil {
		fail(w, err)
		return
	}
	form := editPage(p)
	if !readForm(w, r, &form) {
		return
	}
	err = changed(b.db.ExecContext(r.Context(),
		"UPDATE posts SET title = ?, body = ? WHERE id = ?",
		form.Post.Title, form.Post.Body, p.ID))
	if err != nil {
		fail(w, err)
		return
	}
	http.Redirect(w, r, postURL(p.ID), http.StatusSeeOther)
}

func (b blog) remove(w http.ResponseWriter, r *http.Request) {
	err := changed(b.db.ExecContext(r.Context(), "DELETE FROM posts WHERE id = ?", postID(r)))
	if err != nil {
		fail(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func postID(r *http.Request) int64 {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil || id <= 0 {
		return 0 // The application only creates positive IDs.
	}
	return id
}

func postURL(id int64) string {
	return "/posts/" + strconv.FormatInt(id, 10)
}

func readForm(w http.ResponseWriter, r *http.Request, data *page) bool {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "application/x-www-form-urlencoded" {
		http.Error(w, "Expected an HTML form", http.StatusUnsupportedMediaType)
		return false
	}
	r.Body = http.MaxBytesReader(w, r.Body, 128<<10)
	if err := r.ParseForm(); err != nil {
		var limit *http.MaxBytesError
		if errors.As(err, &limit) {
			http.Error(w, "Form is too large", http.StatusRequestEntityTooLarge)
		} else {
			http.Error(w, "Invalid form", http.StatusBadRequest)
		}
		return false
	}
	data.Post.Title = strings.TrimSpace(r.PostForm.Get("title"))
	data.Post.Body = strings.TrimSpace(r.PostForm.Get("body"))
	if validText(data.Post.Title, 120) && validText(data.Post.Body, 10000) {
		return true
	}
	data.Message = "Enter a title of 1–120 characters and a body of 1–10,000 characters."
	render(w, http.StatusUnprocessableEntity, "form", *data)
	return false
}

func validText(s string, limit int) bool {
	return s != "" && !strings.ContainsRune(s, 0) &&
		utf8.ValidString(s) && utf8.RuneCountInString(s) <= limit
}

func render(w http.ResponseWriter, status int, name string, data page) {
	var buf bytes.Buffer
	if err := pages.ExecuteTemplate(&buf, name, data); err != nil {
		fail(w, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	_, _ = buf.WriteTo(w)
}

func fail(w http.ResponseWriter, err error) {
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}
	log.Printf("blog: %v", err)
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

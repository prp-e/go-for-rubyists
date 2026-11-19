# 14. Build a blog with Chi and SQLite

The reading list returned JSON, which works well when another program is the reader. This time the reader has a browser and wants to write something. Let's build Gopher Notes, a small blog where you can list, read, create, edit, and delete posts. The pages use plain HTML and CSS, and the posts stay in a SQLite file after you close the terminal.

## Run it, then take it apart

Get Go 1.26 or newer and a C compiler. The [SQLite driver](https://github.com/mattn/go-sqlite3) uses cgo to compile SQLite. On macOS the Command Line Tools provide Clang. On Linux use your distribution's C build tools, and on Windows use a compatible GCC toolchain. You don't have to start a separate database server.

From the repository root:

```sh
cd examples/blog
CGO_ENABLED=1 go test ./...
CGO_ENABLED=1 go run .
```

In PowerShell, set `$env:CGO_ENABLED = "1"` before running the Go commands. Open [the blog in your browser](http://127.0.0.1:8081/) and try New post. Close any server you started for Chapter 13 first, since the two examples use the same port. The program binds to loopback. This is a single-user example with no login, so keep it local unless you add authentication and authorization for the writing routes.

To build along, create a fresh directory and copy the example's `go.mod` and `go.sum` into it. Then add the files in the order below. If you're following in the clone, the tested versions are already in `examples/blog`.

## 1. Give a post somewhere to live

In [store.go](https://github.com/prp-e/go-for-rubyists/blob/book-rewrite/examples/blog/store.go), define a Go struct with an integer ID and text fields for the title and body. The blank import of `github.com/mattn/go-sqlite3` registers the driver with `database/sql`. Open it with `sql.Open("sqlite3", "blog.db")` and run this schema at startup:

```sql
CREATE TABLE IF NOT EXISTS posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL CHECK (length(title) BETWEEN 1 AND 120),
    body TEXT NOT NULL CHECK (length(body) BETWEEN 1 AND 10000)
) STRICT;
```

The ID identifies a link, so we don't want deleting the last post to make its URL point at the next one. SQLite's `AUTOINCREMENT` prevents the engine from reusing automatically assigned IDs from deleted committed rows. Ordering by descending ID puts the newest insertion first. The index query selects just IDs and titles. The individual post query reads the body as well.

Create the table once when the app starts. `IF NOT EXISTS` makes that safe on restart. It won't transform an old table into a new one when the schema changes. That's when you need a migration. Keep backups of the database if you care about the posts. The file lives in the directory where you start the program.

This example limits the pool to one connection, which lets requests queue for short operations instead of having several connections compete to write. Close every set of rows before doing more database work on that connection. Pass the request context to database calls so canceled requests can stop waiting. Put user input in SQL parameters, including the post ID. Here's the read for a single post:

```go
p := post{ID: id}
err := db.QueryRowContext(ctx,
    "SELECT title, body FROM posts WHERE id = ?", id,
).Scan(&p.Title, &p.Body)
```

If `Scan` returns `sql.ErrNoRows`, the browser should get a 404. An update or delete affecting zero rows has the same meaning. Check `RowsAffected` after the statement succeeds. Insertion uses `LastInsertId` on the result so the redirect goes to the row this request created.

## 2. Give each action a URL

In [web.go](https://github.com/prp-e/go-for-rubyists/blob/book-rewrite/examples/blog/web.go), register the routes with Chi. `GET /` lists posts and `GET /posts/new` opens the form. `POST /posts` creates the post. Chapter 13's route group adds GET routes for reading and editing one post and POST routes for saving an edit or deleting a post. A GET never changes the stored posts. HTML forms support GET and POST directly, so we can keep those four actions simple.

When someone submits the form, cap the request body before calling `ParseForm`. Read from `r.PostForm` so a query parameter can't fill in a missing body field. Trim the title and body and check their lengths. Return 422 with the populated form if the text is invalid. Returning an empty form after someone typed three paragraphs is a good way to make sure they never write a fourth.

After a successful write, return a 303 redirect. The browser then navigates with GET, which makes refreshing the new page safe. For a new post we use its new URL. After deletion we go back to the index. The router uses the standard library's cross-origin protection to reject cross-origin browser writes. That complements the use of POST forms. It doesn't provide accounts or decide who is allowed to edit a post.

## 3. Turn Go values into HTML

Create the shared layout and the three page templates in [templates](https://github.com/prp-e/go-for-rubyists/tree/book-rewrite/examples/blog/templates). If you've used ERB, the basic idea will look familiar. The template receives data from a handler. The dot means the current data, and inside a `range` the dot becomes the current item.

```html
<ul>
  {{range .Posts}}
  <li><a href="/posts/{{.ID}}">{{.Title}}</a></li>
  {{else}}
  <li>No posts yet.</li>
  {{end}}
</ul>
```

Use `html/template` so a title like `Go & Ruby` renders correctly and body text that looks like markup stays text. Keep user-submitted text as a string. In [blog.css](https://github.com/prp-e/go-for-rubyists/blob/book-rewrite/examples/blog/static/blog.css), `white-space: pre-wrap` preserves line breaks in the post body without requiring HTML in the database.

Embed the HTML and CSS with `go:embed` so the compiled binary carries its assets. Parse the templates once at startup with `template.ParseFS`. Each request executes a named template into a buffer, then writes the status and the HTML. If rendering fails, you still have a chance to return a 500 before sending half a page.

## 4. Start the server and check the round trip

The [main.go](https://github.com/prp-e/go-for-rubyists/blob/book-rewrite/examples/blog/main.go) command opens the database, passes the pool to the router, and runs an HTTP server with timeouts as in Chapter 12. It waits for active requests during shutdown before closing the pool. Run the blog, create a post with a couple of paragraphs, and edit its title. Restart the server and check that the edit persists. Delete the post and revisit its old URL. You should see 404.

The [tests](https://github.com/prp-e/go-for-rubyists/blob/book-rewrite/examples/blog/web_test.go) use a temporary SQLite file and call the router directly. They check the forms, redirects, HTML escaping, rejection of cross-origin writes, and the persistence of an edit after reopening the database. Test a missing row through the router too. It's the difference between proving the SQL returns no row and proving the visitor gets the right response.

**Try it.** Add a search form to the index. Submit a GET with a query parameter, use a bound parameter in the search query, and keep the entered text in the form. Decide whether percent and underscore are search wildcards or literal characters, and test that choice.

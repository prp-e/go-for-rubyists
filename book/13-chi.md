# 13. From Sinatra to Chi

Remember how little ceremony it took to write a route in Sinatra? I wanted to keep that feeling when I started writing Go. The standard library already gave us a perfectly good server in Chapter 9. [Chi](https://github.com/go-chi/chi) adds a router, route groups, and middleware helpers. Your handlers still speak `net/http`, so the tests you already know how to write still work.

## Start with one route

Create an empty directory and run `go mod init example.com/chi-demo`, then `go get github.com/go-chi/chi/v5@v5.5.0`. Save this as `main.go`. When you're done, run `go run .` and visit `http://127.0.0.1:8081/hello/Ruby`.

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Get("/hello/{name}", hello)
    log.Fatal(http.ListenAndServe("127.0.0.1:8081", r))
}

func hello(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintf(w, "Hello, %s!\n", name)
}
```

The request goes through the logger, then the recovery middleware, then your handler. When the handler returns, the logger can report what happened. `Recoverer` can catch a panic from a downstream handler and report a server error if the response can still be written. I still want to handle the ordinary errors in the handler. Otherwise my plan for a missing blog post would involve a stack trace, which is a lot to ask from someone who just mistyped a link.

## Group routes by what they share

In Sinatra I might put a filter in front of several routes. In Chi I can put shared middleware around the routes that need it. `r.Use` applies to the whole router and should be registered before routes. `r.With(middleware).Get(...)` applies middleware to a particular route. `r.Group` creates a group without adding a path prefix, while `r.Route("/posts", ...)` creates a subrouter under a prefix. Here is how the blog will share the post ID between several paths:

```go
r.Route("/posts/{id}", func(r chi.Router) {
    r.Get("/", blog.show)
    r.Get("/edit", blog.edit)
    r.Post("/edit", blog.update)
    r.Post("/delete", blog.remove)
})
```

The names of the path parameters are local to the routing pattern. Outside the router, an ordinary request has no Chi routing context. Test through `router.ServeHTTP` when you want to assert that a route matched, its path parameter was extracted, or middleware ran.

`chi.Router` implements `http.Handler`. Pass it to `http.Server`, wrap it with existing middleware, or call it from `httptest`. The framework makes routing choices. Your own code still decides how to read a request, save data, and render an error. Coming from Rails, that's more work on the first day. When I want to understand where an update failed, I can follow the call from the router to SQL without guessing which callback ran first.

**Try it.** Add a GET route under `/health` that replies with 204. Test the route through the router, then send a POST to the same URL and check that Chi returns 405.
